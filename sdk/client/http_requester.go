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
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/conductor-sdk/conductor-go/sdk/authentication"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

type HttpRequester struct {
	httpSettings *settings.HttpSettings
	httpClient   *http.Client
	tokenManager authentication.TokenManager
}

func NewHttpRequester(authenticationSettings *settings.AuthenticationSettings, httpSettings *settings.HttpSettings, httpClient *http.Client, tokenExpiration *authentication.TokenExpiration, tokenManager authentication.TokenManager) *HttpRequester {

	if authenticationSettings != nil && !authenticationSettings.IsEmpty() {
		if tokenManager == nil {
			tokenManager = nil
			tokenManager = authentication.NewTokenManager(*authenticationSettings, tokenExpiration)
		}
	}
	return &HttpRequester{
		httpSettings: httpSettings,
		httpClient:   httpClient,
		tokenManager: tokenManager,
	}
}

// prepareRequest build the request
func (h *HttpRequester) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("cannot specify postBody and multipart form at the same time")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile("file", filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("cannot specify postBody and x-www-form-urlencoded form at the same time")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(h.httpSettings.BaseUrl + path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	if body != nil {
		localVarRequest, err = http.NewRequestWithContext(ctx, method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequestWithContext(ctx, method, url.String(), nil)
	}

	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	for header, value := range h.httpSettings.Headers {
		if header == "Content-Type" {
			// a client can only produce content of one single type, so skip adding the Content-Type header
			// when it already exists
			if _, ok := localVarRequest.Header["Content-Type"]; ok {
				continue
			}
		}
		localVarRequest.Header.Add(header, value)
	}

	if h.tokenManager != nil {
		token, err := h.tokenManager.RefreshToken(h.httpSettings, h.httpClient)
		if err == nil {
			localVarRequest.Header.Add("X-Authorization", token)
		}
	}

	return localVarRequest, nil
}
