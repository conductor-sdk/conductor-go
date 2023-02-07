//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package authentication

import "time"

// sets the default expiration time for each generated token and a cleanupInterval to old delete entries
type TokenExpiration struct {
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

func NewTokenExpiration(defaultExpiration time.Duration, cleanupInterval time.Duration) *TokenExpiration {
	return &TokenExpiration{
		DefaultExpiration: defaultExpiration,
		CleanupInterval:   cleanupInterval,
	}
}

func NewDefaultTokenExpiration() *TokenExpiration {
	return &TokenExpiration{
		DefaultExpiration: 30 * time.Minute,
		CleanupInterval:   2 * time.Hour,
	}
}
