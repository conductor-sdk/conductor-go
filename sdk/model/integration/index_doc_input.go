// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package integration

type IndexDocInput struct {
	LlmProvider           string                 `json:"llmProvider,omitempty"`
	Model                 string                 `json:"model,omitempty"`
	EmbeddingModelProvider string                 `json:"embeddingModelProvider,omitempty"`
	EmbeddingModel        string                 `json:"embeddingModel,omitempty"`
	VectorDB              string                 `json:"vectorDB,omitempty"`
	Text                  string                 `json:"text,omitempty"`
	DocId                 string                 `json:"docId,omitempty"`
	Url                   string                 `json:"url,omitempty"`
	MediaType             string                 `json:"mediaType,omitempty"`
	Namespace             string                 `json:"namespace,omitempty"`
	Index                 string                 `json:"index,omitempty"`
	ChunkSize             int                    `json:"chunkSize,omitempty"`
	ChunkOverlap          int                    `json:"chunkOverlap,omitempty"`
	Metadata              map[string]interface{} `json:"metadata,omitempty"`
	Dimensions            *int                   `json:"dimensions,omitempty"`
}