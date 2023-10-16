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
	"github.com/stretchr/testify/assert"
)

const WorkflowName = "TestGoSDKWorkflowWithTags"

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
		Name:        WorkflowName,
		Description: "Test Workflow created by GO SDK",
		Tasks:       workflowTasks,
		Tags:        tagObjects,
	}

	testdata.MetadataClient.RegisterWorkflowDefWithTags(context.Background(), true, workflowDefWithTags)

	tags, _, err2 := testdata.TagsClient.GetWorkflowTags(context.Background(), WorkflowName)

	if err2 == nil {
		assert.Equal(t, len(tags), 1)
		assert.Equal(t, tags[0].Key, tag0.Key)
		assert.Equal(t, *tags[0].Value, *tag0.Value)

	} else {
		t.Fatal(err2)
	}
}

func TestUpdateWorkflowDefWithTags(t *testing.T) {
	task := model.WorkflowTask{
		Name:              "simple_task",
		TaskReferenceName: "simple_task_ref",
		Description:       "Test Simple Task",
	}

	workflowTasks := []model.WorkflowTask{}
	workflowTasks = append(workflowTasks, task)

	var value1 interface{} = "wf_value_1"
	var value2 interface{} = "wf_value_2"

	tagObjects := []model.TagObject{}

	tag1 := model.TagObject{
		Key:   "key_1",
		Value: &value1,
	}

	tag2 := model.TagObject{
		Key:   "key_2",
		Value: &value2,
	}

	tagObjects = append(tagObjects, tag1)
	tagObjects = append(tagObjects, tag2)

	workflowDefWithTags := model.ExtendedWorkflowDef{
		Name:        WorkflowName,
		Description: "Test Workflow created by GO SDK",
		Tasks:       workflowTasks,
		Tags:        tagObjects,
	}

	workflowDefs := []model.ExtendedWorkflowDef{workflowDefWithTags}

	testdata.MetadataClient.UpdateWorkflowDefWithTags(context.Background(), workflowDefs)

	tags, _, err2 := testdata.TagsClient.GetWorkflowTags(context.Background(), WorkflowName)

	if err2 == nil {
		assert.Equal(t, len(tags), 2)
		assert.Equal(t, tags[0].Key, tag1.Key)
		assert.Equal(t, *tags[0].Value, *tag1.Value)
		assert.Equal(t, tags[1].Key, tag2.Key)
		assert.Equal(t, *tags[1].Value, *tag2.Value)

	} else {
		t.Fatal(err2)
	}

	var value3 interface{} = "wf_value_3"

	tag3 := model.TagObject{
		Key:   "key_3",
		Value: &value3,
	}

	tagObjects = append(tagObjects, tag3)

	workflowDefWithTags.Tags = tagObjects
	workflowDefWithTags.OverwriteTags = false
	workflowDefs = []model.ExtendedWorkflowDef{workflowDefWithTags}

	testdata.MetadataClient.UpdateWorkflowDefWithTags(context.Background(), workflowDefs)

	tags, _, err2 = testdata.TagsClient.GetWorkflowTags(context.Background(), WorkflowName)

	if err2 == nil {
		assert.Equal(t, len(tags), 3)
		assert.Equal(t, tags[0].Key, tag1.Key)
		assert.Equal(t, *tags[0].Value, *tag1.Value)
		assert.Equal(t, tags[1].Key, tag2.Key)
		assert.Equal(t, *tags[1].Value, *tag2.Value)
		assert.Equal(t, tags[2].Key, tag3.Key)
		assert.Equal(t, *tags[2].Value, *tag3.Value)

	} else {
		t.Fatal(err2)
	}

	_, err3 := testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), WorkflowName, 1)

	if err3 != nil {
		t.Fatal(
			"Failed to delete workflow. Reason: ", err3.Error(),
		)
	}
}
