package model

type ExternalStorageHandler func(data map[string]interface{}) (string, error)
