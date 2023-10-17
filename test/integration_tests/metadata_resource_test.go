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
	task := model.WorkflowTask{
		Name:              "simple_task",
		TaskReferenceName: "simple_task_ref",
		Description:       "Test Simple Task",
	}

	workflowTasks := []model.WorkflowTask{}
	workflowTasks = append(workflowTasks, task)

	workflowDef := model.WorkflowDef{
		Name:        WorkflowName,
		Description: "Test Workflow created by GO SDK",
		Tasks:       workflowTasks,
	}

	tag0 := model.MetadataTag{
		Key:   "key_0",
		Value: "value_0",
	}

	metadataTags := []model.MetadataTag{tag0}

	testdata.MetadataClient.RegisterWorkflowDefWithTags(context.Background(), true, workflowDef, metadataTags)

	tags, _, err2 := testdata.TagsClient.GetWorkflowTags(context.Background(), WorkflowName)

	if err2 == nil {
		assert.Equal(t, len(tags), 1)
		assert.Equal(t, tags[0].Key, tag0.Key)
		assert.Equal(t, *tags[0].Value, tag0.Value)

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

	workflowTasks := []model.WorkflowTask{task}

	workflowDef := model.WorkflowDef{
		Name:        WorkflowName,
		Description: "Test Workflow updated by GO SDK",
		Tasks:       workflowTasks,
	}

	workflowDefs := []model.WorkflowDef{workflowDef}

	testdata.MetadataClient.Update(context.Background(), workflowDefs)

	tags, _, err := testdata.TagsClient.GetWorkflowTags(context.Background(), WorkflowName)

	if err == nil {
		assert.Equal(t, 1, len(tags))
		assert.Equal(t, "key_0", tags[0].Key)
		assert.Equal(t, "value_0", *tags[0].Value)

	} else {
		t.Fatal(err)
	}

	tag1 := model.MetadataTag{
		Key:   "key_1",
		Value: "wf_value_1",
	}

	tag2 := model.MetadataTag{
		Key:   "key_2",
		Value: "wf_value_2",
	}

	metadataTags := []model.MetadataTag{tag1, tag2}

	testdata.MetadataClient.UpdateWorkflowDefWithTags(context.Background(), workflowDef, metadataTags, true)

	tags, _, err = testdata.TagsClient.GetWorkflowTags(context.Background(), WorkflowName)

	if err == nil {
		assert.Equal(t, 2, len(tags))
		assert.Equal(t, tag1.Key, tags[1].Key)
		assert.Equal(t, tag1.Value, *tags[1].Value)
		assert.Equal(t, tag2.Key, tags[0].Key)
		assert.Equal(t, tag2.Value, *tags[0].Value)

	} else {
		t.Fatal(err)
	}

	tag3 := model.MetadataTag{
		Key:   "key_3",
		Value: "wf_value_3",
	}

	metadataTags = []model.MetadataTag{tag3}

	testdata.MetadataClient.UpdateWorkflowDefWithTags(context.Background(), workflowDef, metadataTags, false)

	tags, _, err = testdata.TagsClient.GetWorkflowTags(context.Background(), WorkflowName)

	if err == nil {
		assert.Equal(t, 3, len(tags))
		assert.Equal(t, tag1.Key, tags[2].Key)
		assert.Equal(t, tag1.Value, *tags[2].Value)
		assert.Equal(t, tag2.Key, tags[0].Key)
		assert.Equal(t, tag2.Value, *tags[0].Value)
		assert.Equal(t, tag3.Key, tags[1].Key)
		assert.Equal(t, tag3.Value, *tags[1].Value)

	} else {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_, err := testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), WorkflowName, 1)

		if err != nil {
			t.Fatal(
				"Failed to delete workflow. Reason: ", err.Error(),
			)
		}
	})
}
