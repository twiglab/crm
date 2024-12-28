package pgx

import (
	"database/sql"

	"entgo.io/ent/dialect"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/twiglab/crm/psdk/cmdb"
)

func OpenDB[C, O any](dbconf cmdb.Database, newFunc func(cmdb.Database) *sql.DB, openFunc func(*sql.DB, string, ...O) (C, error), o ...O) C {
	c := cmdb.Database{
		DSN:     dbconf.DSN,
		Driver:  "pgx",
		Dialect: dialect.Postgres,
		Configs: dbconf.Configs,
	}

	return cmdb.OpenDB(c, newFunc, openFunc, o...)
}
