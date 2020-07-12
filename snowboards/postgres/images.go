package postgres

import (
	"context"
	"github.com/elgris/sqrl"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ztsu/snowboardsdb/snowboards"
)

const (
	imagesTableName = `"snowboards"."images"`
)

type imageStore struct {
	pg *pgxpool.Pool
}

func NewImageStore(pg *pgxpool.Pool) *imageStore {
	return &imageStore{pg: pg}
}

func (r *imageStore) Count(ctx context.Context, query snowboards.ImageQuery) (int, error) {
	var (
		total int
	)

	q := SelectFromImages(`count("id")`).Where(query)

	sql, args, err := q.ToSql()
	if err != nil {
		return 0, err
	}

	err = r.pg.QueryRow(ctx, sql, args...).Scan(&total)

	return total, err
}

func (r *imageStore) List(ctx context.Context, query snowboards.ImageQuery) ([]*snowboards.Image, error) {
	sb := SelectFromImages(
		"id",
		`"snowboardId"`,
		`"url"`,
		`"size"`,
		`"colorOfBase"`,
	).Where(query).LimitOffset(query)

	if len(query.Sort) > 0 {
		for _, s := range query.Sort {
			switch s {
			case snowboards.ImageQuerySortID:
				sb.OrderBy(`"id"`)
			case snowboards.ImageQuerySortIDDesc:
				sb.OrderBy(`"id" desc`)
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

	images := make([]*snowboards.Image, 0)
	
	for rows.Next() {
		image := new(snowboards.Image)

		err := rows.Scan(
			&image.ID,
			&image.SnowboardID,
			&image.URL,
			&image.Size,
			&image.ColorOfBase,
		)
		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}

//
type imageSelectBuilder struct {
	*sqrl.SelectBuilder
}

func SelectFromImages(columns ...string) *imageSelectBuilder {
	return &imageSelectBuilder{
		sqrl.Select(columns...).
			From(imagesTableName).
			PlaceholderFormat(sqrl.Dollar),
	}
}

func (brands *imageSelectBuilder) Where(query snowboards.ImageQuery) *imageSelectBuilder {
	if len(query.ID) > 0 {
		brands.SelectBuilder = brands.SelectBuilder.Where(sqrl.Eq{"id": query.ID})
	}

	if len(query.SnowboardID) > 0 {
		brands.SelectBuilder = brands.SelectBuilder.Where(sqrl.Eq{`"snowboardId"`: query.SnowboardID})
	}

	return brands
}

func (brands *imageSelectBuilder) LimitOffset(query snowboards.ImageQuery) *imageSelectBuilder {
	if query.Limit != nil {
		brands.SelectBuilder = brands.SelectBuilder.Limit(*query.Limit)
	}

	if query.Offset != nil {
		brands.SelectBuilder = brands.SelectBuilder.Offset(*query.Offset)
	}

	return brands
}
