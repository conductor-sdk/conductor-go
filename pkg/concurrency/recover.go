package concurrency

import (
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metrics_counter"
	"github.com/sirupsen/logrus"
)

func HandlePanicError(message string) {
	err := recover()
	if err == nil {
		return
	}
	metrics_counter.IncrementUncaughtException(message)
	logrus.Warning(
		"Uncaught panic",
		", message: ", message,
		", error: ", err,
	)
}
