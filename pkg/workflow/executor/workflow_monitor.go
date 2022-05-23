package executor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/concurrency"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/sirupsen/logrus"
)

type WorkflowExecutionChannel chan *http_model.Workflow

type WorkflowMonitor struct {
	mutex                        sync.Mutex
	refreshInterval              time.Duration
	executionChannelByWorkflowId map[string]executionChannel
	workflowClient               *conductor_http_client.WorkflowResourceApiService
}

type executionChannel struct {
	ClientWorkflowExecutionChannel  WorkflowExecutionChannel
	TimeoutWorkflowExecutionChannel chan bool
}

const (
	defaultMonitorRunningWorkflowsRefreshInterval = 100 * time.Millisecond
)

func NewWorkflowMonitor(workflowClient *conductor_http_client.WorkflowResourceApiService) *WorkflowMonitor {
	return &WorkflowMonitor{
		refreshInterval:              defaultMonitorRunningWorkflowsRefreshInterval,
		executionChannelByWorkflowId: make(map[string]executionChannel),
		workflowClient:               workflowClient,
	}
}

func (w *WorkflowMonitor) MonitorRunningWorkflows() {
	defer concurrency.OnError("monitor_running_workflows")
	for {
		err := w.monitorRunningWorkflows()
		if err != nil {
			logrus.Warning(
				"Failed to monitor running workflows",
				", error: ", err.Error(),
			)
		}
		time.Sleep(w.refreshInterval)
	}
}

func (w *WorkflowMonitor) GenerateWorkflowExecutionChannel(workflowId string) (WorkflowExecutionChannel, error) {
	workflowExecutionChannel := make(WorkflowExecutionChannel, 1)
	w.addWorkflowExecutionChannel(
		workflowId,
		executionChannel{
			ClientWorkflowExecutionChannel:  workflowExecutionChannel,
			TimeoutWorkflowExecutionChannel: nil,
		},
	)
	return workflowExecutionChannel, nil
}

func (w *WorkflowMonitor) GenerateWorkflowExecutionChannelWithTimeout(workflowId string, timeout time.Duration) (WorkflowExecutionChannel, error) {
	workflowExecutionChannel := make(WorkflowExecutionChannel, 1)
	timeoutExecutionChannel := make(chan bool, 1)
	err := w.addWorkflowExecutionChannel(
		workflowId,
		executionChannel{
			ClientWorkflowExecutionChannel:  workflowExecutionChannel,
			TimeoutWorkflowExecutionChannel: timeoutExecutionChannel,
		},
	)
	if err != nil {
		return nil, err
	}
	go w.waitForWorkflowCompletionUntilTimeout(
		workflowId,
		timeoutExecutionChannel,
		timeout,
	)
	return workflowExecutionChannel, nil
}

func (w *WorkflowMonitor) monitorRunningWorkflows() error {
	workflowsInTerminalState, err := w.getWorkflowsInTerminalState()
	if err != nil {
		return err
	}
	for _, workflow := range workflowsInTerminalState {
		err = w.notifyFinishedWorkflow(workflow.WorkflowId, workflow)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *WorkflowMonitor) getWorkflowsInTerminalState() ([]*http_model.Workflow, error) {
	runningWorkflowIdList, err := w.getRunningWorkflowIdList()
	if err != nil {
		return nil, err
	}
	workflowsInTerminalState := make([]*http_model.Workflow, 0)
	for _, workflowId := range runningWorkflowIdList {
		workflow, _, err := w.workflowClient.GetExecutionStatus(
			context.Background(),
			workflowId,
			nil,
		)
		if err != nil {
			return nil, err
		}
		if isWorkflowInTerminalState(&workflow) {
			workflowsInTerminalState = append(workflowsInTerminalState, &workflow)
		}
	}
	return workflowsInTerminalState, nil
}

func (w *WorkflowMonitor) waitForWorkflowCompletionUntilTimeout(workflowId string, timeoutWorkflowExecutionChannel chan bool, timeout time.Duration) {
	defer concurrency.OnError(
		fmt.Sprint(
			"Failed to waitForWorkflowCompletionUntilTimeout",
			", workflowId: ", workflowId,
			", timeout: ", timeout,
		),
	)
	select {
	case <-timeoutWorkflowExecutionChannel:
		logrus.Debug(
			"Stopped waiting for workflow completion",
			", workflowId: ", workflowId,
		)
		break
	case <-time.After(timeout):
		logrus.Debug(
			fmt.Sprint(
				"Timeout waiting for completion of workflow",
				", with id: ", workflowId,
			),
		)
		w.notifyFinishedWorkflow(workflowId, nil)
		break
	}
}

func (w *WorkflowMonitor) addWorkflowExecutionChannel(workflowId string, ec executionChannel) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.executionChannelByWorkflowId[workflowId] = ec
	logrus.Debug(
		fmt.Sprint(
			"Added workflow execution channel",
			", workflowId: ", workflowId,
		),
	)
	return nil
}

func (w *WorkflowMonitor) getRunningWorkflowIdList() ([]string, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	i := 0
	runningWorkflowIdList := make([]string, len(w.executionChannelByWorkflowId))
	for workflowId := range w.executionChannelByWorkflowId {
		runningWorkflowIdList[i] = workflowId
		i += 1
	}
	logrus.Debug("Got running workflowId list")
	return runningWorkflowIdList, nil
}

func (w *WorkflowMonitor) notifyFinishedWorkflow(workflowId string, workflow *http_model.Workflow) error {
	executionChannel, err := w.getExecutionChannel(workflowId)
	if err != nil {
		return err
	}
	logrus.Debug("Notifying finished workflow: ", *workflow)
	executionChannel.ClientWorkflowExecutionChannel <- workflow
	logrus.Debug("Sent workflow through client channel")
	executionChannel.TimeoutWorkflowExecutionChannel <- true
	logrus.Debug("Sent signal to stop waiting for workflow execution")
	return w.removeExecutionChannel(workflowId)
}

func (w *WorkflowMonitor) getExecutionChannel(workflowId string) (executionChannel, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	executionChannel, ok := w.executionChannelByWorkflowId[workflowId]
	if !ok {
		return executionChannel, fmt.Errorf("execution channel not found for workflowId: %s", workflowId)
	}
	return executionChannel, nil
}

func (w *WorkflowMonitor) removeExecutionChannel(workflowId string) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	delete(w.executionChannelByWorkflowId, workflowId)
	logrus.Debug(
		fmt.Sprint(
			"Deleted workflow execution channel",
			", workflowId: ", workflowId,
		),
	)
	return nil
}
