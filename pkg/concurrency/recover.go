package concurrency

import (
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metrics_counter"
	log "github.com/sirupsen/logrus"
)

func OnError(message string) {
	if err := recover(); err != nil {
		metrics_counter.IncrementUncaughtException(message)
		log.Error(
			"Uncaught error",
			", message: ", message,
			", error: ", err,
		)
	}
}
