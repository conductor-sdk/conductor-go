// This is the task definition definied as part of the WorkflowDef. The tasks definied in
// the Workflow definition are saved as part of WorkflowDef#getTasks
package model

type WorkflowTask struct {
	Name                           string                    `json:"name"`
	TaskReferenceName              string                    `json:"taskReferenceName"`
	Description                    string                    `json:"description,omitempty"`
	InputParameters                map[string]interface{}    `json:"inputParameters,omitempty"`
	Type_                          string                    `json:"type,omitempty"`
	DynamicTaskNameParam           string                    `json:"dynamicTaskNameParam,omitempty"`
	// Deprecated: Use EvaluatorType and Expression combination
	CaseValueParam                 string                    `json:"caseValueParam,omitempty"`
	// Deprecated: Use EvaluatorType and Expression combination
	CaseExpression                 string                    `json:"caseExpression,omitempty"`
	ScriptExpression               string                    `json:"scriptExpression,omitempty"`
	DecisionCases                  map[string][]WorkflowTask `json:"decisionCases,omitempty"`
	// Deprecated
	DynamicForkJoinTasksParam      string                    `json:"dynamicForkJoinTasksParam,omitempty"`
	DynamicForkTasksParam          string                    `json:"dynamicForkTasksParam,omitempty"`
	DynamicForkTasksInputParamName string                    `json:"dynamicForkTasksInputParamName,omitempty"`
	DefaultCase                    []WorkflowTask            `json:"defaultCase,omitempty"`
	ForkTasks                      [][]WorkflowTask          `json:"forkTasks,omitempty"`
	StartDelay                     int32                     `json:"startDelay,omitempty"`
	SubWorkflowParam               *SubWorkflowParams        `json:"subWorkflowParam,omitempty"`
	JoinOn                         []string                  `json:"joinOn,omitempty"`
	Sink                           string                    `json:"sink,omitempty"`
	Optional                       bool                      `json:"optional,omitempty"`
	TaskDefinition                 *TaskDef                  `json:"taskDefinition,omitempty"`
	RateLimited                    bool                      `json:"rateLimited,omitempty"`
	DefaultExclusiveJoinTask       []string                  `json:"defaultExclusiveJoinTask,omitempty"`
	AsyncComplete                  bool                      `json:"asyncComplete,omitempty"`
	LoopCondition                  string                    `json:"loopCondition,omitempty"`
	LoopOver                       []WorkflowTask            `json:"loopOver,omitempty"`
	RetryCount                     int32                     `json:"retryCount,omitempty"`
	EvaluatorType                  string                    `json:"evaluatorType,omitempty"`
	Expression                     string                    `json:"expression,omitempty"`
	// Deprecated
	WorkflowTaskType               string                    `json:"workflowTaskType,omitempty"`
	JoinStatus                     string                    `json:"joinStatus,omitempty"`
	CacheConfig                    *CacheConfig              `json:"cacheConfig,omitempty"`
	Permissive                     bool                      `json:"permissive,omitempty"`
	OnStateChange                  map[string][]StateChangeEvent `json:"onStateChange,omitempty"`
}

type CacheConfig struct {
	Key          string `json:"key,omitempty"`
	TtlInSeconds int    `json:"ttlInSecond,omitempty"`
}