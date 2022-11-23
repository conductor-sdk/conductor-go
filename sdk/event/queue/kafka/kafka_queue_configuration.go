//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package kafka

import "github.com/conductor-sdk/conductor-go/sdk/event/queue"

var (
	queueName = "kafka"

	bootstrapServerConfigKey = "bootstrap.servers"
)

func NewKafkaQueueConfiguration(queueTopicName string) *queue.QueueConfiguration {
	return queue.NewQueueConfiguration(queueTopicName, queueName)
}

func NewKafkaConsumer(bootstrapServersConfig string) *queue.QueueWorkerConfiguration {
	return queue.NewQueueWorkerConfiguration().
		WithConfiguration(bootstrapServerConfigKey, bootstrapServersConfig)
}

func NewKafkaProducer(bootstrapServersConfig string) *queue.QueueWorkerConfiguration {
	return queue.NewQueueWorkerConfiguration().
		WithConfiguration(bootstrapServerConfigKey, bootstrapServersConfig)
}
