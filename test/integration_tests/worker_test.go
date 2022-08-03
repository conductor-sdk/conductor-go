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
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/internal/testdata"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
)

func TestWorkerBatchSize(t *testing.T) {
	simpleTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SIMPLE").
		Version(1).
		Add(simpleTask)
	err := testdata.TaskRunner.StartWorker(
		simpleTask.ReferenceName(),
		testdata.SimpleWorker,
		5,
		testdata.WorkerPollInterval,
	)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	if testdata.TaskRunner.GetBatchSizeForTask(simpleTask.ReferenceName()) != 5 {
		t.Fatal("unexpected batch size")
	}
	err = testdata.ValidateWorkflowBulk(simpleTaskWorkflow, workflowValidationTimeout, workflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.TaskRunner.SetBatchSize(
		simpleTask.ReferenceName(),
		0,
	)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	if testdata.TaskRunner.GetBatchSizeForTask(simpleTask.ReferenceName()) != 0 {
		t.Fatal("unexpected batch size")
	}
	err = testdata.TaskRunner.SetBatchSize(
		simpleTask.ReferenceName(),
		8,
	)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	if testdata.TaskRunner.GetBatchSizeForTask(simpleTask.ReferenceName()) != 8 {
		t.Fatal("unexpected batch size")
	}
	err = testdata.ValidateWorkflowBulk(simpleTaskWorkflow, workflowValidationTimeout, workflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
}
