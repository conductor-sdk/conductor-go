package util

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	maxRetryAttempts int = 3
)

func RetryFunction(doFunc func(*http.Request) (*http.Response, error), request *http.Request) (*http.Response, error) {
	var err error
	var response *http.Response
	for attempt := 0; attempt < maxRetryAttempts; attempt++ {
		if attempt > 0 {
			// Wait for [10s, 20s, 30s] before next attempt
			amount := attempt * 10
			time.Sleep(time.Duration(amount) * time.Second)
		}
		response, err = doFunc(request)
		if err == nil {
			return response, nil
		}
		log.Debug(
			"Failed to make request",
			", reason: ", err.Error(),
		)
	}
	return nil, err
}
