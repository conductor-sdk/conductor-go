// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
package model

type WorkflowTask struct {
	AsyncComplete                  bool                      `json:"asyncComplete,omitempty"`
	CaseExpression                 string                    `json:"caseExpression,omitempty"`
	CaseValueParam                 string                    `json:"caseValueParam,omitempty"`
	DecisionCases                  map[string][]WorkflowTask `json:"decisionCases,omitempty"`
	DefaultCase                    []WorkflowTask            `json:"defaultCase,omitempty"`
	DefaultExclusiveJoinTask       []string                  `json:"defaultExclusiveJoinTask,omitempty"`
	Description                    string                    `json:"description,omitempty"`
	DynamicForkJoinTasksParam      string                    `json:"dynamicForkJoinTasksParam,omitempty"`
	DynamicForkTasksInputParamName string                    `json:"dynamicForkTasksInputParamName,omitempty"`
	DynamicForkTasksParam          string                    `json:"dynamicForkTasksParam,omitempty"`
	DynamicTaskNameParam           string                    `json:"dynamicTaskNameParam,omitempty"`
	EvaluatorType                  string                    `json:"evaluatorType,omitempty"`
	Expression                     string                    `json:"expression,omitempty"`
	ForkTasks                      [][]WorkflowTask          `json:"forkTasks,omitempty"`
	InputParameters                map[string]interface{}    `json:"inputParameters,omitempty"`
	JoinOn                         []string                  `json:"joinOn,omitempty"`
	LoopCondition                  string                    `json:"loopCondition,omitempty"`
	LoopOver                       []WorkflowTask            `json:"loopOver,omitempty"`
	Name                           string                    `json:"name"`
	Optional                       bool                      `json:"optional,omitempty"`
	RateLimited                    bool                      `json:"rateLimited,omitempty"`
	RetryCount                     int32                     `json:"retryCount,omitempty"`
	ScriptExpression               string                    `json:"scriptExpression,omitempty"`
	Sink                           string                    `json:"sink,omitempty"`
	StartDelay                     int32                     `json:"startDelay,omitempty"`
	SubWorkflowParam               *SubWorkflowParams        `json:"subWorkflowParam,omitempty"`
	TaskDefinition                 *TaskDef                  `json:"taskDefinition,omitempty"`
	TaskReferenceName              string                    `json:"taskReferenceName"`
	Type_                          string                    `json:"type,omitempty"`
	WorkflowTaskType               string                    `json:"workflowTaskType,omitempty"`
}
