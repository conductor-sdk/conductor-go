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
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/authentication"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/sirupsen/logrus"
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

func newAPIClient(authenticationSettings *settings.AuthenticationSettings, httpSettings *settings.HttpSettings, tokenExpiration *authentication.TokenExpiration, tokenManager authentication.TokenManager) *APIClient {
	if httpSettings == nil {
		httpSettings = settings.NewHttpDefaultSettings()
	}
	baseDialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	netTransport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		DialContext:         baseDialer.DialContext,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		DisableCompression:  false,
	}
	client := http.Client{
		Transport:     netTransport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       30 * time.Second,
	}
	return &APIClient{
		httpRequester: NewHttpRequester(
			authenticationSettings, httpSettings, &client, tokenExpiration, tokenManager,
		),
	}
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.httpRequester.httpClient.Do(request)
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}

	if strings.Contains(contentType, "application/xml") {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	} else if strings.Contains(contentType, "application/json") {
		if err = json.Unmarshal(b, v); err != nil {
			// Hacky - if json unmarshalling fails, return a string.
			// it's because the backend might reply with content-type: application/json and a string.
			rv := reflect.ValueOf(v)
			if rv.Kind() == reflect.Ptr && rv.Elem().Kind() == reflect.String {
				rv.Elem().SetString(string(b))
				return nil
			}
			return err
		}
		return nil
	} else if strings.Contains(contentType, "text/plain") {
		rv := reflect.ValueOf(v)
		if rv.IsNil() {
			return errors.New("undefined response type")
		}
		rv.Elem().SetString(string(b))
		return nil
	}

	return errors.New("undefined response type")
}

func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte,
) (localVarRequest *http.Request, err error) {
	return c.httpRequester.prepareRequest(
		ctx, path, method, postBody, headerParams, queryParams, formParams, fileName, fileBytes,
	)
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
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

func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}
	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

func getDecompressedBody(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	var reader io.ReadCloser
	var err error
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			logrus.Error("Unable to decompress the response ", err.Error())
			if err == io.EOF {
				return nil, nil
			}
			return nil, err
		}
	default:
		reader = response.Body
	}
	defer reader.Close()
	return io.ReadAll(reader)
}

func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

func isSuccessfulStatus(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
