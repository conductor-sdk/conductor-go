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
	APIKey                       ConfigKey = "api_key"
	User                         ConfigKey = "user"
	Endpoint                     ConfigKey = "endpoint"
	AuthURL                      ConfigKey = "authUrl"
	Environment                  ConfigKey = "environment"
	ProjectName                  ConfigKey = "projectName"
	IndexName                    ConfigKey = "indexName"
	Publisher                    ConfigKey = "publisher"
	Password                     ConfigKey = "password"
	Namespace                    ConfigKey = "namespace"
	BatchSize                    ConfigKey = "batchSize"
	BatchWaitTime                ConfigKey = "batchWaitTime"
	VisibilityTimeout            ConfigKey = "visibilityTimeout"
	ConnectionType               ConfigKey = "connectionType"
	Consumer                     ConfigKey = "consumer"
	Stream                       ConfigKey = "stream"
	BatchPollConsumersCount      ConfigKey = "batchPollConsumersCount"
	ConsumerType                 ConfigKey = "consumer_type"
	Region                       ConfigKey = "region"
	AWSAccountId                 ConfigKey = "awsAccountId"
	ExternalId                   ConfigKey = "externalId"
	RoleArn                      ConfigKey = "roleArn"
	Protocol                     ConfigKey = "protocol"
	Mechanism                    ConfigKey = "mechanism"
	Port                         ConfigKey = "port"
	SchemaRegistryURL            ConfigKey = "schemaRegistryUrl"
	SchemaRegistryApiKey         ConfigKey = "schemaRegistryApiKey"
	SchemaRegistryApiSecret      ConfigKey = "schemaRegistryApiSecret"
	AuthenticationType           ConfigKey = "authenticationType"
	TruststoreAuthenticationType ConfigKey = "truststoreAuthenticationType"
	TLS                          ConfigKey = "tls"
	CipherSuite                  ConfigKey = "cipherSuite"
	PubSubMethod                 ConfigKey = "pubSubMethod"
	KeyStorePassword             ConfigKey = "keyStorePassword"
	KeyStoreLocation             ConfigKey = "keyStoreLocation"
	SchemaRegistryAuthType       ConfigKey = "schemaRegistryAuthType"
	ValueSubjectNameStrategy     ConfigKey = "valueSubjectNameStrategy"
	DatasourceURL                ConfigKey = "datasourceURL"
	JdbcDriver                   ConfigKey = "jdbcDriver"
	Subscription                 ConfigKey = "subscription"
	ServiceAccountCredentials    ConfigKey = "serviceAccountCredentials"
	File                         ConfigKey = "file"
	TLSFile                      ConfigKey = "tlsFile"
	QueueManager                 ConfigKey = "queueManager"
	GroupId                      ConfigKey = "groupId"
	Channel                      ConfigKey = "channel"
	Dimensions                   ConfigKey = "dimensions"
	DistanceMetric               ConfigKey = "distance_metric"
	IndexingMethod               ConfigKey = "indexing_method"
	InvertedListCount            ConfigKey = "inverted_list_count"
	PullPeriod                   ConfigKey = "pullPeriod"
	PullBatchWaitMillis          ConfigKey = "pullBatchWaitMillis"
)
