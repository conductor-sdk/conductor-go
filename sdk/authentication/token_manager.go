//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package authentication

import (
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

const (
	tokenKey = "TOKEN_KEY"
)

type TokenManager struct {
	mutex       sync.RWMutex
	credentials settings.AuthenticationSettings
	database    cache.Cache
}

func NewTokenManager(credentials settings.AuthenticationSettings, defaultExpiration time.Duration, cleanupInterval time.Duration) *TokenManager {
	return &TokenManager{
		credentials: credentials,
		database:    *cache.New(defaultExpiration, cleanupInterval),
	}
}

func (t *TokenManager) RefreshToken(httpSettings *settings.HttpSettings, httpClient *http.Client) (string, error) {
	token := t.getTokenIfCached()
	if token != "" {
		return token, nil
	}
	return t.refreshToken(httpSettings, httpClient)
}

func (t *TokenManager) getTokenIfCached() string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	token, found := t.database.Get(tokenKey)
	if found {
		return token.(string)
	}
	return ""
}

func (t *TokenManager) refreshToken(httpSettings *settings.HttpSettings, httpClient *http.Client) (string, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	log.Debug("Refreshing authentication token")
	token, response, err := getToken(t.credentials, httpSettings, httpClient)
	if err != nil {
		log.Warning(
			"Failed to refresh authentication token",
			", response: ", response,
			", error: ", err,
		)
		t.database.Delete(tokenKey)
		return "", err
	}
	log.Debug("Refreshed authentication token")
	t.database.Set(tokenKey, token.Token, cache.DefaultExpiration)
	return token.Token, nil
}
