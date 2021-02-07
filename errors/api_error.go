package errors

type APIError interface {
	Error() string
	Unwrap() error
	Code() int
	Message() string
	Extensions() map[string]interface{}
}
