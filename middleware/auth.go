package middleware

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"xorm.io/xorm"

	"github.com/octo-5/karrot-api/model"
	"github.com/octo-5/karrot-api/util"
)

var (
	authUserKey  = ContextKey("authUserKey")
	bearerRegexp = regexp.MustCompile(`^\s*Bearer\s+`)
)

func MustGetUser(ctx context.Context) *model.User {
	return MustGet(ctx, authUserKey).(*model.User)
}

func GetUser(ctx context.Context) *model.User {
	return ctx.Value(authUserKey).(*model.User)
}

func AuthUser(db *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		cfg := MustGetConfig(ctx)

		var user *model.User

		accessToken := getAuthorization(c)

		if accessToken != "" {
			token, err := util.DecodeJWT(accessToken, cfg.JWTSecret)
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					log.WithError(err).Error("invalid signature in jwt")
					c.AbortWithStatus(500)
				}
				c.AbortWithStatus(400)
			}

			if token != nil && token.Valid() == nil {
				target := model.User{}
				found, err := db.Context(ctx).ID(token.UserId).Get(&target)
				if err != nil {
					log.WithError(err).Error("auth middleware: get user query error")
					c.AbortWithStatus(500)
				}
				if found {
					user = &target
				}
			}
		}

		ctx = context.WithValue(ctx, authUserKey, user)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func getAuthorization(c *gin.Context) string {
	val := c.Request.Header.Get("Authorization")
	return bearerRegexp.ReplaceAllString(val, "")
}

func getCookie(c *gin.Context, name string) string {
	token, err := c.Cookie(name)
	if errors.Is(err, http.ErrNoCookie) {
		return ""
	}
	return strings.TrimSpace(token)
}

func setCookie(c *gin.Context, name, val string, maxAge int) {
	cfg := MustGetConfig(c.Request.Context())
	c.SetCookie(name, val, maxAge, "/", cfg.APIDomain, true, true)
}
