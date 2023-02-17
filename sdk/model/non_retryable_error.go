package model

type NonRetryableError struct {
	error
}

func NewNonRetryableError(err error) *NonRetryableError {
	return &NonRetryableError{
		error: err,
	}
}

func (e *NonRetryableError) Error() string {
	return e.error.Error()
}
