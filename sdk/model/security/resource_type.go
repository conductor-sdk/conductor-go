// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package security

// ResourceType represents the types of resources in the system
type ResourceType string

const (
	ResourceTypeWorkflow             ResourceType = "WORKFLOW"
	ResourceTypeWorkflowDef          ResourceType = "WORKFLOW_DEF"
	ResourceTypeWorkflowSchedule     ResourceType = "WORKFLOW_SCHEDULE"
	ResourceTypeEventHandler         ResourceType = "EVENT_HANDLER"
	ResourceTypeTaskDef              ResourceType = "TASK_DEF"
	ResourceTypeTaskRefName          ResourceType = "TASK_REF_NAME"
	ResourceTypeTaskId               ResourceType = "TASK_ID"
	ResourceTypeApplication          ResourceType = "APPLICATION"
	ResourceTypeUser                 ResourceType = "USER"
	ResourceTypeSecretName           ResourceType = "SECRET_NAME"
	ResourceTypeEnvVariable          ResourceType = "ENV_VARIABLE"
	ResourceTypeTag                  ResourceType = "TAG"
	ResourceTypeDomain               ResourceType = "DOMAIN"
	ResourceTypeIntegrationProvider  ResourceType = "INTEGRATION_PROVIDER"
	ResourceTypeIntegration          ResourceType = "INTEGRATION"
	ResourceTypePrompt               ResourceType = "PROMPT"
	ResourceTypeUserFormTemplate     ResourceType = "USER_FORM_TEMPLATE"
	ResourceTypeSchema               ResourceType = "SCHEMA"
	ResourceTypeClusterConfig        ResourceType = "CLUSTER_CONFIG"
	ResourceTypeWebhook              ResourceType = "WEBHOOK"
)