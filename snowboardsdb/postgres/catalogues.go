package postgres

import (
	"context"
	"github.com/elgris/sqrl"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ztsu/snowboardsdb/snowboardsdb"
)

const (
	cataloguesTableName = `"snowboards"."catalogues"`
)

type cataloguesStore struct {
	pg *pgxpool.Pool
}

func NewCataloguesStore(pg *pgxpool.Pool) *cataloguesStore {
	return &cataloguesStore{pg: pg}
}

func (store *cataloguesStore) Count(ctx context.Context, query snowboardsdb.CataloguesQuery) (int, error) {
	var (
		total int
	)

	q := SelectFromCatalogues(`count("id")`).Where(query)

	sql, args, err := q.ToSql()
	if err != nil {
		return 0, err
	}

	err = store.pg.QueryRow(ctx, sql, args...).Scan(&total)

	return total, err
}

func (store *cataloguesStore) List(ctx context.Context, query snowboardsdb.CataloguesQuery) ([]*snowboardsdb.Catalogue, error) {
	q := SelectFromCatalogues("id", `"brandId"`, "season", "type", "url", "size").
		Where(query).
		Sort(query).
		LimitOffset(query)

	sql, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := store.pg.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	catalogues := make([]*snowboardsdb.Catalogue, 0)

	for rows.Next() {
		catalogue := new(snowboardsdb.Catalogue)

		err := rows.Scan(
			&catalogue.ID,
			&catalogue.BrandID,
			&catalogue.Season,
			&catalogue.Type,
			&catalogue.URL,
			&catalogue.Size,
		)
		if err != nil {
			return nil, err
		}

		catalogues = append(catalogues, catalogue)
	}

	return catalogues, nil
}

//
type cataloguesSelectBuilder struct {
	*sqrl.SelectBuilder
}

func SelectFromCatalogues(columns ...string) *cataloguesSelectBuilder {
	return &cataloguesSelectBuilder{
		sqrl.Select(columns...).
			From(cataloguesTableName).
			PlaceholderFormat(sqrl.Dollar),
	}
}

func (catalogues *cataloguesSelectBuilder) Where(query snowboardsdb.CataloguesQuery) *cataloguesSelectBuilder {
	if len(query.ID) > 0 {
		catalogues.SelectBuilder = catalogues.SelectBuilder.Where(sqrl.Eq{"id": query.ID})
	}

	if len(query.Season) > 0 {
		catalogues.SelectBuilder = catalogues.SelectBuilder.Where(sqrl.Eq{"season": query.Season})
	}

	if len(query.BrandID) > 0 {
		catalogues.SelectBuilder = catalogues.SelectBuilder.Where(sqrl.Eq{`"brandId"`: query.BrandID})
	}

	return catalogues
}

func (catalogues *cataloguesSelectBuilder) Sort(query snowboardsdb.CataloguesQuery) *cataloguesSelectBuilder {
	if len(query.Sort) > 0 {
		for _, s := range query.Sort {
			switch s {
			case snowboardsdb.CataloguesQuerySortSeason:
				catalogues.SelectBuilder = catalogues.SelectBuilder.OrderBy(`"season"`)
			case snowboardsdb.CataloguesQuerySortSeasonDesc:
				catalogues.SelectBuilder = catalogues.SelectBuilder.OrderBy(`"season" desc`)
			}
		}
	}

	return catalogues
}

func (catalogues *cataloguesSelectBuilder) LimitOffset(query snowboardsdb.CataloguesQuery) *cataloguesSelectBuilder {
	if query.Limit != nil {
		catalogues.SelectBuilder = catalogues.SelectBuilder.Limit(*query.Limit)
	}

	if query.Offset != nil {
		catalogues.SelectBuilder = catalogues.SelectBuilder.Offset(*query.Offset)
	}

	return catalogues
}
