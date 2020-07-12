package dataloader

import (
	"context"
	"github.com/ztsu/snowboardsdb/snowboards"
)

func NewBrandLoaderConfig(ctx context.Context, store snowboards.BrandsStore) BrandLoaderConfig {
	return BrandLoaderConfig{
		MaxBatch: 100,
		Fetch: func(keys []int) ([]*snowboards.Brand, []error) {
			result := make([]*snowboards.Brand, len(keys))
			errors := make([]error, len(keys))

			brands, err := store.List(ctx, snowboards.BrandsQuery{ID: keys})
			if err != nil {
				return nil, []error{err}
			}

			brandsByID := make(map[int]*snowboards.Brand)
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
