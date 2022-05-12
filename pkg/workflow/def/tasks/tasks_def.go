package tasks

//TaskDef Task definition
type TaskDef struct {
	CreateTime                  int64         `json:"createTime"`
	CreatedBy                   string        `json:"createdBy"`
	Name                        string        `json:"name"`
	Description                 string        `json:"description"`
	RetryCount                  int           `json:"retryCount"`
	TimeoutSeconds              int           `json:"timeoutSeconds"`
	InputKeys                   []interface{} `json:"inputKeys"`
	OutputKeys                  []interface{} `json:"outputKeys"`
	TimeoutPolicy               string        `json:"timeoutPolicy"`
	RetryLogic                  string        `json:"retryLogic"`
	RetryDelaySeconds           int           `json:"retryDelaySeconds"`
	ResponseTimeoutSeconds      int           `json:"responseTimeoutSeconds"`
	InputTemplate               struct{}      `json:"inputTemplate"`
	RateLimitPerFrequency       int           `json:"rateLimitPerFrequency"`
	RateLimitFrequencyInSeconds int           `json:"rateLimitFrequencyInSeconds"`
	OwnerEmail                  string        `json:"ownerEmail"`
	BackoffScaleFactor          int           `json:"backoffScaleFactor"`
}

// WorkflowTask Represents the WorkflowTask inside the workflow
type WorkflowTask struct {
	Name                           interface{}   `json:"name"`
	TaskReferenceName              interface{}   `json:"taskReferenceName"`
	Description                    interface{}   `json:"description"`
	InputParameters                struct{}      `json:"inputParameters"`
	Type                           string        `json:"type"`
	DynamicTaskNameParam           interface{}   `json:"dynamicTaskNameParam"`
	CaseValueParam                 interface{}   `json:"caseValueParam"`
	CaseExpression                 interface{}   `json:"caseExpression"`
	ScriptExpression               interface{}   `json:"scriptExpression"`
	DecisionCases                  struct{}      `json:"decisionCases"`
	DynamicForkJoinTasksParam      interface{}   `json:"dynamicForkJoinTasksParam"`
	DynamicForkTasksParam          interface{}   `json:"dynamicForkTasksParam"`
	DynamicForkTasksInputParamName interface{}   `json:"dynamicForkTasksInputParamName"`
	DefaultCase                    []interface{} `json:"defaultCase"`
	ForkTasks                      []interface{} `json:"forkTasks"`
	StartDelay                     int           `json:"startDelay"`
	SubWorkflowParam               interface{}   `json:"subWorkflowParam"`
	JoinOn                         []interface{} `json:"joinOn"`
	Sink                           interface{}   `json:"sink"`
	Optional                       bool          `json:"optional"`
	TaskDefinition                 interface{}   `json:"taskDefinition"`
	RateLimited                    interface{}   `json:"rateLimited"`
	DefaultExclusiveJoinTask       []interface{} `json:"defaultExclusiveJoinTask"`
	AsyncComplete                  bool          `json:"asyncComplete"`
	LoopCondition                  interface{}   `json:"loopCondition"`
	LoopOver                       []interface{} `json:"loopOver"`
	RetryCount                     interface{}   `json:"retryCount"`
	EvaluatorType                  interface{}   `json:"evaluatorType"`
	Expression                     interface{}   `json:"expression"`
}

type WorkflowDef struct {
	OwnerApp                      interface{}   `json:"ownerApp"`
	CreateTime                    interface{}   `json:"createTime"`
	UpdateTime                    interface{}   `json:"updateTime"`
	CreatedBy                     interface{}   `json:"createdBy"`
	UpdatedBy                     interface{}   `json:"updatedBy"`
	Name                          interface{}   `json:"name"`
	Description                   interface{}   `json:"description"`
	Version                       int           `json:"version"`
	Tasks                         []interface{} `json:"tasks"`
	InputParameters               []interface{} `json:"inputParameters"`
	OutputParameters              struct{}      `json:"outputParameters"`
	FailureWorkflow               interface{}   `json:"failureWorkflow"`
	SchemaVersion                 int           `json:"schemaVersion"`
	Restartable                   bool          `json:"restartable"`
	WorkflowStatusListenerEnabled bool          `json:"workflowStatusListenerEnabled"`
	OwnerEmail                    interface{}   `json:"ownerEmail"`
	TimeoutPolicy                 string        `json:"timeoutPolicy"`
	TimeoutSeconds                int           `json:"timeoutSeconds"`
	Variables                     struct{}      `json:"variables"`
	InputTemplate                 struct{}      `json:"inputTemplate"`
}
