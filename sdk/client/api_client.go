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
	"github.com/conductor-sdk/conductor-go/sdk/log"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/authentication"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

const (
	CONDUCTOR_AUTH_KEY            = "CONDUCTOR_AUTH_KEY"
	CONDUCTOR_AUTH_SECRET         = "CONDUCTOR_AUTH_SECRET"
	CONDUCTOR_SERVER_URL          = "CONDUCTOR_SERVER_URL"
	CONDUCTOR_CLIENT_HTTP_TIMEOUT = "CONDUCTOR_CLIENT_HTTP_TIMEOUT"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

type APIClient struct {
	httpRequester *HttpRequester
}

type Option func(c *http.Client)

// WithHTTPClient lets callers replace the entire http.Client.
func WithHTTPClient(custom *http.Client) Option {
	return func(c *http.Client) { *c = *custom }
}

// WithRoundTripper lets callers replace or wrap just Transport.
func WithRoundTripper(rt http.RoundTripper) Option {
	return func(c *http.Client) { c.Transport = rt }
}

func NewAPIClient(
	authenticationSettings *settings.AuthenticationSettings,
	httpSettings *settings.HttpSettings,
	opts ...Option,
) *APIClient {
	return newAPIClient(
		authenticationSettings,
		httpSettings,
		nil,
		nil,
		opts...,
	)
}
func NewAPIClientFromEnv() *APIClient {
	return NewAPIClient(NewAuthenticationSettingsFromEnv(), NewHttpSettingsFromEnv())
}

func NewAuthenticationSettingsFromEnv() *settings.AuthenticationSettings {
	return settings.NewAuthenticationSettings(
		os.Getenv(CONDUCTOR_AUTH_KEY),
		os.Getenv(CONDUCTOR_AUTH_SECRET),
	)
}

func NewHttpSettingsFromEnv() *settings.HttpSettings {
	url := os.Getenv(CONDUCTOR_SERVER_URL)
	if url == "" {
		log.Fatalf("Error: %s env variable is not set", CONDUCTOR_SERVER_URL)
	}

	return settings.NewHttpSettings(url)
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

func newAPIClient(authenticationSettings *settings.AuthenticationSettings, httpSettings *settings.HttpSettings, tokenExpiration *authentication.TokenExpiration, tokenManager authentication.TokenManager, opts ...Option) *APIClient {
	if httpSettings == nil {
		httpSettings = settings.NewHttpDefaultSettings()
	}
	var httpTimeout = 30 * time.Second // Set default value once

	timeoutStr := os.Getenv(CONDUCTOR_CLIENT_HTTP_TIMEOUT)
	if timeoutStr != "" {
		// Only try to parse if the environment variable is actually set
		if timeoutInt, err := strconv.Atoi(timeoutStr); err == nil {
			httpTimeout = time.Duration(timeoutInt) * time.Second
		}
		// If parsing fails, we'll keep the default value
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
	client := &http.Client{
		Transport:     netTransport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       httpTimeout,
	}
	for _, opt := range opts {
		opt(client)
	}
	return &APIClient{
		httpRequester: NewHttpRequester(
			authenticationSettings, httpSettings, client, tokenExpiration, tokenManager,
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
	bodyBuf = &bytes.Buffer{}

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
		err = xml.NewEncoder(bodyBuf).Encode(body)
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
			log.Error("Unable to decompress the response ", err.Error())
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

// executeCall performs an HTTP request with centralized error handling
// Supports all CRUD operations through a common interface
func (c *APIClient) executeCall(ctx context.Context, method, path string, queryParams url.Values, body interface{}, contentType string, result interface{}) (*http.Response, error) {
	// Create headers
	headers := make(map[string]string)

	// Set content type if body is provided
	if body != nil {
		cType := "application/json"
		if len(contentType) > 0 && contentType != "" {
			cType = contentType
		}
		headers["Content-Type"] = cType
	}

	// Set accept header for all requests
	headers["Accept"] = "application/json"

	// Prepare the request
	req, err := c.prepareRequest(ctx, path, method, body, headers, queryParams, nil, "", nil)
	if err != nil {
		return nil, err
	}

	// Call the API
	resp, err := c.callAPI(req)
	if err != nil || resp == nil {
		return resp, err
	}

	// Get response body
	respBody, err := getDecompressedBody(resp)
	if err != nil {
		return resp, err
	}

	// Handle successful response
	if isSuccessfulStatus(resp.StatusCode) {
		if result != nil && len(respBody) > 0 {
			err = c.decode(result, respBody, resp.Header.Get("Content-Type"))
		}
		return resp, err
	}

	// Handle error response - create GenericSwaggerError with status code
	newErr := NewGenericSwaggerError(respBody, string(respBody), nil, resp.StatusCode)
	return resp, newErr
}

// Get performs a GET request
func (c *APIClient) Get(ctx context.Context, path string, queryParams url.Values, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "GET", path, queryParams, nil, "", result)
}

// Post performs a POST request
func (c *APIClient) Post(ctx context.Context, path string, body interface{}, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "POST", path, nil, body, "", result)
}

// PostWithContentType performs post with given content type
func (c *APIClient) PostWithContentType(ctx context.Context, path string, body interface{}, contentType string, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "POST", path, nil, body, contentType, result)
}

// PostWithParams performs a POST request with query parameters
func (c *APIClient) PostWithParams(ctx context.Context, path string, queryParams url.Values, body interface{}, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "POST", path, queryParams, body, "", result)
}

// Put performs a PUT request
func (c *APIClient) Put(ctx context.Context, path string, body interface{}, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "PUT", path, nil, body, "", result)
}

// PutWithContentType performs a PUT request
func (c *APIClient) PutWithContentType(ctx context.Context, path string, body interface{}, contentType string, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "PUT", path, nil, body, contentType, result)
}

// PutWithParams performs a PUT request with query parameters
func (c *APIClient) PutWithParams(ctx context.Context, path string, queryParams url.Values, body interface{}, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "PUT", path, queryParams, body, "", result)
}

// Delete performs a DELETE request without a body
func (c *APIClient) Delete(ctx context.Context, path string, queryParams url.Values, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "DELETE", path, queryParams, nil, "", result)
}

// DeleteWithBody performs a DELETE request with a body
func (c *APIClient) DeleteWithBody(ctx context.Context, path string, body interface{}, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "DELETE", path, nil, body, "", result)
}

// Patch performs a PATCH request
func (c *APIClient) Patch(ctx context.Context, path string, body interface{}, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "PATCH", path, nil, body, "", result)
}
