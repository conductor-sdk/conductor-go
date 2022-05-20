package concurrency

import (
	log "github.com/sirupsen/logrus"
)

func GoSafe(callable func(...interface{}), args ...interface{}) {
	defer onError(callable, args...)
	callable(args...)
}

func onError(callable func(...interface{}), args ...interface{}) {
	if err := recover(); err != nil {
		log.Error(
			"Uncaught error, with function: ", callable,
			", args: ", args,
			", error: ", err,
		)
	}
}
