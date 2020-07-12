package graphql

import "context"

// queryResolver resolvers query
type queryResolver struct {
	*rootResolver
}

func (r *queryResolver) Brands(ctx context.Context) (*Brands, error) {
	return &Brands{}, nil
}

func (r *queryResolver) Catalogues(ctx context.Context) (*Catalogues, error) {
	return &Catalogues{}, nil
}

func (r *queryResolver) Snowboards(ctx context.Context) (*Snowboards, error) {
	return &Snowboards{}, nil
}

func (r *queryResolver) Version(ctx context.Context) (string, error) {
	return "0", nil
}
