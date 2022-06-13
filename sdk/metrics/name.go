//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package metrics

type MetricName string

//List of metrics that are collected when metrics server is enabled
const (
	EXTERNAL_PAYLOAD_USED     MetricName = "external_payload_used"
	TASK_EXECUTE_ERROR        MetricName = "task_execute_error"
	TASK_EXECUTE_TIME         MetricName = "task_execute_time"
	TASK_EXECUTION_QUEUE_FULL MetricName = "task_execution_queue_full"
	TASK_PAUSED               MetricName = "task_paused"
	TASK_POLL                 MetricName = "task_poll"
	TASK_POLL_ERROR           MetricName = "task_poll_error"
	TASK_POLL_TIME            MetricName = "task_poll_time"
	TASK_RESULT_SIZE          MetricName = "task_result_size"
	TASK_UPDATE_ERROR         MetricName = "task_update_error"
	TASK_UPDATE_TIME          MetricName = "task_update_time"
	THREAD_UNCAUGHT_EXCEPTION MetricName = "thread_uncaught_exceptions"
	WORKFLOW_INPUT_SIZE       MetricName = "workflow_input_size"
	WORKFLOW_START_ERROR      MetricName = "workflow_start_error"
)
