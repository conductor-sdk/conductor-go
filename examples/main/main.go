package main

import (
	"os"

	"github.com/conductor-sdk/conductor-go/examples/task_execute_function"
	"github.com/conductor-sdk/conductor-go/pkg/metrics"
	"github.com/conductor-sdk/conductor-go/pkg/orkestrator"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	log "github.com/sirupsen/logrus"
)

// Example init function that shows how to configure logging
// Using json formatter and changing level to Debug
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	//Stdout, change to a file for production use case
	log.SetOutput(os.Stdout)
	// Set to debug for demonstration.  Change to Info for production use cases.
	log.SetLevel(log.InfoLevel)
}

// Example main function to start workers
func main() {
	// MetricsSettings is optional,
	// could be nil instead and use default settings
	go metrics.ProvideMetrics(nil)
	// AuthenticationSettings and HttpSettings are optional,
	// could be nil instead and use default settings
	workerOrkestrator := orkestrator.NewWorkerOrkestrator(
		nil,
		settings.NewHttpSettings(
			"http://localhost:8080/api", // conductor http server url
			nil,
		),
	)

	THREAD_LIMIT := 10
	for i := 0; i < THREAD_LIMIT; i++ {
		workerOrkestrator.StartWorker(
			"task1",
			task_execute_function.Example1,
			10,
			10,
		)

	}

	// Wait for all workers to finish, otherwise would terminate them
	workerOrkestrator.WaitWorkers()
}
