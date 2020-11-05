package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/FATHOM5/graphql/b-service/graph/generated"
	"github.com/FATHOM5/graphql/common/database"
)

func (r *queryResolver) Node(ctx context.Context, id string) (database.Node, error) {
	return database.Users[id], nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
