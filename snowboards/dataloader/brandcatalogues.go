package dataloader

import (
	"context"
	"github.com/ztsu/snowboardsdb/snowboards"
)

func NewBrandCataloguesLoaderConfig(ctx context.Context, store snowboards.CataloguesStore) BrandCataloguesLoaderConfig {
	return BrandCataloguesLoaderConfig{
		MaxBatch: 100,
		Fetch: func(keys []int) ([][]*snowboards.Catalogue, []error) {
			result := make([][]*snowboards.Catalogue, len(keys))
			errors := make([]error, len(keys))

			catalogues, err := store.List(ctx, snowboards.CataloguesQuery{BrandID: keys})
			if err != nil {
				return nil, []error{err}
			}

			cataloguesByBrand := make(map[int][]*snowboards.Catalogue)
			for _, c := range catalogues {
				cataloguesByBrand[c.BrandID] = append(cataloguesByBrand[c.BrandID], c)
			}

			for i, id := range keys {
				result[i] = cataloguesByBrand[id]
				errors[i] = err
			}

			return result, errors
		},
	}
}
