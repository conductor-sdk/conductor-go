package integration

type IntegrationUpdate struct {
	Category      string                    `json:"category,omitempty"`
	Configuration map[ConfigKey]interface{} `json:"configuration,omitempty"`
	Description   string                    `json:"description,omitempty"`
	Enabled       bool                      `json:"enabled,omitempty"`
	Type_         string                    `json:"type,omitempty"`
}
type ConfigKey string

const (
	APIKey                    ConfigKey = "api_key"
	User                      ConfigKey = "user"
	Endpoint                  ConfigKey = "endpoint"
	Environment               ConfigKey = "environment"
	ProjectName               ConfigKey = "projectName"
	IndexName                 ConfigKey = "indexName"
	Publisher                 ConfigKey = "publisher"
	Password                  ConfigKey = "password"
	Namespace                 ConfigKey = "namespace"
	BatchSize                 ConfigKey = "batchSize"
	BatchWaitTime             ConfigKey = "batchWaitTime"
	VisibilityTimeout         ConfigKey = "visibilityTimeout"
	ConnectionType            ConfigKey = "connectionType"
	Region                    ConfigKey = "region"
	AWSAccountId              ConfigKey = "awsAccountId"
	ExternalId                ConfigKey = "externalId"
	RoleArn                   ConfigKey = "roleArn"
	Protocol                  ConfigKey = "protocol"
	Port                      ConfigKey = "port"
	SchemaRegistryURL         ConfigKey = "schemaRegistryUrl"
	SchemaRegistryApiKey      ConfigKey = "schemaRegistryApiKey"
	SchemaRegistryApiSecret   ConfigKey = "schemaRegistryApiSecret"
	AuthenticationType        ConfigKey = "authenticationType"
	KeyStorePassword          ConfigKey = "keyStorePassword"
	KeyStoreLocation          ConfigKey = "keyStoreLocation"
	SchemaRegistryAuthType    ConfigKey = "schemaRegistryAuthType"
	ValueSubjectNameStrategy  ConfigKey = "valueSubjectNameStrategy"
	DatasourceURL             ConfigKey = "datasourceURL"
	JdbcDriver                ConfigKey = "jdbcDriver"
	Subscription              ConfigKey = "subscription"
	ServiceAccountCredentials ConfigKey = "serviceAccountCredentials"
	File                      ConfigKey = "file"
	QueueManager              ConfigKey = "queueManager"
	Channel                   ConfigKey = "channel"
)
