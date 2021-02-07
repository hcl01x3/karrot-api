package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	dl "github.com/graph-gophers/dataloader/v6"
	"xorm.io/xorm"

	"github.com/octo-5/karrot-api/datastore"
	"github.com/octo-5/karrot-api/datastore/batch"
)

var dataLoaderKey = ContextKey("dataLoaderKey")

func MustGetLoader(ctx context.Context) *datastore.DataLoader {
	val := MustGet(ctx, dataLoaderKey)
	loader, ok := val.(*datastore.DataLoader)
	if !ok {
		panic("")
	}
	return loader
}

func Dataloader(db *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		loader := &datastore.DataLoader{
			UserById: dl.NewBatchedLoader(batch.UserById(db)),
		}
		ctx := context.WithValue(c.Request.Context(), dataLoaderKey, loader)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
