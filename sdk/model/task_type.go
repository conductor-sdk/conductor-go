// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package model

// TaskType represents the type of task in the workflow
const (
	TaskTypeSimple           string = "SIMPLE"
	TaskTypeDynamic          string = "DYNAMIC"
	TaskTypeForkJoin         string = "FORK_JOIN"
	TaskTypeForkJoinDynamic  string = "FORK_JOIN_DYNAMIC"
	TaskTypeDecision         string = "DECISION"
	TaskTypeSwitch           string = "SWITCH"
	TaskTypeJoin             string = "JOIN"
	TaskTypeDoWhile          string = "DO_WHILE"
	TaskTypeSubWorkflow      string = "SUB_WORKFLOW"
	TaskTypeStartWorkflow    string = "START_WORKFLOW"
	TaskTypeEvent            string = "EVENT"
	TaskTypeWait             string = "WAIT"
	TaskTypeHuman            string = "HUMAN"
	TaskTypeUserDefined      string = "USER_DEFINED"
	TaskTypeHTTP             string = "HTTP"
	TaskTypeLambda           string = "LAMBDA"
	TaskTypeInline           string = "INLINE"
	TaskTypeExclusiveJoin    string = "EXCLUSIVE_JOIN"
	TaskTypeTerminate        string = "TERMINATE"
	TaskTypeKafkaPublish     string = "KAFKA_PUBLISH"
	TaskTypeJsonJQTransform  string = "JSON_JQ_TRANSFORM"
	TaskTypeSetVariable      string = "SET_VARIABLE"
	TaskTypeFork             string = "FORK"
	TaskTypeNoop             string = "NOOP"
)

type TaskType struct {
}