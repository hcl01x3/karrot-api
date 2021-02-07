package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var cacheKey = ContextKey("cacheKey")

func MustGetCache(ctx context.Context) redis.Cmdable {
	val := MustGet(ctx, cacheKey)
	cache, ok := val.(redis.Cmdable)
	if !ok {
		panic("")
	}
	return cache
}

func Cache(cache redis.Cmdable) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), cacheKey, cache)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
