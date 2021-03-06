package graphql

import (
	"context"
	"fmt"
	"github.com/ztsu/snowboardsdb/snowboardsdb"
)

//
type snowboardsResolver struct {
	*rootResolver
}

func (r *snowboardsResolver) List(
	ctx context.Context,
	obj *Snowboards,
	filter *SnowboardListFilter,
	sort SnowboardListSort,
	limit int,
	offset int,
) (*SnowboardList, error) {
	var (
		limitUint64  = uint64(limit)
		offsetUint64 = uint64(offset)
	)

	query := snowboardsdb.SnowboardsQuery{
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
	case SnowboardListSortNameAsc:
		query.Sort = append(query.Sort, snowboardsdb.SnowboardsQuerySortName)
	case SnowboardListSortNameDesc:
		query.Sort = append(query.Sort, snowboardsdb.SnowboardsQuerySortNameDesc)
	case SnowboardListSortSeasonAsc:
		query.Sort = append(query.Sort, snowboardsdb.SnowboardsQuerySortSeason)
	case SnowboardListSortSeasonDesc:
		query.Sort = append(query.Sort, snowboardsdb.SnowboardsQuerySortSeasonDesc)
	}

	snowboards, err := r.Stores.Snowboards.List(ctx, query)
	if err != nil {
		return nil, err
	}

	var (
		items []*Snowboard
	)

	for _, c := range snowboards {
		sb, err := snowboardToGraphQL(c)
		if err != nil {
			// @todo log or return error
			continue
		}

		items = append(items, sb)
	}

	output := &SnowboardList{
		Items: items,
	}

	return output, nil
}

func (r *snowboardsResolver) Total(ctx context.Context, obj *Snowboards, filter *SnowboardListFilter) (int, error) {
	query := snowboardsdb.SnowboardsQuery{}

	if filter != nil {
		query.ID = filter.ID

		for _, s := range filter.Season {
			query.Season = append(query.Season, string(s))
		}

		if len(filter.BrandID) > 0 {
			query.BrandID = filter.BrandID
		}
	}

	return r.Stores.Snowboards.Count(ctx, query)
}

func snowboardToGraphQL(c *snowboardsdb.Snowboard) (*Snowboard, error) {
	if c == nil {
		return nil, nil
	}

	s, err := seasonToGraphQL(c.Season)
	if err != nil {
		return nil, err
	}

	st, err := snowboardTypeToGraphQL(c.Type)
	if err != nil {
		return nil, err
	}

	return &Snowboard{
		ID:     c.ID,
		Name:   c.Name,
		Type:   *st,
		Season: *s,
		Brand:  &Brand{ID: c.BrandID},
	}, nil
}

func snowboardTypeToGraphQL(st snowboardsdb.SnowboardType) (*SnowboardType, error) {
	typeToGraphQL := map[snowboardsdb.SnowboardType]SnowboardType{
		snowboardsdb.SnowboardTypeSnowboard:   SnowboardTypeSnowboard,
		snowboardsdb.SnowboardTypeSplitboard:  SnowboardTypeSplitboard,
		snowboardsdb.SnowboardTypePowsurfer:   SnowboardTypePowsurfer,
		snowboardsdb.SnowboardTypeSplitsurfer: SnowboardTypeSplitsurfer,
		snowboardsdb.SnowboardTypeSnowskate:   SnowboardTypeSnowskate,
	}

	if t, ok := typeToGraphQL[st]; ok {
		return &t, nil
	}

	return nil, fmt.Errorf("unknown snowboard type %s", st)
}
