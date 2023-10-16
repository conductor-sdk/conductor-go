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
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/internal/testdata"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

func TestRegisterWorkflowDefWithTags(t *testing.T) {
	var value0 interface{} = "value_0"
	tagObjects := []model.TagObject{}

	tag0 := model.TagObject{
		Key:   "key_0",
		Value: &value0,
	}

	tagObjects = append(tagObjects, tag0)

	task := model.WorkflowTask{
		Name:              "simple_task",
		TaskReferenceName: "simple_task_ref",
		Description:       "Test Simple Task",
	}

	workflowTasks := []model.WorkflowTask{}
	workflowTasks = append(workflowTasks, task)

	workflowDefWithTags := model.ExtendedWorkflowDef{
		Name:        "TestGoSDKWorkflow",
		Description: "Test Workflow created by GO SDK",
		Tasks:       workflowTasks,
		Tags:        tagObjects,
	}

	testdata.MetadataClient.RegisterWorkflowDefWithTags(context.Background(), true, workflowDefWithTags)
}
