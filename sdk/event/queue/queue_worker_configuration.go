//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package queue

type QueueWorkerConfiguration struct {
	config map[string]string
}

func NewQueueWorkerConfiguration() *QueueWorkerConfiguration {
	return &QueueWorkerConfiguration{
		config: make(map[string]string),
	}
}

func (q *QueueWorkerConfiguration) WithConfiguration(key string, value string) *QueueWorkerConfiguration {
	q.addConfiguration(key, value)
	return q
}

func (q *QueueWorkerConfiguration) GetConfiguration() map[string]string {
	return q.config
}

func (q *QueueWorkerConfiguration) addConfiguration(key string, value string) {
	q.config[key] = value
}
