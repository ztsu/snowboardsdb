package postgres

import (
	"context"
	"github.com/elgris/sqrl"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ztsu/snowboardsdb/snowboards"
)

const (
	personsTableName = `"snowboards"."persons"`
)

type personsStore struct {
	pg *pgxpool.Pool
}

func NewPersonsStore(pg *pgxpool.Pool) *personsStore {
	return &personsStore{pg: pg}
}

func (r *personsStore) List(ctx context.Context, query snowboards.PersonsQuery) ([]*snowboards.Person, error) {
	q := SelectFromPersons("id", "name").Where(query)

	sql, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pg.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	persons := make([]*snowboards.Person, 0)

	for rows.Next() {
		person := new(snowboards.Person)

		err := rows.Scan(
			&person.ID,
			&person.Name,
		)
		if err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

//
type personsSelectBuilder struct {
	*sqrl.SelectBuilder
}

func SelectFromPersons(columns ...string) *personsSelectBuilder {
	return &personsSelectBuilder{
		sqrl.Select(columns...).
			From(personsTableName).
			PlaceholderFormat(sqrl.Dollar),
	}
}

func (persons *personsSelectBuilder) Where(query snowboards.PersonsQuery) *personsSelectBuilder {
	if len(query.ID) > 0 {
		persons.SelectBuilder = persons.SelectBuilder.Where(sqrl.Eq{"id": query.ID})
	}

	return persons
}
