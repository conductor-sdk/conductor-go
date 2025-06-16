// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
package model

type ServiceMethod struct {
	ExampleInput  map[string]interface{} `json:"exampleInput,omitempty"`
	Id            int64                  `json:"id,omitempty"`
	InputType     string                 `json:"inputType,omitempty"`
	MethodName    string                 `json:"methodName,omitempty"`
	MethodType    string                 `json:"methodType,omitempty"`
	OperationName string                 `json:"operationName,omitempty"`
	OutputType    string                 `json:"outputType,omitempty"`
	RequestParams []RequestParam         `json:"requestParams,omitempty"`
}
