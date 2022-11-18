//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package status

type WorkflowStatus string

const (
	RunningWorkflow    WorkflowStatus = "RUNNING"
	CompletedWorkflow  WorkflowStatus = "COMPLETED"
	FailedWorkflow     WorkflowStatus = "FAILED"
	TimedOutWorkflow   WorkflowStatus = "TIMED_OUT"
	TerminatedWorkflow WorkflowStatus = "TERMINATED"
	PausedWorkflow     WorkflowStatus = "PAUSED"
)

var (
	WorkflowTerminalStates = []WorkflowStatus{
		CompletedWorkflow,
		FailedWorkflow,
		TimedOutWorkflow,
		TerminatedWorkflow,
	}
)