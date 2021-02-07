package middleware

import (
	"context"
)

type ContextKey string

func MustGet(ctx context.Context, key ContextKey) interface{} {
	val := ctx.Value(key)
	if val == nil {
		panic("")
	}
	return val
}
