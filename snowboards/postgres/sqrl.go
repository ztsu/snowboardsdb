package postgres

import (
	"github.com/elgris/sqrl"
)

func Select(columns ...string) *sqrl.SelectBuilder {
	return sqrl.Select(columns...).PlaceholderFormat(sqrl.Dollar)
}
