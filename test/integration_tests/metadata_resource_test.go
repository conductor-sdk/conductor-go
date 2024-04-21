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
const TaskName = "TEST_GO_SIMPLE_TASK"

func TestRegisterWorkflowDefWithTags(t *testing.T) {

	_, _ = testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), WorkflowName, 1)
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
		assert.Equal(t, tags[0].Value, tag0.Value)

	} else {
		t.Fatal(err2)
	}

	tags2, err := testdata.MetadataClient.GetTagsForWorkflowDef(context.Background(), WorkflowName)

	if err == nil {
		assert.Equal(t, len(tags2), 1)
		assert.Equal(t, tags2[0].Key, "key_0")
		assert.Equal(t, tags2[0].Value, "value_0")
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

func TestUpdateWorkflowDefWithTags(t *testing.T) {

	_, _ = testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), WorkflowName, 1)

	task := model.WorkflowTask{
		Name:              "simple_task",
		TaskReferenceName: "simple_task_ref",
		Description:       "Test Simple Task",
	}

	workflowDef := model.WorkflowDef{
		Name:        WorkflowName,
		Description: "Test Workflow updated by GO SDK",
		Tasks:       []model.WorkflowTask{task},
	}

	tag0 := model.MetadataTag{
		Key:   "key_0",
		Value: "value_0",
	}

	tag1 := model.MetadataTag{
		Key:   "key_1",
		Value: "wf_value_1",
	}

	tag2 := model.MetadataTag{
		Key:   "key_2",
		Value: "wf_value_2",
	}

	tag3 := model.MetadataTag{
		Key:   "key_3",
		Value: "wf_value_3",
	}

	testCases := []struct {
		tags         []model.MetadataTag
		overwrite    bool
		expectedTags []model.MetadataTag
	}{
		{[]model.MetadataTag{tag0}, false, []model.MetadataTag{tag0}},
		{[]model.MetadataTag{tag1, tag2}, true, []model.MetadataTag{tag1, tag2}},
		{[]model.MetadataTag{tag3}, false, []model.MetadataTag{tag1, tag2, tag3}},
	}

	for _, tc := range testCases {
		testdata.MetadataClient.UpdateWorkflowDefWithTags(context.Background(), workflowDef, tc.tags, tc.overwrite)

		fetchedTags, err := testdata.MetadataClient.GetTagsForWorkflowDef(context.Background(), WorkflowName)

		if err == nil {
			assert.Equal(t, len(fetchedTags), len(tc.expectedTags))
			assert.ElementsMatch(t, fetchedTags, tc.expectedTags)

		} else {
			t.Fatal(err)
		}
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

func TestRegisterTaskDefWithTags(t *testing.T) {
	taskDef := model.TaskDef{
		Name:        TaskName,
		Description: "Test task definition created from Go SDK",
	}

	tag0 := model.MetadataTag{
		Key:   "key_0",
		Value: "value_0",
	}

	tags := []model.MetadataTag{tag0}

	testdata.MetadataClient.RegisterTaskDefWithTags(context.Background(), taskDef, tags)

	fetchedTags, _, err := testdata.TagsClient.GetTaskTags(context.Background(), TaskName)

	if err == nil {
		assert.Equal(t, len(tags), 1)
		assert.Equal(t, fetchedTags[0].Key, tag0.Key)
		assert.Equal(t, fetchedTags[0].Value, tag0.Value)

	} else {
		t.Fatal(err)
	}
}

func TestGetTagsForTaskDef(t *testing.T) {
	tags, err := testdata.MetadataClient.GetTagsForTaskDef(context.Background(), TaskName)

	if err == nil {
		assert.Equal(t, len(tags), 1)
		assert.Equal(t, tags[0].Key, "key_0")
		assert.Equal(t, tags[0].Value, "value_0")
	} else {
		t.Fatal(err)
	}
}

func TestUpdateTaskDefWithTags(t *testing.T) {
	taskDef := model.TaskDef{
		Name:        TaskName,
		Description: "Test task definition updated from Go SDK",
	}

	tag1 := model.MetadataTag{
		Key:   "key_1",
		Value: "value_1",
	}

	tag2 := model.MetadataTag{
		Key:   "key_2",
		Value: "value_2",
	}

	tag3 := model.MetadataTag{
		Key:   "key_3",
		Value: "value_3",
	}

	testCases := []struct {
		tags         []model.MetadataTag
		overwrite    bool
		expectedTags []model.MetadataTag
	}{
		{[]model.MetadataTag{tag1, tag2}, true, []model.MetadataTag{tag1, tag2}},
		{[]model.MetadataTag{tag3}, false, []model.MetadataTag{tag1, tag2, tag3}},
	}

	for _, tc := range testCases {
		testdata.MetadataClient.UpdateTaskDefWithTags(context.Background(), taskDef, tc.tags, tc.overwrite)

		fetchedTags, err := testdata.MetadataClient.GetTagsForTaskDef(context.Background(), TaskName)

		if err == nil {
			assert.Equal(t, len(fetchedTags), len(tc.expectedTags))
			assert.ElementsMatch(t, fetchedTags, tc.expectedTags)

		} else {
			t.Fatal(err)
		}
	}

	t.Cleanup(func() {
		_, err := testdata.MetadataClient.UnregisterTaskDef(context.Background(), TaskName)

		if err != nil {
			t.Fatal(
				"Failed to delete task definition. Reason: ", err.Error(),
			)
		}
	})
}
