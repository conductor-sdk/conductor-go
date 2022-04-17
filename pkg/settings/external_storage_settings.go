package settings

type ExternalStorageSettings struct {
	TaskOutputPayloadThresholdKB    int64
	TaskOutputMaxPayloadThresholdKB int64
}

func NewExternalStorageSettings(
	taskOutputPayloadThresholdKB int64,
	taskOutputMaxPayloadThresholdKB int64,
) *ExternalStorageSettings {
	return &ExternalStorageSettings{
		TaskOutputPayloadThresholdKB:    taskOutputPayloadThresholdKB,
		TaskOutputMaxPayloadThresholdKB: taskOutputMaxPayloadThresholdKB,
	}
}
