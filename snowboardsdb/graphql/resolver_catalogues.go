package graphql

import (
	"context"
	"fmt"
	"github.com/ztsu/snowboardsdb/snowboardsdb"
)

//
type cataloguesResolver struct {
	*rootResolver
}

func (r *cataloguesResolver) List(
	ctx context.Context,
	obj *Catalogues,
	filter *CatalogueListFilter,
	sort CatalogueListSort,
	limit int,
	offset int,
) (*CatalogueList, error) {
	var (
		limitUint64  = uint64(limit)
		offsetUint64 = uint64(offset)
	)

	query := snowboardsdb.CataloguesQuery{
		Limit:  &limitUint64,
		Offset: &offsetUint64,
	}

	if filter != nil {
		query.ID = filter.ID

		for _, s := range filter.Season {
			query.Season = append(query.Season, string(s))
		}

		if len(filter.BrandID) > 0 {
			query.BrandID = filter.BrandID
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
			// @todo log or return error
			continue
		}

		items = append(items, cg)
	}

	output := &CatalogueList{
		Items: items,
	}

	return output, nil
}

func catalogueToGraphQL(c *snowboardsdb.Catalogue) (Catalogue, error) {
	if c == nil {
		return nil, nil
	}

	s, err := seasonToGraphQL(c.Season)
	if err != nil {
		return nil, err
	}

	if snowboardsdb.CatalogueType(c.Type) == snowboardsdb.CatalogueTypeIssuu {
		return &CatalogueOnIssuu{
			ID:     c.ID,
			Season: *s,
			Link:   c.URL,
			Brand:  &Brand{ID: c.BrandID},
		}, nil
	}

	return nil, fmt.Errorf("unknown catalogue type %s", c.Type)
}

func seasonToGraphQL(season string) (*Season, error) {
	for _, s := range AllSeason {
		if s.String() == season {
			return &s, nil
		}
	}

	return nil, fmt.Errorf("unknown seasons %s", season)
}

func (r *cataloguesResolver) Total(ctx context.Context, obj *Catalogues, filter *CatalogueListFilter) (int, error) {
	query := snowboardsdb.CataloguesQuery{}

	if filter != nil {
		query.ID = filter.ID

		for _, s := range filter.Season {
			query.Season = append(query.Season, string(s))
		}

		if len(filter.BrandID) > 0 {
			query.BrandID = filter.BrandID
		}
	}

	return r.Stores.Catalogues.Count(ctx, query)
}
