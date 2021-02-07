package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

var dbKey = ContextKey("dbKey")

func MustGetDB(ctx context.Context) *xorm.Engine {
	val := MustGet(ctx, dbKey)
	db, ok := val.(*xorm.Engine)
	if !ok {
		panic("")
	}
	return db
}

func Database(db *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), dbKey, db)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
