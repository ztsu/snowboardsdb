package graphql

import (
	"github.com/ztsu/snowboardsdb/snowboards"
)

type Stores struct {
	Brands     snowboards.BrandsStore
	Persons    snowboards.PersonsStore
	Catalogues snowboards.CataloguesStore
	Snowboards snowboards.SnowboardsStore
	Images     snowboards.ImageStore
}

type rootResolver struct {
	*Stores
}

func NewRootResolver(services *Stores) ResolverRoot {
	return &rootResolver{services}
}

func (r *rootResolver) Brand() BrandResolver { return &brandResolver{r} }

func (r *rootResolver) BrandCatalogues() BrandCataloguesResolver { return &brandCataloguesResolver{r} }

func (r *rootResolver) Brands() BrandsResolver { return &brandsResolver{r} }

func (r *rootResolver) CatalogueOnIssuu() CatalogueOnIssuuResolver { return &catalogueOnIssuuResolver{r} }

func (r *rootResolver) Catalogues() CataloguesResolver { return &cataloguesResolver{r} }

func (r *rootResolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *rootResolver) Query() QueryResolver { return &queryResolver{r} }

func (r *rootResolver) Snowboard() SnowboardResolver { return &snowboardResolver{r} }

func (r *rootResolver) Snowboards() SnowboardsResolver { return &snowboardsResolver{r} }
