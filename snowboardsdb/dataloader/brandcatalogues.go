package dataloader

import (
	"context"
	"github.com/ztsu/snowboardsdb/snowboardsdb"
)

func NewBrandCataloguesLoaderConfig(ctx context.Context, store snowboardsdb.CataloguesStore) BrandCataloguesLoaderConfig {
	return BrandCataloguesLoaderConfig{
		MaxBatch: 100,
		Fetch: func(keys []int) ([][]*snowboardsdb.Catalogue, []error) {
			result := make([][]*snowboardsdb.Catalogue, len(keys))
			errors := make([]error, len(keys))

			catalogues, err := store.List(ctx, snowboardsdb.CataloguesQuery{BrandID: keys})
			if err != nil {
				return nil, []error{err}
			}

			cataloguesByBrand := make(map[int][]*snowboardsdb.Catalogue)
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
