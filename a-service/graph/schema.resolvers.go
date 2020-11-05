package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/FATHOM5/graphql/a-service/graph/generated"
	"github.com/FATHOM5/graphql/common/database"
)

func (r *queryResolver) Node(ctx context.Context, id string) (database.Node, error) {
	return database.Users[id], nil
}

func (r *queryResolver) User(ctx context.Context, name string) (*database.User, error) {
	u := database.Users[name]
	return &u, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*database.User, error) {
	return database.AllUsersList(), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
