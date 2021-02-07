package errors

import (
	"net/http"
)

var _ APIError = (*NotFoundError)(nil)

type NotFoundError struct {
	msg  string
	code int
}

func NotFound(msg string) *NotFoundError {
	code := http.StatusNotFound
	if msg == "" {
		msg = http.StatusText(code)
	}
	return &NotFoundError{
		msg:  msg,
		code: code,
	}
}

func (e *NotFoundError) Unwrap() error {
	return nil
}

func (e *NotFoundError) Error() string {
	return e.msg
}

func (e *NotFoundError) Code() int {
	return e.code
}

func (e *NotFoundError) Message() string {
	return e.msg
}

func (e *NotFoundError) Extensions() map[string]interface{} {
	return map[string]interface{}{"code": e.code}
}
