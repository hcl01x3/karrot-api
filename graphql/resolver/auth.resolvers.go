package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/octo-5/karrot-api/graphql/executable"
	"github.com/octo-5/karrot-api/model"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LogInInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns executable.MutationResolver implementation.
func (r *Resolver) Mutation() executable.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
