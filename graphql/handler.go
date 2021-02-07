package graphql

import (
	"context"
	"net/http"
	"time"

	gql "github.com/99designs/gqlgen/graphql"
	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"xorm.io/xorm"

	apiErr "github.com/octo-5/karrot-api/errors"
	exec "github.com/octo-5/karrot-api/graphql/executable"
	"github.com/octo-5/karrot-api/graphql/resolver"
	"github.com/octo-5/karrot-api/model"
)

func Playground(title, gqlPath string) gin.HandlerFunc {
	graphqlUI := playground.Handler(title, gqlPath)
	return func(ctx *gin.Context) {
		graphqlUI(ctx.Writer, ctx.Request)
	}
}

func Handler(db *xorm.Engine) gin.HandlerFunc {
	cfg := exec.Config{
		Resolvers: &resolver.Resolver{},
	}

	cfg.Directives.IsAuth = IsAuth
	cfg.Directives.IsNotAuth = IsNotAuth
	cfg.Directives.HasRole = HasRole

	cfg.Complexity.Query.Users = pagingComplexity
	cfg.Complexity.Query.Todos = pagingComplexity
	cfg.Complexity.User.Todos = pagingComplexity

	schema := exec.NewExecutableSchema(cfg)

	srv := gqlHandler.New(schema)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.FixedComplexityLimit(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	srv.SetErrorPresenter(errorPresenter)

	return func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func pagingComplexity(childComplexity int, input model.PagingInput) int {
	return childComplexity * input.After
}

func errorPresenter(ctx context.Context, err error) *gqlerror.Error {
	gqlErr := gql.DefaultErrorPresenter(ctx, err)

	/*
		if errors.Is(cause, sql.ErrNoRows) {
			gqlErr.Message = "not found"
			gqlErr.Extensions = map[string]interface{}{"code": http.StatusNotFound}
			return gqlErr
		}

		if e, ok := cause.(*pq.Error); ok {
			switch e.Code {
			case "23505":
				gqlErr.Message = "duplicated value"
				gqlErr.Extensions = map[string]interface{}{"code": http.StatusConflict}
			}
			return gqlErr
		}
	*/

	/*
		if e, ok := err.(*apiErr.InternalServerError); ok {
			gqlErr.Message = e.Message()
			gqlErr.Extensions = e.Extensions()
			return gqlErr
		}
	*/
	if _, ok := err.(*gqlerror.Error); ok {
		return gqlErr
	} else if e, ok := err.(*apiErr.UnauthorizedError); ok {
		gqlErr.Message = e.Message()
		gqlErr.Extensions = e.Extensions()
	} else if e, ok := err.(*apiErr.BadRequestError); ok {
		gqlErr.Message = e.Message()
		gqlErr.Extensions = e.Extensions()
	} else {
		log.WithError(err).Error("unexpected error occurred")
		gqlErr.Message = "internal server error"
		gqlErr.Extensions = map[string]interface{}{"code": http.StatusInternalServerError}
	}
	return gqlErr
}
