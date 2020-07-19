package dataloader

import (
	"context"
	"github.com/ztsu/snowboardsdb/snowboardsdb"
)

func NewBrandLoaderConfig(ctx context.Context, store snowboardsdb.BrandsStore) BrandLoaderConfig {
	return BrandLoaderConfig{
		MaxBatch: 100,
		Fetch: func(keys []int) ([]*snowboardsdb.Brand, []error) {
			result := make([]*snowboardsdb.Brand, len(keys))
			errors := make([]error, len(keys))

			brands, err := store.List(ctx, snowboardsdb.BrandsQuery{ID: keys})
			if err != nil {
				return nil, []error{err}
			}

			brandsByID := make(map[int]*snowboardsdb.Brand)
			for _, per := range brands {
				brandsByID[per.ID] = per
			}

			for i, id := range keys {
				result[i] = brandsByID[id]
				errors[i] = err
			}

			return result, errors
		},
	}
}
