package graphql

import (
	"context"
	"github.com/ztsu/snowboardsdb/snowboards"
)

//
type brandResolver struct {
	*rootResolver
}

func (r *brandResolver) Founders(ctx context.Context, obj *Brand) ([]*Person, error) {
	var (
		ids []int
	)

	for _, f := range obj.Founders {
		ids = append(ids, f.ID)
	}

	if len(ids) > 0 {
		p, err := DataLoaders(ctx).Persons.LoadAll(ids)
		if len(err) > 0 && err[0] != nil {
			return nil, err[0]
		}

		founders := make([]*Person, len(obj.Founders))
		for i, _ := range ids {
			founders[i] = personToGraphQL(p[i])
		}

		return founders, nil
	}

	return obj.Founders, nil
}

func personToGraphQL(p *snowboards.Person) *Person {
	if p == nil {
		return nil
	}

	return &Person{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (r *brandResolver) Catalogues(ctx context.Context, obj *Brand) (*BrandCatalogues, error) {
	return &BrandCatalogues{Brand: obj}, nil
}
