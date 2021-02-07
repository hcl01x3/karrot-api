package errors

import (
	"net/http"
)

var _ APIError = (*BadRequestError)(nil)

type BadRequestError struct {
	msg  string
	code int
	errs []error
}

func BadRequest(errs []error) *BadRequestError {
	code := http.StatusBadRequest
	return &BadRequestError{
		msg:  http.StatusText(code),
		code: code,
		errs: errs,
	}
}

func (e *BadRequestError) Unwrap() error {
	return nil
}

func (e *BadRequestError) Error() string {
	return e.msg
}

func (e *BadRequestError) Code() int {
	return e.code
}

func (e *BadRequestError) Message() string {
	return e.msg
}

func (e *BadRequestError) Extensions() map[string]interface{} {
	fieldErrors := []string{}
	for _, err := range e.errs {
		fieldErrors = append(fieldErrors, err.Error())
	}
	return map[string]interface{}{"code": e.code, "fieldsErrors": fieldErrors}
}
