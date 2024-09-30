//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package client

import (
	"net"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/authentication"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

const (
	CONDUCTOR_AUTH_KEY    = "CONDUCTOR_AUTH_KEY"
	CONDUCTOR_AUTH_SECRET = "CONDUCTOR_AUTH_SECRET"
	CONDUCTOR_SERVER_URL  = "CONDUCTOR_SERVER_URL"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

type APIClient struct {
	dialer        *net.Dialer
	netTransport  *http.Transport
	httpClient    *http.Client
	httpRequester *HttpRequester
	tokenManager  authentication.TokenManager
}

func NewAPIClient(
	authenticationSettings *settings.AuthenticationSettings,
	httpSettings *settings.HttpSettings,
) *APIClient {
	return newAPIClient(
		authenticationSettings,
		httpSettings,
		nil,
		nil,
	)
}
func NewAPIClientFromEnv() *APIClient {
	authenticationSettings := settings.NewAuthenticationSettings(os.Getenv(CONDUCTOR_AUTH_KEY), os.Getenv(CONDUCTOR_AUTH_SECRET))
	httpSettings := settings.NewHttpSettings(os.Getenv(CONDUCTOR_SERVER_URL))
	return NewAPIClient(authenticationSettings, httpSettings)
}

func NewAPIClientWithTokenExpiration(
	authenticationSettings *settings.AuthenticationSettings,
	httpSettings *settings.HttpSettings,
	tokenExpiration *authentication.TokenExpiration,
) *APIClient {
	return newAPIClient(
		authenticationSettings,
		httpSettings,
		tokenExpiration,
		nil,
	)
}

func NewAPIClientWithTokenManager(
	authenticationSettings *settings.AuthenticationSettings,
	httpSettings *settings.HttpSettings,
	tokenExpiration *authentication.TokenExpiration,
	tokenManager authentication.TokenManager,
) *APIClient {
	return newAPIClient(
		authenticationSettings,
		httpSettings,
		tokenExpiration,
		tokenManager,
	)
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func (client *APIClient) ConfigureDialer(configurer func(dialer *net.Dialer)) *APIClient {
	configurer(client.dialer)

	return client
}

func (client *APIClient) ConfigureTransport(configurer func(transport *http.Transport)) *APIClient {
	configurer(client.netTransport)

	return client
}

func (client *APIClient) ConfigureHttpClient(configurer func(client *http.Client)) *APIClient {
	configurer(client.httpClient)

	return client
}
