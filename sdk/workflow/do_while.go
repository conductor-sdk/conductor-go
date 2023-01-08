//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package workflow

import (
	"fmt"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

const (
	loopCondition = "loop_count"
)

// NewDoWhileTask Do...While task Create a new DoWhile task.
// terminationCondition is a Javascript expression that evaluates to True or False
func NewDoWhileTask(taskRefName string, terminationCondition string, tasks ...model.WorkflowTask) *Task {
	return &Task{
		model.WorkflowTask{
			Name:              taskRefName,
			TaskReferenceName: taskRefName,
			WorkflowTaskType:  string(DO_WHILE),
			InputParameters:   map[string]interface{}{},
			LoopCondition:     terminationCondition,
			LoopOver:          tasks,
		},
	}
}

// NewLoopTask Loop over N times when N is specified as iterations
func NewLoopTask(taskRefName string, iterations int32, tasks ...model.WorkflowTask) *Task {
	return &Task{
		model.WorkflowTask{
			Name:              taskRefName,
			TaskReferenceName: taskRefName,
			WorkflowTaskType:  string(DO_WHILE),
			InputParameters: map[string]interface{}{
				loopCondition: iterations,
			},
			LoopCondition: getForLoopCondition(taskRefName, loopCondition),
			LoopOver:      tasks,
		},
	}
}

func getForLoopCondition(loopValue string, taskReferenceName string) string {
	return fmt.Sprintf(
		"if ( $.%s['iteration'] < $.%s ) { true; } else { false; }",
		taskReferenceName, loopValue,
	)
}
