//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package settings

type HttpSettings struct {
	BaseUrl string
	Headers map[string]string
}

func NewHttpDefaultSettings() *HttpSettings {
	return NewHttpSettings(
		"http://localhost:8080/api",
	)
}

func NewHttpSettings(baseUrl string) *HttpSettings {
	return &HttpSettings{
		BaseUrl: baseUrl,
		Headers: map[string]string{
			"Content-Type":    "application/json",
			"Accept":          "application/json",
			"Accept-Encoding": "gzip",
		},
	}
}
