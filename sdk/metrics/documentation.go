//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package metrics

type MetricDocumentation string

const (
	EXTERNAL_PAYLOAD_USED_DOC     MetricDocumentation = "Incremented each time external payload storage is used"
	TASK_ACK_ERROR_DOC            MetricDocumentation = "Task ack has encountered an exception"
	TASK_ACK_FAILED_DOC           MetricDocumentation = "Task ack failed"
	TASK_EXECUTE_ERROR_DOC        MetricDocumentation = "Execution error"
	TASK_EXECUTE_TIME_DOC         MetricDocumentation = "Time to execute a task"
	TASK_EXECUTION_QUEUE_FULL_DOC MetricDocumentation = "Counter to record execution queue has saturated"
	TASK_PAUSED_DOC               MetricDocumentation = "Counter for number of times the task has been polled, when the worker has been paused"
	TASK_POLL_DOC                 MetricDocumentation = "Incremented each time polling is done"
	TASK_POLL_ERROR_DOC           MetricDocumentation = "Client error when polling for a task queue"
	TASK_POLL_TIME_DOC            MetricDocumentation = "Time to poll for a batch of tasks"
	TASK_RESULT_SIZE_DOC          MetricDocumentation = "Records output payload size of a task"
	TASK_UPDATE_ERROR_DOC         MetricDocumentation = "Task status cannot be updated back to server"
	TASK_UPDATE_TIME_DOC          MetricDocumentation = "Time to update for a task"
	THREAD_UNCAUGHT_EXCEPTION_DOC MetricDocumentation = "thread_uncaught_exceptions"
	WORKFLOW_START_ERROR_DOC      MetricDocumentation = "Counter for workflow start errors"
	WORKFLOW_INPUT_SIZE_DOC       MetricDocumentation = "Records input payload size of a workflow"
)
