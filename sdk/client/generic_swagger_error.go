//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package client

import "fmt"

// GenericSwaggerError Provides access to the body, error and model on returned errors.
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

type GenericSwaggerErrorV2 struct {
	body  string
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericSwaggerErrorV2) Error() string {
	return fmt.Sprintf("error: %s, body: %s", e.error, e.body)
}

// Body returns the raw bytes of the response
func (e GenericSwaggerErrorV2) Body() string {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericSwaggerErrorV2) Model() interface{} {
	return e.model
}

// Error returns non-empty string if there was an error.
func (e GenericSwaggerError) Error() string {
	return fmt.Sprintf("error: %s, body: %s", e.error, e.body)
}

// Body returns the raw bytes of the response
func (e GenericSwaggerError) Body() string {
	return string(e.body)
}

// Model returns the unpacked model of the error
func (e GenericSwaggerError) Model() interface{} {
	return e.model
}
