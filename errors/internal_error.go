package errors

import (
	"net/http"
)

var _ APIError = (*InternalServerError)(nil)

type InternalServerError struct {
	msg   string
	cause error
	code  int
	errs  []error
}

func InternalError(cause error) *InternalServerError {
	code := http.StatusInternalServerError
	return &InternalServerError{
		msg:   http.StatusText(code),
		code:  code,
		cause: cause,
	}
}

func (e *InternalServerError) Unwrap() error {
	return e.cause
}

func (e *InternalServerError) Error() string {
	return e.cause.Error()
}

func (e *InternalServerError) Code() int {
	return e.code
}

func (e *InternalServerError) Message() string {
	return e.msg
}

func (e *InternalServerError) Extensions() map[string]interface{} {
	return map[string]interface{}{"code": e.code}
}
