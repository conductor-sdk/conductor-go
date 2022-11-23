//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package queue

import (
	"encoding/json"
	"fmt"
)

var (
	requiredKeyConsumer = "consumer"
	requiredKeyProducer = "producer"
)

type QueueConfiguration struct {
	QueueType                string
	QueueName                string
	queueWorkerConfiguration map[string]QueueWorkerConfiguration
}

func NewQueueConfiguration(queueName string, queueType string) *QueueConfiguration {
	return &QueueConfiguration{
		QueueType:                queueType,
		QueueName:                queueName,
		queueWorkerConfiguration: make(map[string]QueueWorkerConfiguration),
	}
}

func (q *QueueConfiguration) GetConfiguration() (configuration string, err error) {
	for _, requiredKey := range []string{requiredKeyConsumer, requiredKeyProducer} {
		if !q.isQueueWorkerRegistered(requiredKey) {
			return "", fmt.Errorf("queue configuration key not set %s", requiredKey)
		}
	}
	body := make(map[string]map[string]string)
	for key, value := range q.queueWorkerConfiguration {
		body[key] = value.GetConfiguration()
	}
	jsonString, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

func (q *QueueConfiguration) WithConsumer(queueWorkerConfiguration *QueueWorkerConfiguration) *QueueConfiguration {
	q.addQueueWorkerConfiguration(requiredKeyConsumer, *queueWorkerConfiguration)
	return q
}

func (q *QueueConfiguration) WithProducer(queueWorkerConfiguration *QueueWorkerConfiguration) *QueueConfiguration {
	q.addQueueWorkerConfiguration(requiredKeyProducer, *queueWorkerConfiguration)
	return q
}

func (q *QueueConfiguration) addQueueWorkerConfiguration(key string, value QueueWorkerConfiguration) {
	q.queueWorkerConfiguration[key] = value
}

func (q *QueueConfiguration) isQueueWorkerRegistered(key string) bool {
	_, ok := q.queueWorkerConfiguration[key]
	return ok
}
