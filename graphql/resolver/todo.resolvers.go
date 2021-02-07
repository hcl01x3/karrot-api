package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/octo-5/karrot-api/graphql/executable"
	"github.com/octo-5/karrot-api/model"
)

func (r *mutationResolver) NewTodo(ctx context.Context, input model.NewTodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context, input model.PagingInput) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todo(ctx context.Context, id int64) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) Author(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns executable.QueryResolver implementation.
func (r *Resolver) Query() executable.QueryResolver { return &queryResolver{r} }

// Todo returns executable.TodoResolver implementation.
func (r *Resolver) Todo() executable.TodoResolver { return &todoResolver{r} }

type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
