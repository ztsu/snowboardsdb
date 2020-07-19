package graphql

import (
	"context"
	"github.com/ztsu/snowboardsdb/snowboardsdb"
)

//
type brandCataloguesResolver struct {
	*rootResolver
}

func (r *brandCataloguesResolver) Brand(ctx context.Context, obj *BrandCatalogues) (*Brand, error) {
	return obj.Brand, nil
}

func (r *brandCataloguesResolver) List(ctx context.Context, obj *BrandCatalogues, filter *BrandCataloguesListFilter, sort CatalogueListSort, limit int, offset int) (*CatalogueList, error) {
	var (
		limitUint64  = uint64(100)
		offsetUint64 = uint64(0)
	)

	query := snowboardsdb.CataloguesQuery{
		BrandID: []int{obj.Brand.ID},
		Limit:   &limitUint64,
		Offset:  &offsetUint64,
	}

	if filter != nil {
		for _, s := range filter.Season {
			query.Season = append(query.Season, string(s))
		}
	}

	switch sort {
	case CatalogueListSortIDAsc:
		query.Sort = append(query.Sort, snowboardsdb.CataloguesQuerySortID)
	case CatalogueListSortIDDesc:
		query.Sort = append(query.Sort, snowboardsdb.CataloguesQuerySortIDDesc)
	case CatalogueListSortSeasonAsc:
		query.Sort = append(query.Sort, snowboardsdb.CataloguesQuerySortSeason)
	case CatalogueListSortSeasonDesc:
		query.Sort = append(query.Sort, snowboardsdb.CataloguesQuerySortSeasonDesc)
	}

	catalogues, err := r.Stores.Catalogues.List(ctx, query)
	if err != nil {
		return nil, err
	}

	var (
		items []Catalogue
	)

	for _, c := range catalogues {
		cg, err := catalogueToGraphQL(c)
		if err != nil {
			continue // @todo
		}

		items = append(items, cg)
	}

	output := &CatalogueList{
		Items: items,
	}

	return output, nil
}

func (r *brandCataloguesResolver) Total(ctx context.Context, obj *BrandCatalogues, filter *BrandCataloguesListFilter) (int, error) {
	query := snowboardsdb.CataloguesQuery{
		BrandID: []int{obj.Brand.ID},
	}

	if filter != nil {
		for _, s := range filter.Season {
			query.Season = append(query.Season, string(s))
		}
	}

	return r.Stores.Catalogues.Count(ctx, query)
}
