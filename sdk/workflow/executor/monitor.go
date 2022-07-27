//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package executor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/concurrency"
	"github.com/conductor-sdk/conductor-go/sdk/model"

	log "github.com/sirupsen/logrus"
)

type WorkflowMonitor struct {
	mutex                        sync.Mutex
	refreshInterval              time.Duration
	executionChannelByWorkflowId map[string]WorkflowExecutionChannel
	workflowClient               *client.WorkflowResourceApiService
}

const (
	defaultMonitorRunningWorkflowsRefreshInterval = 100 * time.Millisecond
)

func NewWorkflowMonitor(workflowClient *client.WorkflowResourceApiService) *WorkflowMonitor {
	workflowMonitor := &WorkflowMonitor{
		refreshInterval:              defaultMonitorRunningWorkflowsRefreshInterval,
		executionChannelByWorkflowId: make(map[string]WorkflowExecutionChannel),
		workflowClient:               workflowClient,
	}
	go workflowMonitor.monitorRunningWorkflowsDaemon()
	return workflowMonitor
}

func (w *WorkflowMonitor) generateWorkflowExecutionChannel(workflowId string) (WorkflowExecutionChannel, error) {
	channel := make(WorkflowExecutionChannel, 1)
	err := w.addWorkflowExecutionChannel(workflowId, channel)
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (w *WorkflowMonitor) monitorRunningWorkflowsDaemon() {
	defer concurrency.HandlePanicError("monitor_running_workflows")
	for {
		err := w.monitorRunningWorkflows()
		if err != nil {
			log.Warning(
				"Failed to monitor running workflows",
				", error: ", err.Error(),
			)
		}
		time.Sleep(w.refreshInterval)
	}
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

func (w *WorkflowMonitor) getWorkflowsInTerminalState() ([]*model.Workflow, error) {
	runningWorkflowIdList, err := w.getRunningWorkflowIdList()
	if err != nil {
		return nil, err
	}
	workflowsInTerminalState := make([]*model.Workflow, 0)
	for _, workflowId := range runningWorkflowIdList {
		workflow, response, err := w.workflowClient.GetExecutionStatus(
			context.Background(),
			workflowId,
			nil,
		)
		if err != nil {
			log.Debug(
				"Failed to get workflow execution status",
				", reason: ", err.Error(),
				", workflowId: ", workflowId,
				", response: ", response,
			)
			return nil, err
		}
		if isWorkflowInTerminalState(&workflow) {
			workflowsInTerminalState = append(workflowsInTerminalState, &workflow)
		}
	}
	return workflowsInTerminalState, nil
}

func (w *WorkflowMonitor) addWorkflowExecutionChannel(workflowId string, executionChannel WorkflowExecutionChannel) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.executionChannelByWorkflowId[workflowId] = executionChannel
	log.Debug(
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
	return runningWorkflowIdList, nil
}

func (w *WorkflowMonitor) notifyFinishedWorkflow(workflowId string, workflow *model.Workflow) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	log.Debug(fmt.Sprintf("Notifying finished workflowId: %s", workflowId))
	executionChannel, ok := w.executionChannelByWorkflowId[workflowId]
	if !ok {
		return fmt.Errorf("execution channel not found for workflowId: %s", workflowId)
	}
	executionChannel <- workflow
	log.Debug("Sent finished workflow through channel")
	close(executionChannel)
	log.Debug("Closed client workflow execution channel")
	delete(w.executionChannelByWorkflowId, workflowId)
	log.Debug("Deleted workflow execution channel")
	return nil
}

func isWorkflowInTerminalState(workflow *model.Workflow) bool {
	for _, terminalState := range model.WorkflowTerminalStates {
		if workflow.Status == terminalState {
			return true
		}
	}
	return false
}
