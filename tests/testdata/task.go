//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package examples

import (
	"github.com/conductor-sdk/conductor-go/sdk/workflow/definition"
)

var HttpTask = definition.NewHttpTask(
	"go_task_of_http_type", // task name
	&definition.HttpInput{ // http input
		Uri: "https://catfact.ninja/fact",
	},
)

var SimpleTask = definition.NewSimpleTask(
	"go_task_of_simple_type",
	"go_task_of_simple_type",
)
