//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package workflow

import "github.com/conductor-sdk/conductor-go/sdk/model"

const (
	sqsEventPrefix       = "sqs"
	conductorEventPrefix = "conductor"
)

// EventTask Task to publish Events to external queuing systems like SQS, NATS, AMQP etc.
func NewSqsEventTask(taskRefName string, queueName string) *Task {
	return newEventTask(
		taskRefName,
		sqsEventPrefix,
		queueName,
	)
}

func NewConductorEventTask(taskRefName string, eventName string) *Task {
	return newEventTask(
		taskRefName,
		conductorEventPrefix,
		eventName,
	)
}

func newEventTask(taskRefName string, eventPrefix string, eventSuffix string) *Task {
	return &Task{
		model.WorkflowTask{
			Name:              taskRefName,
			TaskReferenceName: taskRefName,
			WorkflowTaskType:  string(EVENT),
			Sink:              eventPrefix + ":" + eventSuffix,
		},
	}
}
