package graphql

import "context"

//
type mutationResolver struct {
	*rootResolver
}

func (r *mutationResolver) Test(ctx context.Context) (string, error) {
	return "0", nil
}
