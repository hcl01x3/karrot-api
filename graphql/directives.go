package graphql

import (
	"context"

	gql "github.com/99designs/gqlgen/graphql"

	apiErr "github.com/octo-5/karrot-api/errors"
	mid "github.com/octo-5/karrot-api/middleware"
	"github.com/octo-5/karrot-api/model"
)

func IsAuth(ctx context.Context, obj interface{}, next gql.Resolver) (interface{}, error) {
	if mid.GetUser(ctx) == nil {
		return nil, apiErr.Unauthrozied("current user isn't authenticated.")
	}
	return next(ctx)
}

func IsNotAuth(ctx context.Context, obj interface{}, next gql.Resolver) (interface{}, error) {
	if mid.GetUser(ctx) == nil {
		//TODO
		return nil, apiErr.Unauthrozied("current user isn't authenticated.")
	}
	return next(ctx)
}

func HasRole(ctx context.Context, obj interface{}, next gql.Resolver, role model.Role) (interface{}, error) {
	currUser := mid.GetUser(ctx)
	if currUser == nil || currUser.Role != role {
		return nil, apiErr.Unauthrozied("current user hasn't authorized to perform this action.")
	}
	return next(ctx)
}
