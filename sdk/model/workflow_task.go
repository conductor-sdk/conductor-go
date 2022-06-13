//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

type WorkflowTask struct {
	Name                           string                    `json:"name"`
	TaskReferenceName              string                    `json:"taskReferenceName"`
	Description                    string                    `json:"description,omitempty"`
	InputParameters                map[string]interface{}    `json:"inputParameters,omitempty"`
	Type_                          string                    `json:"type,omitempty"`
	DynamicTaskNameParam           string                    `json:"dynamicTaskNameParam,omitempty"`
	CaseValueParam                 string                    `json:"caseValueParam,omitempty"`
	CaseExpression                 string                    `json:"caseExpression,omitempty"`
	ScriptExpression               string                    `json:"scriptExpression,omitempty"`
	DecisionCases                  map[string][]WorkflowTask `json:"decisionCases,omitempty"`
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
	WorkflowTaskType               string                    `json:"workflowTaskType,omitempty"`
}
