package middleware

import (
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RecoverPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				headers := strings.Split(string(httpRequest), "\r\n")

				for idx, header := range headers {
					current := strings.Split(header, ":")
					if current[0] == "Authorization" {
						headers[idx] = current[0] + ": *"
					}
				}

				headersToStr := strings.Join(headers, "\r\n")

				log.WithFields(
					log.Fields{
						"request": headersToStr,
						"stack":   string(debug.Stack()),
						"error":   err,
					},
				).Error(err)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
