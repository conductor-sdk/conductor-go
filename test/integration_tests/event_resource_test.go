//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration_tests

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/internal/testdata"
	"github.com/conductor-sdk/conductor-go/sdk/event/queue"
	"github.com/conductor-sdk/conductor-go/sdk/event/queue/kafka"
)

var (
	kafkaQueueTopicName         = "test_sdk_java_kafka_queue_name"
	kafkaBootstrapServersConfig = "localhost:9092"
)

func TestKafkaQueueConfiguration(t *testing.T) {
	kafkaQueueConfiguration := getKafkaQueueConfiguration()
	_, err := testdata.WorkflowExecutor.DeleteQueueConfiguration(*kafkaQueueConfiguration)
	if err != nil {
		t.Fatal(err)
	}
	_, response, _ := testdata.WorkflowExecutor.GetQueueConfiguration(*kafkaQueueConfiguration)
	if response.StatusCode != 404 {
		t.Fatal("no queue configuration should be found")
	}
	_, err = testdata.WorkflowExecutor.PutQueueConfiguration(*kafkaQueueConfiguration)
	if err != nil {
		t.Fatal(err)
	}
	receivedQueueConfig, _, err := testdata.WorkflowExecutor.GetQueueConfiguration(*kafkaQueueConfiguration)
	if err != nil {
		t.Fatal(err)
	}
	expectedConfigString, _ := kafkaQueueConfiguration.GetConfiguration()
	var expectedConfig map[string]interface{}
	err = json.Unmarshal([]byte(expectedConfigString), &expectedConfig)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(receivedQueueConfig, expectedConfig) {
		t.Fatal(fmt.Errorf("received config differs from expected config"))
	}
	_, err = testdata.WorkflowExecutor.DeleteQueueConfiguration(*kafkaQueueConfiguration)
	if err != nil {
		t.Fatal(err)
	}
}

func getKafkaQueueConfiguration() *queue.QueueConfiguration {
	return kafka.NewKafkaQueueConfiguration(kafkaQueueTopicName).
		WithConsumer(
			kafka.NewKafkaConsumer(kafkaBootstrapServersConfig),
		).
		WithProducer(
			kafka.NewKafkaProducer(kafkaBootstrapServersConfig),
		)
}
