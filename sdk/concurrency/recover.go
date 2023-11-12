//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package concurrency

import (
	"github.com/conductor-sdk/conductor-go/sdk/metrics"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

func HandlePanicError(message string) {
	err := recover()
	if err == nil {
		return
	}
	metrics.IncrementUncaughtException(message)

	log.Error(
		"Uncaught panic",
		", message: ", message,
		", error: ", err,
		", stack: ", string(debug.Stack()),
	)
}
