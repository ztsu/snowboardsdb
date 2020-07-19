package graphql

import (
	"context"
	"log"
)

//
type catalogueOnIssuuResolver struct {
	*rootResolver
}

func (resolver *catalogueOnIssuuResolver) Brand(ctx context.Context, obj *CatalogueOnIssuu) (BrandResolveResult, error) {
	if b, ok := obj.Brand.(*Brand); ok {
		brand, err := DataLoaders(ctx).Brands.Load(b.ID)
		if err != nil {
			log.Printf("can't resolve brand: %s", err)
			return BrandResolveError{Message: "can't resolve brand"}, nil
		}

		if brand == nil {
			return BrandNotFoundError{Message: "brand not found"}, nil
		}

		return brandToGraphQL(brand), nil
	}

	log.Printf("can't resolve brand: obj is not *Brand")

	return BrandResolveError{Message: "can't resolve brand"}, nil
}
