package concurrency

import (
	"github.com/conductor-sdk/conductor-go/metrics"
	log "github.com/sirupsen/logrus"
)

func HandlePanicError(message string) {
	err := recover()
	if err == nil {
		return
	}
	metrics.IncrementUncaughtException(message)
	log.Warning(
		"Uncaught panic",
		", message: ", message,
		", error: ", err,
	)
}
