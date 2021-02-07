package middleware

import (
	"math"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RequestLogging(skipPaths ...string) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	isContains := func(target []string, value string) bool {
		for _, v := range target {
			if v == value {
				return true
			}
		}
		return false
	}

	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path

		if isContains(skipPaths, path) {
			ctx.Next()
			return
		}

		start := time.Now()
		ctx.Next()
		stop := time.Since(start)

		statusCode := ctx.Writer.Status()
		dataLength := ctx.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		logEntry := log.WithFields(log.Fields{
			"hostname":    hostname,
			"path":        path,
			"method":      ctx.Request.Method,
			"referer":     ctx.Request.Referer(),
			"statusCode":  statusCode,
			"dataLength":  dataLength,
			"latency":     int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0)),
			"clientIP":    ctx.ClientIP(),
			"clientAgent": ctx.Request.UserAgent(),
		})

		if len(ctx.Errors) > 0 {
			logEntry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := http.StatusText(statusCode)
			if statusCode >= http.StatusInternalServerError {
				logEntry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				logEntry.Warn(msg)
			} else {
				logEntry.Info(msg)
			}
		}
	}
}
