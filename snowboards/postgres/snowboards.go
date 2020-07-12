package postgres

import (
	"context"
	"github.com/elgris/sqrl"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ztsu/snowboardsdb/snowboards"
)

const (
	snowboardsTableName = `"snowboards"."snowboards"`
)

type snowboardsStore struct {
	pg *pgxpool.Pool
}

func NewSnowboardsStore(pg *pgxpool.Pool) *snowboardsStore {
	return &snowboardsStore{pg: pg}
}

func (store *snowboardsStore) Count(ctx context.Context, query snowboards.SnowboardsQuery) (int, error) {
	var (
		total int
	)

	q := SelectFromSnowboards(`count("id")`).Where(query)

	sql, args, err := q.ToSql()
	if err != nil {
		return 0, err
	}

	err = store.pg.QueryRow(ctx, sql, args...).Scan(&total)

	return total, err
}

func (store *snowboardsStore) List(ctx context.Context, query snowboards.SnowboardsQuery) ([]*snowboards.Snowboard, error) {
	q := SelectFromSnowboards("id", `"brandId"`, "name", "season", "type").
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

	list := make([]*snowboards.Snowboard, 0)

	for rows.Next() {
		snowboard := new(snowboards.Snowboard)

		err := rows.Scan(
			&snowboard.ID,
			&snowboard.BrandID,
			&snowboard.Name,
			&snowboard.Season,
			&snowboard.Type,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, snowboard)
	}

	return list, nil
}

//
type snowboardsSelectBuilder struct {
	*sqrl.SelectBuilder
}

func SelectFromSnowboards(columns ...string) *snowboardsSelectBuilder {
	return &snowboardsSelectBuilder{
		sqrl.Select(columns...).
			From(snowboardsTableName).
			PlaceholderFormat(sqrl.Dollar),
	}
}

func (builder *snowboardsSelectBuilder) Where(query snowboards.SnowboardsQuery) *snowboardsSelectBuilder {
	if len(query.ID) > 0 {
		builder.SelectBuilder = builder.SelectBuilder.Where(sqrl.Eq{"id": query.ID})
	}

	if len(query.Season) > 0 {
		builder.SelectBuilder = builder.SelectBuilder.Where(sqrl.Eq{"season": query.Season})
	}

	if len(query.BrandID) > 0 {
		builder.SelectBuilder = builder.SelectBuilder.Where(sqrl.Eq{`"brandId"`: query.BrandID})
	}

	return builder
}

func (builder *snowboardsSelectBuilder) Sort(query snowboards.SnowboardsQuery) *snowboardsSelectBuilder {
	if len(query.Sort) > 0 {
		for _, s := range query.Sort {
			switch s {
			case snowboards.SnowboardsQuerySortID:
				builder.SelectBuilder = builder.SelectBuilder.OrderBy(`"id"`)
			case snowboards.SnowboardsQuerySortIDDesc:
				builder.SelectBuilder = builder.SelectBuilder.OrderBy(`"id" desc`)
			case snowboards.SnowboardsQuerySortName:
				builder.SelectBuilder = builder.SelectBuilder.OrderBy(`"name"`)
			case snowboards.SnowboardsQuerySortNameDesc:
				builder.SelectBuilder = builder.SelectBuilder.OrderBy(`"name" desc`)
			case snowboards.SnowboardsQuerySortSeason:
				builder.SelectBuilder = builder.SelectBuilder.OrderBy(`"season"`)
			case snowboards.SnowboardsQuerySortSeasonDesc:
				builder.SelectBuilder = builder.SelectBuilder.OrderBy(`"season" desc`)
			}
		}
	}

	return builder
}

func (builder *snowboardsSelectBuilder) LimitOffset(query snowboards.SnowboardsQuery) *snowboardsSelectBuilder {
	if query.Limit != nil {
		builder.SelectBuilder = builder.SelectBuilder.Limit(*query.Limit)
	}

	if query.Offset != nil {
		builder.SelectBuilder = builder.SelectBuilder.Offset(*query.Offset)
	}

	return builder
}
