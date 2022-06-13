package settings

import (
	"github.com/conductor-sdk/conductor-go/model"
)

type ExternalStorageSettings struct {
	TaskOutputPayloadThresholdKB    int64
	TaskOutputMaxPayloadThresholdKB int64
	ExternalStorageHandler          model.ExternalStorageHandler
}

func NewExternalStorageSettings(
	taskOutputPayloadThresholdKB int64,
	taskOutputMaxPayloadThresholdKB int64,
	externalStorageHandler model.ExternalStorageHandler,
) *ExternalStorageSettings {
	return &ExternalStorageSettings{
		TaskOutputPayloadThresholdKB:    taskOutputPayloadThresholdKB,
		TaskOutputMaxPayloadThresholdKB: taskOutputMaxPayloadThresholdKB,
		ExternalStorageHandler:          externalStorageHandler,
	}
}
