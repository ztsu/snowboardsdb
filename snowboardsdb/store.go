package snowboardsdb

import "context"

type BrandsStore interface {
	Count(ctx context.Context, query BrandsQuery) (int, error)
	List(ctx context.Context, query BrandsQuery) ([]*Brand, error)
}

type BrandsQuery struct {
	ID       []int
	NameLike *string
	Sort     []BrandsQuerySort
	Limit    *uint64
	Offset   *uint64
}

type BrandsQuerySort int

const (
	BrandsQuerySortID BrandsQuerySort = iota + 1
	BrandsQuerySortIDDesc
	BrandsQuerySortName
	BrandsQuerySortNameDesc
)

type PersonsStore interface {
	List(ctx context.Context, query PersonsQuery) ([]*Person, error)
}

type PersonsQuery struct {
	ID     []int
	Limit  *uint64
	Offset *uint64
}

type CataloguesStore interface {
	Count(ctx context.Context, query CataloguesQuery) (int, error)
	List(ctx context.Context, query CataloguesQuery) ([]*Catalogue, error)
}

type CataloguesQuery struct {
	ID      []int
	BrandID []int
	Season  []string
	Sort    []CataloguesQuerySort
	Limit   *uint64
	Offset  *uint64
}

type CataloguesQuerySort int

const (
	CataloguesQuerySortID CataloguesQuerySort = iota + 1
	CataloguesQuerySortIDDesc
	CataloguesQuerySortSeason
	CataloguesQuerySortSeasonDesc
)

type SnowboardsStore interface {
	Count(ctx context.Context, query SnowboardsQuery) (int, error)
	List(ctx context.Context, query SnowboardsQuery) ([]*Snowboard, error)
}

type SnowboardsQuery struct {
	ID      []int
	BrandID []int
	Season  []string
	Sort    []SnowboardsQuerySort
	Limit   *uint64
	Offset  *uint64
}

type SnowboardsQuerySort int

const (
	SnowboardsQuerySortID SnowboardsQuerySort = iota + 1
	SnowboardsQuerySortIDDesc
	SnowboardsQuerySortName
	SnowboardsQuerySortNameDesc
	SnowboardsQuerySortSeason
	SnowboardsQuerySortSeasonDesc
)

type ImageStore interface {
	Count(ctx context.Context, query ImageQuery) (int, error)
	List(ctx context.Context, query ImageQuery) ([]*Image, error)
}

type ImageQuery struct {
	ID          []int
	SnowboardID []int
	Sort        []ImageQuerySort
	Limit       *uint64
	Offset      *uint64
}

type ImageQuerySort int

const (
	ImageQuerySortID ImageQuerySort = iota + 1
	ImageQuerySortIDDesc
)
