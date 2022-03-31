package model

type TaskExecuteFunction func(t *Task) (*TaskResult, error)
