//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var counterByName = map[MetricName]*prometheus.CounterVec{}

var counterTemplates = map[MetricName]*MetricDetails{
	TASK_POLL: NewMetricDetails(
		TASK_POLL,
		TASK_POLL_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_EXECUTION_QUEUE_FULL: NewMetricDetails(
		TASK_EXECUTION_QUEUE_FULL,
		TASK_EXECUTION_QUEUE_FULL_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	THREAD_UNCAUGHT_EXCEPTION: NewMetricDetails(
		THREAD_UNCAUGHT_EXCEPTION,
		THREAD_UNCAUGHT_EXCEPTION_DOC,
		[]MetricLabel{},
	),
	TASK_POLL_ERROR: NewMetricDetails(
		TASK_POLL_ERROR,
		TASK_POLL_ERROR_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_PAUSED: NewMetricDetails(
		TASK_PAUSED,
		TASK_PAUSED_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_EXECUTE_ERROR: NewMetricDetails(
		TASK_EXECUTE_ERROR,
		TASK_EXECUTE_ERROR_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),

	TASK_UPDATE_ERROR: NewMetricDetails(
		TASK_UPDATE_ERROR,
		TASK_UPDATE_ERROR_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	EXTERNAL_PAYLOAD_USED: NewMetricDetails(
		EXTERNAL_PAYLOAD_USED,
		EXTERNAL_PAYLOAD_USED_DOC,
		[]MetricLabel{
			ENTITY_NAME,
			OPERATION,
			PAYLOAD_TYPE,
		},
	),
	WORKFLOW_START_ERROR: NewMetricDetails(
		WORKFLOW_START_ERROR,
		WORKFLOW_START_ERROR_DOC,
		[]MetricLabel{
			WORKFLOW_TYPE,
		},
	),
}

func IncrementTaskPoll(taskType string) {
	incrementCounter(
		TASK_POLL,
		[]string{
			taskType,
		},
	)
}

func IncrementTaskExecutionQueueFull(taskType string) {
	incrementCounter(
		TASK_EXECUTION_QUEUE_FULL,
		[]string{
			taskType,
		},
	)
}

func IncrementUncaughtException(message string) {
	incrementCounter(
		THREAD_UNCAUGHT_EXCEPTION,
		[]string{
			message,
		},
	)
}

func IncrementTaskPollError(taskType string, err error) {
	incrementCounter(
		TASK_POLL_ERROR,
		[]string{
			taskType,
		},
	)
}

func IncrementTaskPaused(taskType string) {
	incrementCounter(
		TASK_PAUSED,
		[]string{
			taskType,
		},
	)
}

func IncrementTaskExecuteError(taskType string, err error) {
	incrementCounter(
		TASK_EXECUTE_ERROR,
		[]string{
			taskType,
		},
	)
}

func IncrementTaskUpdateError(taskType string, err error) {
	incrementCounter(
		TASK_UPDATE_ERROR,
		[]string{
			taskType,
		},
	)
}

func IncrementExternalPayloadUsed(entityName string, operation string, payloadType string) {
	incrementCounter(
		EXTERNAL_PAYLOAD_USED,
		[]string{
			entityName,
			operation,
			payloadType,
		},
	)
}

func IncrementWorkflowStartError(workflowType string, err error) {
	incrementCounter(
		WORKFLOW_START_ERROR,
		[]string{
			workflowType,
		},
	)
}

func incrementCounter(metricName MetricName, labelValues []string) {
	// We skip incrementing if metrics collection is not yet enabled
	if !collectionEnabled {
		return
	}

	counter := getCounter(metricName, labelValues)
	if *counter != nil {
		(*counter).Inc()
	}
}

func newCounter(metricDetails *MetricDetails) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metricDetails.Name,
			Help: metricDetails.Description,
		},
		metricDetails.Labels,
	)
}

func getCounter(metricName MetricName, labelValues []string) *prometheus.Counter {
	counterVec, ok := counterByName[metricName]
	if !ok {
		return nil
	}
	counter, _ := counterVec.GetMetricWithLabelValues(
		labelValues...,
	)
	return &counter
}
