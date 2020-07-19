package dataloader

import (
	"context"
	"github.com/ztsu/snowboardsdb/snowboardsdb"
)

func NewPersonLoaderConfig(ctx context.Context, store snowboardsdb.PersonsStore) PersonLoaderConfig {
	return PersonLoaderConfig{
		MaxBatch: 100,
		Fetch: func(keys []int) ([]*snowboardsdb.Person, []error) {
			result := make([]*snowboardsdb.Person, len(keys))
			errors := make([]error, len(keys))

			persons, err := store.List(ctx, snowboardsdb.PersonsQuery{ID: keys})
			if err != nil {
				return nil, []error{err}
			}

			personsByID := make(map[int]*snowboardsdb.Person)
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
