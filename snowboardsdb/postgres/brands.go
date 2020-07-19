package postgres

import (
	"context"
	"github.com/elgris/sqrl"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ztsu/snowboardsdb/snowboardsdb"
)

const (
	brandsTableName = `"snowboards"."brands"`
)

type brandsStore struct {
	pg *pgxpool.Pool
}

func NewBrandsStore(pg *pgxpool.Pool) *brandsStore {
	return &brandsStore{pg: pg}
}

func (r *brandsStore) Count(ctx context.Context, query snowboardsdb.BrandsQuery) (int, error) {
	var (
		total int
	)

	q := SelectFromBrands(`count("id")`).Where(query)

	sql, args, err := q.ToSql()
	if err != nil {
		return 0, err
	}

	err = r.pg.QueryRow(ctx, sql, args...).Scan(&total)

	return total, err
}

func (r *brandsStore) List(ctx context.Context, query snowboardsdb.BrandsQuery) ([]*snowboardsdb.Brand, error) {
	sb := SelectFromBrands(
		"id",
		`"name"`,
		`"websiteUrl"`,
		`"founders"`,
		`"foundedIn"`,
		`"originsFrom"`,
	).Where(query).LimitOffset(query)

	if len(query.Sort) > 0 {
		for _, s := range query.Sort {
			switch s {
			case snowboardsdb.BrandsQuerySortName:
				sb.OrderBy(`"name"`)
			case snowboardsdb.BrandsQuerySortNameDesc:
				sb.OrderBy(`"name" desc`)
			}
		}
	}

	sql, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pg.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	brands := make([]*snowboardsdb.Brand, 0)

	for rows.Next() {
		brand := new(snowboardsdb.Brand)

		err := rows.Scan(
			&brand.ID,
			&brand.Name,
			&brand.WebsiteURL,
			&brand.Founders,
			&brand.FoundedIn,
			&brand.OriginsFrom,
		)
		if err != nil {
			return nil, err
		}

		brands = append(brands, brand)
	}

	return brands, nil
}

//
type brandsSelectBuilder struct {
	*sqrl.SelectBuilder
}

func SelectFromBrands(columns ...string) *brandsSelectBuilder {
	return &brandsSelectBuilder{
		sqrl.Select(columns...).
			From(brandsTableName).
			PlaceholderFormat(sqrl.Dollar),
	}
}

func (brands *brandsSelectBuilder) Where(query snowboardsdb.BrandsQuery) *brandsSelectBuilder {
	if query.NameLike != nil {
		brands.SelectBuilder = brands.SelectBuilder.Where(
			sqrl.Expr("LOWER(name) LIKE ?", *query.NameLike), // @todo replace with regexp
		)
	}

	if len(query.ID) > 0 {
		brands.SelectBuilder = brands.SelectBuilder.Where(sqrl.Eq{"id": query.ID})
	}

	return brands
}

func (brands *brandsSelectBuilder) LimitOffset(query snowboardsdb.BrandsQuery) *brandsSelectBuilder {
	if query.Limit != nil {
		brands.SelectBuilder = brands.SelectBuilder.Limit(*query.Limit)
	}

	if query.Offset != nil {
		brands.SelectBuilder = brands.SelectBuilder.Offset(*query.Offset)
	}

	return brands
}
