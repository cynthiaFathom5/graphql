package graphgo generate ./...

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/FATHOM5/graphql/b-service/graph/generated"
	"github.com/FATHOM5/graphql/common/database"
)

func (r *queryResolver) Users(ctx context.Context) ([]*database.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
