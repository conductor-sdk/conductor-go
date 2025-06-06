//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

import (
	"encoding/json"
	"fmt"
)

type BulkResponse struct {
	BulkErrorResults map[string]string `json:"bulkErrorResults,omitempty"`

	// Keep for backward compatibility but mark as deprecated
	BulkSuccessfulResults []string `json:"bulkSuccessfulResults,omitempty"` // Deprecated: Use GetSuccessfulResults() instead

	// Internal field for new functionality
	bulkSuccessfulResultsRaw interface{} `json:"-"`

	Message string `json:"message,omitempty"`
}

// Custom unmarshaling
func (br *BulkResponse) UnmarshalJSON(data []byte) error {
	type Alias BulkResponse
	aux := &struct {
		BulkSuccessfulResults json.RawMessage `json:"bulkSuccessfulResults,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(br),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Store raw data
	if len(aux.BulkSuccessfulResults) > 0 {
		// Try []string first
		var stringResults []string
		if err := json.Unmarshal(aux.BulkSuccessfulResults, &stringResults); err == nil {
			br.BulkSuccessfulResults = stringResults
			br.bulkSuccessfulResultsRaw = stringResults
			return nil
		}

		// Try []interface{}
		var interfaceResults []interface{}
		if err := json.Unmarshal(aux.BulkSuccessfulResults, &interfaceResults); err == nil {
			br.bulkSuccessfulResultsRaw = interfaceResults
			// Convert to strings for backward compatibility
			br.BulkSuccessfulResults = make([]string, len(interfaceResults))
			for i, item := range interfaceResults {
				br.BulkSuccessfulResults[i] = fmt.Sprintf("%v", item)
			}
			return nil
		}
	}

	return nil
}

// New recommended methods
func (br *BulkResponse) GetSuccessfulResults() interface{} {
	return br.bulkSuccessfulResultsRaw
}

func (br *BulkResponse) GetSuccessfulResultsAsStrings() []string {
	return br.BulkSuccessfulResults
}
