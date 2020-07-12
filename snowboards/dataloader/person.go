package dataloader

import (
	"context"
	"github.com/ztsu/snowboardsdb/snowboards"
)

func NewPersonLoaderConfig(ctx context.Context, store snowboards.PersonsStore) PersonLoaderConfig {
	return PersonLoaderConfig{
		MaxBatch: 100,
		Fetch: func(keys []int) ([]*snowboards.Person, []error) {
			result := make([]*snowboards.Person, len(keys))
			errors := make([]error, len(keys))

			persons, err := store.List(ctx, snowboards.PersonsQuery{ID: keys})
			if err != nil {
				return nil, []error{err}
			}

			personsByID := make(map[int]*snowboards.Person)
			for _, per := range persons {
				personsByID[per.ID] = per
			}

			for i, id := range keys {
				result[i] = personsByID[id]
				errors[i] = err
			}

			return result, errors
		},
	}
}
