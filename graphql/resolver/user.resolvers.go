package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/octo-5/karrot-api/graphql/executable"
	"github.com/octo-5/karrot-api/model"
)

func (r *mutationResolver) NewUser(ctx context.Context, input model.NewUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePassword(ctx context.Context, input model.UpdatePasswordInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id int64) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, input model.PagingInput) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User, input model.PagingInput) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns executable.UserResolver implementation.
func (r *Resolver) User() executable.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
