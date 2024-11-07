package executor

import (
	"context"
	"errors"
	"fmt"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/event/queue"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func (e *WorkflowExecutor) RegisterWorkflowWithContext(ctx context.Context, overwrite bool, workflow *model.WorkflowDef) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, err := e.metadataClient.RegisterWorkflowDef(ctx, overwrite, *workflow)
	return err
}

func (e *WorkflowExecutor) UnRegisterWorkflowWithContext(ctx context.Context, name string, version int32) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, err := e.metadataClient.UnregisterWorkflowDef(ctx, name, version)
	return err
}

func (e *WorkflowExecutor) ExecuteWorkflowWithContext(ctx context.Context, startWorkflowRequest *model.StartWorkflowRequest, waitUntilTask string) (run *model.WorkflowRun, err error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	requestId := ""
	version := startWorkflowRequest.Version
	workflowRun, _, err := e.workflowClient.ExecuteWorkflow(ctx, *startWorkflowRequest, requestId, startWorkflowRequest.Name, version, waitUntilTask)
	if err != nil {
		return nil, err
	}
	return &workflowRun, nil
}

func (e *WorkflowExecutor) StartWorkflowWithContext(ctx context.Context, startWorkflowRequest *model.StartWorkflowRequest) (workflowId string, err error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	id, _, err := e.workflowClient.StartWorkflowWithRequest(
		ctx,
		*startWorkflowRequest,
	)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (e *WorkflowExecutor) GetWorkflowWithContext(ctx context.Context, workflowId string, includeTasks bool) (*model.Workflow, error) {
	return e.getWorkflowWithContext(ctx, 4, workflowId, includeTasks)
}

func (e *WorkflowExecutor) getWorkflowWithContext(ctx context.Context, retry int, workflowId string, includeTasks bool) (*model.Workflow, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	workflow, response, err := e.workflowClient.GetExecutionStatus(
		ctx,
		workflowId,
		&client.WorkflowResourceApiGetExecutionStatusOpts{
			IncludeTasks: optional.NewBool(includeTasks)},
	)

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("no such workflow by Id %s", workflowId)
	}

	if response.StatusCode > 399 && response.StatusCode < 500 && response.StatusCode != 429 {
		return nil, err
	}

	if err != nil {
		if retry < 0 {
			return nil, err
		} else {
			time.Sleep(time.Duration(4-retry) * 10 * time.Second)
			retry = retry - 1
			return e.getWorkflowWithContext(ctx, retry, workflowId, includeTasks)
		}
	}

	return &workflow, nil
}

func (e *WorkflowExecutor) GetWorkflowStatusWithContext(ctx context.Context, workflowId string, includeOutput bool, includeVariables bool) (*model.WorkflowState, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	state, response, err := e.workflowClient.GetWorkflowState(ctx, workflowId, includeOutput, includeVariables)
	if response != nil && response.StatusCode == 404 {
		return nil, nil
	}
	return &state, err
}

func (e *WorkflowExecutor) GetByCorrelationIdsWithContext(ctx context.Context, workflowName string, includeClosed bool, includeTasks bool, correlationIds ...string) (map[string][]model.Workflow, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	workflows, _, err := e.workflowClient.GetWorkflows(
		ctx,
		correlationIds,
		workflowName,
		&client.WorkflowResourceApiGetWorkflowsOpts{
			IncludeClosed: optional.NewBool(includeClosed),
			IncludeTasks:  optional.NewBool(includeTasks),
		})
	if err != nil {
		return nil, err
	}
	return workflows, nil
}

func (e *WorkflowExecutor) GetByCorrelationIdsAndNamesWithContext(ctx context.Context, includeClosed bool, includeTasks bool, correlationIds []string, workflowNames []string) (map[string][]model.Workflow, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	workflows, _, err := e.workflowClient.GetWorkflowsBatch(
		ctx,
		map[string][]string{
			"workflowNames":  workflowNames,
			"correlationIds": correlationIds,
		},
		&client.WorkflowResourceApiGetWorkflowsOpts{
			IncludeClosed: optional.NewBool(includeClosed),
			IncludeTasks:  optional.NewBool(includeTasks),
		})

	if err != nil {
		return nil, err
	}

	return workflows, nil
}

func (e *WorkflowExecutor) SearchWithContext(ctx context.Context, start int32, size int32, query string, freeText string) ([]model.WorkflowSummary, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	workflows, _, err := e.workflowClient.Search(
		ctx,
		&client.WorkflowResourceApiSearchOpts{
			Start:    optional.NewInt32(start),
			Size:     optional.NewInt32(size),
			FreeText: optional.NewString(freeText),
			Query:    optional.NewString(query),
		},
	)
	if err != nil {
		return nil, err
	}

	return workflows.Results, nil
}

func (e *WorkflowExecutor) PauseWithContext(ctx context.Context, workflowId string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, err := e.workflowClient.PauseWorkflow(ctx, workflowId)
	if err != nil {
		return err
	}
	return nil
}

func (e *WorkflowExecutor) ResumeWithContext(ctx context.Context, workflowId string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, err := e.workflowClient.ResumeWorkflow(ctx, workflowId)
	if err != nil {
		return err
	}
	return nil
}

func (e *WorkflowExecutor) TerminateWithContext(ctx context.Context, workflowId string, reason string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if strings.TrimSpace(workflowId) == "" {
		err := errors.New("workflow id cannot be empty when calling terminate workflow API")
		log.Error("Failed to terminate workflow: ", err.Error())
		return err
	}
	_, err := e.workflowClient.Terminate(ctx, workflowId,
		&client.WorkflowResourceApiTerminateOpts{Reason: optional.NewString(reason), TriggerFailureWorkflow: optional.NewBool(false)},
	)

	if err != nil {
		return err
	}

	return nil
}

func (e *WorkflowExecutor) TerminateWithFailureWithContext(ctx context.Context, workflowId string, reason string, triggerFailureWorkflow bool) error {
	if strings.TrimSpace(workflowId) == "" {
		err := errors.New("workflow id cannot be empty when calling terminate workflow API")
		log.Error("Failed to terminate workflow: ", err.Error())
		return err
	}
	_, err := e.workflowClient.Terminate(ctx, workflowId,
		&client.WorkflowResourceApiTerminateOpts{Reason: optional.NewString(reason), TriggerFailureWorkflow: optional.NewBool(triggerFailureWorkflow)},
	)

	if err != nil {
		return err
	}

	return nil
}

func (e *WorkflowExecutor) RestartWithContext(ctx context.Context, workflowId string, useLatestDefinition bool) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, err := e.workflowClient.Restart(
		ctx,
		workflowId,
		&client.WorkflowResourceApiRestartOpts{
			UseLatestDefinitions: optional.NewBool(useLatestDefinition),
		})

	if err != nil {
		return err
	}

	return nil
}

func (e *WorkflowExecutor) RetryWithContext(ctx context.Context, workflowId string, resumeSubworkflowTasks bool) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, err := e.workflowClient.Retry(
		ctx,
		workflowId,
		&client.WorkflowResourceApiRetryOpts{
			ResumeSubworkflowTasks: optional.NewBool(resumeSubworkflowTasks),
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (e *WorkflowExecutor) ReRunWithContext(ctx context.Context, workflowId string, reRunRequest model.RerunWorkflowRequest) (id string, error error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	id, _, err := e.workflowClient.Rerun(
		ctx,
		reRunRequest,
		workflowId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (e *WorkflowExecutor) SkipTasksFromWorkflowWithContext(ctx context.Context, workflowId string, taskReferenceName string, skipTaskRequest model.SkipTaskRequest) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, err := e.workflowClient.SkipTaskFromWorkflow(
		ctx,
		workflowId,
		taskReferenceName,
		skipTaskRequest,
	)

	if err != nil {
		return err
	}

	return nil
}

func (e *WorkflowExecutor) UpdateTaskWithContext(ctx context.Context, taskId string, workflowInstanceId string, status model.TaskResultStatus, output interface{}) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	taskResult, err := getTaskResultFromOutput(taskId, workflowInstanceId, output)
	if err != nil {
		return err
	}

	taskResult.Status = status
	_, _, err = e.taskClient.UpdateTask(ctx, taskResult)
	if err != nil {
		return err
	}

	return nil
}

func (e *WorkflowExecutor) UpdateTaskByRefNameWithContext(ctx context.Context, taskRefName string, workflowInstanceId string, status model.TaskResultStatus, output interface{}) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	outputData, err := model.ConvertToMap(output)
	if err != nil {
		return err
	}

	_, _, err = e.taskClient.UpdateTaskByRefName(ctx, outputData, workflowInstanceId, taskRefName, string(status))
	if err != nil {
		return err
	}

	return nil
}

func (e *WorkflowExecutor) GetTaskWithContext(ctx context.Context, taskId string) (task *model.Task, err error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	t, _, err := e.taskClient.GetTask(ctx, taskId)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (e *WorkflowExecutor) RemoveWorkflowWithContext(ctx context.Context, workflowId string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, err := e.workflowClient.Delete(ctx, workflowId, &client.WorkflowResourceApiDeleteOpts{ArchiveWorkflow: optional.NewBool(false)})
	if err != nil {
		return err
	}

	return nil
}

func (e *WorkflowExecutor) DeleteQueueConfigurationWithContext(ctx context.Context, queueConfiguration queue.QueueConfiguration) (*http.Response, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return e.eventClient.DeleteQueueConfig(ctx, queueConfiguration.QueueType, queueConfiguration.QueueName)
}

func (e *WorkflowExecutor) GetQueueConfigurationWithContext(ctx context.Context, queueConfiguration queue.QueueConfiguration) (map[string]interface{}, *http.Response, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}

	return e.eventClient.GetQueueConfig(ctx, queueConfiguration.QueueType, queueConfiguration.QueueName)
}

func (e *WorkflowExecutor) PutQueueConfigurationWithContext(ctx context.Context, queueConfiguration queue.QueueConfiguration) (*http.Response, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	body, err := queueConfiguration.GetConfiguration()
	if err != nil {
		return nil, err
	}

	return e.eventClient.PutQueueConfig(ctx, body, queueConfiguration.QueueType, queueConfiguration.QueueName)
}

func getTaskResultFromOutput(taskId string, workflowInstanceId string, taskExecutionOutput interface{}) (*model.TaskResult, error) {
	taskResult, ok := taskExecutionOutput.(*model.TaskResult)
	if !ok {
		taskResult = model.NewTaskResult(taskId, workflowInstanceId)
		outputData, err := model.ConvertToMap(taskExecutionOutput)
		if err != nil {
			return nil, err
		}
		taskResult.OutputData = outputData
		taskResult.Status = model.CompletedTask
	}

	return taskResult, nil
}

func (e *WorkflowExecutor) executeWorkflowWithContext(ctx context.Context, workflow *model.WorkflowDef, request *model.StartWorkflowRequest) (workflowId string, err error) {
	startWorkflowRequest := model.StartWorkflowRequest{
		Name:                            request.Name,
		Version:                         request.Version,
		CorrelationId:                   request.CorrelationId,
		Input:                           request.Input,
		TaskToDomain:                    request.TaskToDomain,
		ExternalInputPayloadStoragePath: request.ExternalInputPayloadStoragePath,
		Priority:                        request.Priority,
	}
	if workflow != nil {
		startWorkflowRequest.WorkflowDef = workflow
	}
	workflowId, response, err := e.workflowClient.StartWorkflowWithRequest(
		ctx,
		startWorkflowRequest,
	)
	if err != nil {
		log.Debug(
			"Failed to start workflow",
			", reason: ", err.Error(),
			", name: ", request.Name,
			", version: ", request.Version,
			", input: ", request.Input,
			", workflowId: ", workflowId,
			", response: ", response,
		)
		return "", err
	}
	log.Debug(
		"Started workflow",
		", workflowId: ", workflowId,
		", name: ", request.Name,
		", version: ", request.Version,
		", input: ", request.Input,
	)

	return workflowId, err
}
