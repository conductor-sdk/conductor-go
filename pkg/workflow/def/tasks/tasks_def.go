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
	DecisionCases struct{}      `json:"decisionCases"`
	DefaultCase   []interface{} `json:"defaultCase"`
	ForkTasks     []interface{} `json:"forkTasks"`
	StartDelay    int           `json:"startDelay"`
	JoinOn        []interface{} `json:"joinOn"`

	DefaultExclusiveJoinTask []interface{} `json:"defaultExclusiveJoinTask"`
	AsyncComplete            bool          `json:"asyncComplete"`
	LoopOver                 []interface{} `json:"loopOver"`
}
