package client

type OrkesClients struct {
	apiClient *APIClient
}

func (clients *OrkesClients) GetMetadataClient() MetadataClient {
	return NewMetadataClient(clients.apiClient)
}
func (clients *OrkesClients) GetWorkflowClient() WorkflowClient {
	return NewWorkflowClient(clients.apiClient)
}
func (clients *OrkesClients) GetTaskClient() TaskClient {
	return NewTaskClient(clients.apiClient)
}
func (clients *OrkesClients) GetIntegrationClient() IntegrationClient {
	return NewIntegrationClient(clients.apiClient)
}
func (clients *OrkesClients) GetPromptClient() PromptClient {
	return NewPromptClient(clients.apiClient)
}
func (clients *OrkesClients) GetEnvironmentClient() EnvironmentClient {
	return NewEnvironmentClient(clients.apiClient)
}
func (clients *OrkesClients) GetSchedulerClient() SchedulerClient {
	return NewSchedulerClient(clients.apiClient)
}
func (clients *OrkesClients) GetWebhooksConfigClient() WebhooksConfigClient {
	return NewWebhooksConfigClient(clients.apiClient)
}
func (clients *OrkesClients) GetHumanTaskClient() HumanTaskClient {
	return NewHumanTaskClient(clients.apiClient)
}
func (clients *OrkesClients) GetEventHandlerClient() EventHandlerClient {
	return NewEventHandlerClient(clients.apiClient)
}
func (clients *OrkesClients) GetWorkflowBulkClient() WorkflowBulkClient {
	return NewWorkflowBulkClient(clients.apiClient)
}
func (clients *OrkesClients) GetAuthorizationClient() AuthorizationClient {
	return NewAuthorizationClient(clients.apiClient)
}
func (clients *OrkesClients) GetUserClient() UserClient {
	return NewUserClient(clients.apiClient)
}
func (clients *OrkesClients) GetGroupClient() GroupClient {
	return NewGroupClient(clients.apiClient)
}
func (clients *OrkesClients) GetApplicationClient() ApplicationClient {
	return NewApplicationClient(clients.apiClient)
}
func (clients *OrkesClients) GetSecretsClient() SecretsClient {
	return NewSecretsClient(clients.apiClient)
}
