package def

import (
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/tasks"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
	"net/http"
)

func NewConductorWorkflow(executor *executor.WorkflowExecutor) *conductorWorkflow {
	return &conductorWorkflow{
		executor: executor,
	}
}

type conductorWorkflow struct {
	executor *executor.WorkflowExecutor
	name     string
	version  int32
	tasks    []tasks.Task
}

func (workflow *conductorWorkflow) Name(name string) *conductorWorkflow {
	workflow.name = name
	return workflow
}
func (workflow *conductorWorkflow) Version(version int32) *conductorWorkflow {
	workflow.version = version
	return workflow
}
func (workflow *conductorWorkflow) Add(task tasks.Task) *conductorWorkflow {
	workflow.tasks = append(workflow.tasks, task)
	return workflow
}

func (workflow *conductorWorkflow) Register(overwrite bool) (*http.Response, error) {
	response, error := workflow.executor.RegisterWorkflow(workflow.ToWorkflowDef())
	return response, error
}

func (workflow *conductorWorkflow) execute() (string, error) {
	return "", nil
}

func (workflow *conductorWorkflow) ToWorkflowDef() *http_model.WorkflowDef {
	workflowTasks := make([]http_model.WorkflowTask, 0)
	for _, task := range workflow.tasks {
		workflowTask := task.ToWorkflowTask()
		workflowTasks = append(workflowTasks, *workflowTask)
	}

	def := &http_model.WorkflowDef{
		OwnerApp:         "",
		Name:             workflow.name,
		Description:      "",
		Version:          workflow.version,
		Tasks:            workflowTasks,
		InputParameters:  nil,
		OutputParameters: nil,
		FailureWorkflow:  "",
		SchemaVersion:    2,
		OwnerEmail:       "",
		TimeoutPolicy:    "",
		TimeoutSeconds:   0,
		Variables:        nil,
		InputTemplate:    nil,
	}
	return def
}
