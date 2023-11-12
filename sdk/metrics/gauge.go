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

var gaugeByName = map[MetricName]*prometheus.GaugeVec{}

var gaugeTemplates = map[MetricName]*MetricDetails{
	WORKFLOW_INPUT_SIZE: NewMetricDetails(
		WORKFLOW_INPUT_SIZE,
		WORKFLOW_INPUT_SIZE_DOC,
		[]MetricLabel{
			WORKFLOW_TYPE,
			WORKFLOW_VERSION,
		},
	),
	TASK_RESULT_SIZE: NewMetricDetails(
		TASK_RESULT_SIZE,
		TASK_RESULT_SIZE_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_POLL_TIME: NewMetricDetails(
		TASK_POLL_TIME,
		TASK_POLL_TIME_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_EXECUTE_TIME: NewMetricDetails(
		TASK_EXECUTE_TIME,
		TASK_EXECUTE_TIME_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_UPDATE_TIME: NewMetricDetails(
		TASK_UPDATE_TIME,
		TASK_UPDATE_TIME_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
}

func RecordWorkflowInputPayloadSize(workflowType string, version string, payloadSize float64) {
	setGauge(
		WORKFLOW_INPUT_SIZE,
		[]string{
			workflowType,
			version,
		},
		payloadSize,
	)
}

func RecordTaskResultPayloadSize(taskType string, payloadSize float64) {
	setGauge(
		TASK_RESULT_SIZE,
		[]string{
			taskType,
		},
		payloadSize,
	)
}

func RecordTaskPollTime(taskType string, timeSpent float64) {
	setGauge(
		TASK_POLL_TIME,
		[]string{
			taskType,
		},
		timeSpent,
	)
}

func RecordTaskUpdateTime(taskType string, timeSpent float64) {
	setGauge(
		TASK_UPDATE_TIME,
		[]string{
			taskType,
		},
		timeSpent,
	)
}

func RecordTaskExecuteTime(taskType string, timeSpent float64) {
	setGauge(
		TASK_EXECUTE_TIME,
		[]string{
			taskType,
		},
		timeSpent,
	)
}

func newGauge(metricDetails *MetricDetails) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metricDetails.Name,
			Help: metricDetails.Description,
		},
		metricDetails.Labels,
	)
}

func setGauge(metricName MetricName, labelValues []string, value float64) {
	// We skip setting gauge if Metrics collection is not enabled
	if !collectionEnabled {
		return
	}

	gauge := getGauge(metricName, labelValues)
	if *gauge != nil {
		(*gauge).Set(value)
	}
}

func getGauge(metricName MetricName, labelValues []string) *prometheus.Gauge {
	gaugeVec, ok := gaugeByName[metricName]
	if !ok {
		return nil
	}
	gauge, _ := gaugeVec.GetMetricWithLabelValues(
		labelValues...,
	)
	return &gauge
}
