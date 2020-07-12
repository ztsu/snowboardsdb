package graphql

import (
	"context"
	"fmt"
	"github.com/ztsu/snowboardsdb/snowboards"
	"strings"
)

// resolves query.snowboards.brands (list, total)
type brandsResolver struct {
	*rootResolver
}

func (r *brandsResolver) List(
	ctx context.Context,
	obj *Brands,
	filter *BrandListFilter,
	sort BrandListSort,
	limit int,
	offset int,
) (*BrandList, error) {
	query := sortBrandsQuery(limitBrandsQuery(brandListFilterToQuery(filter), limit, offset), sort)

	brands, err := r.Stores.Brands.List(ctx, query)
	if err != nil {
		return nil, err
	}

	var (
		items []*Brand
	)

	for _, b := range brands {
		items = append(items, brandToGraphQL(b))
	}

	output := &BrandList{
		Items: items,
	}

	return output, nil
}

func brandListFilterToQuery(filter *BrandListFilter) snowboards.BrandsQuery {
	query := snowboards.BrandsQuery{}

	if filter != nil {
		query.ID = filter.ID

		if filter.NameStartsWith != nil {
			query.NameLike = str(fmt.Sprintf("%s%%", strings.ToLower(*filter.NameStartsWith)))
		}
	}

	return query
}

func limitBrandsQuery(query snowboards.BrandsQuery, limit, offset int) snowboards.BrandsQuery {
	var (
		limitUint64  = uint64(limit)
		offsetUint64 = uint64(offset)
	)

	query.Limit = &limitUint64
	query.Offset = &offsetUint64

	return query
}

func sortBrandsQuery(query snowboards.BrandsQuery, sort BrandListSort) snowboards.BrandsQuery {
	switch sort {
	case BrandListSortIDAsc:
		query.Sort = append(query.Sort, snowboards.BrandsQuerySortID)
	case BrandListSortIDDesc:
		query.Sort = append(query.Sort, snowboards.BrandsQuerySortIDDesc)
	case BrandListSortNameAsc:
		query.Sort = append(query.Sort, snowboards.BrandsQuerySortName)
	case BrandListSortNameDesc:
		query.Sort = append(query.Sort, snowboards.BrandsQuerySortNameDesc)
	}

	return query
}

func brandToGraphQL(b *snowboards.Brand) *Brand {
	if b == nil {
		return nil
	}

	var (
		founders []*Person
	)

	for _, p := range b.Founders {
		founders = append(founders, &Person{ID: p})
	}

	return &Brand{
		ID:          b.ID,
		Name:        b.Name,
		WebsiteURL:  b.WebsiteURL,
		Founders:    founders,
		FoundedIn:   b.FoundedIn,
		OriginsFrom: b.OriginsFrom,
	}
}

func (r *brandsResolver) Total(ctx context.Context, obj *Brands, filter *BrandListFilter) (int, error) {
	return r.Stores.Brands.Count(ctx, brandListFilterToQuery(filter))
}
