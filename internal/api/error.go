package api

type Error struct {
	StatusCode int
	Err        error
	Message    any
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func NewApiError(statusCode int, err error) *Error {
	return &Error{
		StatusCode: statusCode,
		Err:        err,
	}
}
