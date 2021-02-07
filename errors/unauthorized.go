package errors

import (
	"net/http"
)

var _ APIError = (*UnauthorizedError)(nil)

type UnauthorizedError struct {
	msg  string
	code int
}

func Unauthrozied(msg string) *UnauthorizedError {
	code := http.StatusUnauthorized
	if msg == "" {
		msg = http.StatusText(code)
	}
	return &UnauthorizedError{
		code: code,
		msg:  msg,
	}
}

func (e *UnauthorizedError) Unwrap() error {
	return nil
}

func (e *UnauthorizedError) Error() string {
	return e.msg
}

func (e *UnauthorizedError) Code() int {
	return e.code
}

func (e *UnauthorizedError) Message() string {
	return e.msg
}

func (e *UnauthorizedError) Extensions() map[string]interface{} {
	return map[string]interface{}{"code": e.code}
}
