// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package integration

type ConfigKey string

const (
	APIKey                    ConfigKey = "api_key"
	User                      ConfigKey = "user"
	Endpoint                  ConfigKey = "endpoint"
	AuthURL                   ConfigKey = "authUrl"
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
	Consumer                  ConfigKey = "consumer"
	Stream                    ConfigKey = "stream"
	BatchPollConsumersCount   ConfigKey = "batchPollConsumersCount"
	ConsumerType              ConfigKey = "consumer_type"
	Region                    ConfigKey = "region"
	AWSAccountId              ConfigKey = "awsAccountId"
	ExternalId                ConfigKey = "externalId"
	RoleArn                   ConfigKey = "roleArn"
	Protocol                  ConfigKey = "protocol"
	Mechanism                 ConfigKey = "mechanism"
	Port                      ConfigKey = "port"
	SchemaRegistryURL         ConfigKey = "schemaRegistryUrl"
	SchemaRegistryApiKey      ConfigKey = "schemaRegistryApiKey"
	SchemaRegistryApiSecret   ConfigKey = "schemaRegistryApiSecret"
	AuthenticationType        ConfigKey = "authenticationType"
	TruststoreAuthenticationType ConfigKey = "truststoreAuthenticationType"
	TLS                       ConfigKey = "tls"
	CipherSuite               ConfigKey = "cipherSuite"
	PubSubMethod              ConfigKey = "pubSubMethod"
	KeyStorePassword          ConfigKey = "keyStorePassword"
	KeyStoreLocation          ConfigKey = "keyStoreLocation"
	SchemaRegistryAuthType    ConfigKey = "schemaRegistryAuthType"
	ValueSubjectNameStrategy  ConfigKey = "valueSubjectNameStrategy"
	DatasourceURL             ConfigKey = "datasourceURL"
	JdbcDriver                ConfigKey = "jdbcDriver"
	Subscription              ConfigKey = "subscription"
	ServiceAccountCredentials ConfigKey = "serviceAccountCredentials"
	File                      ConfigKey = "file"
	TLSFile                   ConfigKey = "tlsFile"
	QueueManager              ConfigKey = "queueManager"
	GroupId                   ConfigKey = "groupId"
	Channel                   ConfigKey = "channel"
	Dimensions                ConfigKey = "dimensions"
	DistanceMetric            ConfigKey = "distance_metric"
	IndexingMethod            ConfigKey = "indexing_method"
	InvertedListCount         ConfigKey = "inverted_list_count"
	PullPeriod                ConfigKey = "pullPeriod"
	PullBatchWaitMillis       ConfigKey = "pullBatchWaitMillis"
)