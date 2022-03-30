package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/conductor-sdk/conductor/client/go/conductor_client/model"
	"github.com/conductor-sdk/conductor/client/go/settings"
	log "github.com/sirupsen/logrus"
)

type HttpClient struct {
	authenticationSettings *settings.AuthenticationSettings
	client                 *http.Client
	httpSettings           *settings.HttpSettings
}

func NewHttpClientWithAuthentication(authenticationSettings *settings.AuthenticationSettings, httpSettings *settings.HttpSettings) *HttpClient {
	httpClient := new(HttpClient)
	httpClient.authenticationSettings = authenticationSettings
	httpClient.client = &http.Client{}
	if httpSettings == nil {
		httpSettings = settings.NewHttpDefaultSettings()
	}
	httpClient.httpSettings = httpSettings
	return httpClient
}

func (c *HttpClient) logSendRequest(url string, requestType string, body string) {
	log.Debug("Sending [", requestType, "] request to Server (", url, "):")
	log.Debug("Body:")
	log.Debug(body)
}

func (c *HttpClient) logResponse(statusCode string, response string) {
	log.Debug("Received response from Server (", c.httpSettings.BaseUrl, "):")
	log.Debug("Status: ", statusCode)
	log.Debug("Response:")
	log.Debug(response)
}

func genParamString(paramMap map[string]string) string {
	if paramMap == nil || len(paramMap) == 0 {
		return ""
	}

	output := "?"
	for key, value := range paramMap {
		output += key
		output += "="
		output += value
		output += "&"
	}
	return output
}

func (c *HttpClient) httpRequest(url string, requestType string, headers map[string]string, body string) (string, error) {
	var req *http.Request
	var err error

	if requestType == "GET" {
		req, err = http.NewRequest(requestType, url, nil)
	} else {
		var bodyStr = []byte(body)
		req, err = http.NewRequest(requestType, url, bytes.NewBuffer(bodyStr))
	}
	if err != nil {
		return "", err
	}

	// Default Headers
	for key, value := range c.httpSettings.Headers {
		req.Header.Set(key, value)
	}

	// Custom Headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	c.logSendRequest(url, requestType, body)

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	// If successful HTTP call, but Client/Server error, we return error
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return "", fmt.Errorf("%d Http Client Error for url: %s", resp.StatusCode, url)
	}
	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		return "", fmt.Errorf("%d Http Server Error for url: %s", resp.StatusCode, url)
	}

	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	responseString := string(response)
	if err != nil {
		log.Error("ERROR reading response for URL: ", url, err)
		return "", err
	}
	c.logResponse(resp.Status, responseString)

	return responseString, nil
}

func (c *HttpClient) Get(url string, queryParamsMap map[string]string, headers map[string]string) (string, error) {
	urlString := url + genParamString(queryParamsMap)
	resp, err := c.httpRequest(urlString, "GET", headers, "")
	if err != nil {
		log.Debug("Http GET Error for URL: ", urlString)
		return "", err
	}
	return resp, nil
}

func (c *HttpClient) Put(url string, queryParamsMap map[string]string, headers map[string]string, body string) (string, error) {
	urlString := url + genParamString(queryParamsMap)
	resp, err := c.httpRequest(urlString, "PUT", headers, body)
	if err != nil {
		log.Error("Http PUT Error for URL: ", urlString, err)
		return "", err
	}
	return resp, nil
}

func (c *HttpClient) Post(url string, queryParamsMap map[string]string, headers map[string]string, body string) (string, error) {
	urlString := url + genParamString(queryParamsMap)
	resp, err := c.httpRequest(urlString, "POST", headers, body)
	if err != nil {
		log.Error("Http POST Error for URL: ", urlString, "Error: ", err)
		return "", err
	}
	return resp, nil
}

func (c *HttpClient) Delete(url string, queryParamsMap map[string]string, headers map[string]string, body string) (string, error) {
	urlString := url + genParamString(queryParamsMap)
	resp, err := c.httpRequest(urlString, "DELETE", headers, body)
	if err != nil {
		log.Error("Http DELETE Error for URL: ", urlString)
		return "", err
	}
	return resp, nil
}

func (c *HttpClient) MakeUrl(path string, args ...string) string {
	url := c.httpSettings.BaseUrl
	r := strings.NewReplacer(args...)
	return url + r.Replace(path)
}

/**********************/
/* Auth Functions */
/**********************/

func (c *HttpClient) RefreshToken() {
	if c.authenticationSettings == nil {
		return
	}
	url := c.MakeUrl("/token")
	body := c.authenticationSettings.GetFormattedSettings()
	resp, err := c.Post(
		url,
		nil,
		nil,
		body,
	)
	if err != nil {
		log.Error("Http RefreshToken: Failed to get token, error: ", err)
		return
	}
	token, err := model.GetTokenFromResponse(resp)
	if err != nil {
		log.Error("Http RefreshToken: Failed to parse token message, error: ", err)
	}
	c.updateToken(token.Token)
}

func (c *HttpClient) updateToken(token string) {
	c.httpSettings.Headers["X-Authorization"] = token
}
