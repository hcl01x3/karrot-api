package middleware

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/octo-5/karrot-api/config"
)

var configKey = ContextKey("configKey")

func MustGetConfig(ctx context.Context) *config.Config {
	val := MustGet(ctx, configKey)
	cfg, ok := val.(*config.Config)
	if !ok {
		panic("")
	}
	return cfg
}

func Config(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), configKey, cfg)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
