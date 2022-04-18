package main

import (
	"os"

	"github.com/conductor-sdk/conductor-go/examples/external_storage_handler"
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
	log.SetLevel(log.DebugLevel)
}

// Example main function to start workers
func main() {
	// MetricsSettings is optional,
	// could be nil instead and use default settings
	go metrics.ProvideMetrics(
		settings.NewMetricsSettings(
			"/metrics", // api endpoint
			2112,       // port
		),
	)
	// AuthenticationSettings and HttpSettings are optional,
	// could be nil instead and use default settings
	workerOrkestrator := orkestrator.NewWorkerOrkestrator(
		settings.NewAuthenticationSettings(
			"keyId",     // key id from your application
			"keySecret", // key secret from your application
		),
		settings.NewHttpSettings(
			"https://play.orkes.io", // conductor http server url
			settings.NewExternalStorageSettings(
				4,  // taskOutputPayloadThresholdKB
				10, // taskOutputMaxPayloadThresholdKB
				external_storage_handler.S3UploadAndGetPath, // ExternalStorageHandler function
			),
		),
	)
	workerOrkestrator.StartWorker(
		"go_task_example",              // task definition name
		task_execute_function.Example1, // task execution function
		1,                              // parallel go routines amount
		5000,                           // 5000ms
	)
	// Wait for all workers to finish, otherwise would terminate them
	workerOrkestrator.WaitWorkers()
}
