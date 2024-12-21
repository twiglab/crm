package orm

import (
	"context"
	"database/sql"

	_ "github.com/twiglab/crm/card/orm/ent/runtime"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	"github.com/twiglab/crm/card/orm/ent"
)

func OpenClient(in *sql.DB) *ent.Client {
	drv := entsql.OpenDB(dialect.Postgres, in)
	return ent.NewClient(ent.Driver(drv))
}

func FromURL(ctx context.Context, url string, ops ...stdlib.OptionOpenDB) (*sql.DB, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}
	return stdlib.OpenDBFromPool(pool, ops...), nil
}
