package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"github.com/twiglab/crm/adv/orm/schema/x"
)

type Adv struct {
	ent.Schema
}

func (Adv) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.Nil).
			Default(x.ID).
			Unique().
			Immutable(),

		field.String("code").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable().
			DefaultFunc(x.Code36).
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		field.String("mall_code").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		field.String("mall_name").
			MaxLen(64).
			NotEmpty().
			Unique().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(64)", // Override MySQL.
				dialect.Postgres: "char(64)", // Override Postgres.
				dialect.SQLite:   "char(64)", // Override Postgres.
			}),

		field.String("h3_index_6").
			MaxLen(36).
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		field.String("h3_index_7").
			MaxLen(36).
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		field.String("h3_index_8").
			MaxLen(36).
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		field.String("img_path").
			MaxLen(64).
			NotEmpty().
			Unique().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(64)", // Override MySQL.
				dialect.Postgres: "char(64)", // Override Postgres.
				dialect.SQLite:   "char(64)", // Override Postgres.
			}),

		field.String("url").
			MaxLen(64).
			NotEmpty().
			Unique().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(64)", // Override MySQL.
				dialect.Postgres: "char(64)", // Override Postgres.
				dialect.SQLite:   "char(64)", // Override Postgres.
			}),

		field.String("ruler").
			MaxLen(256).
			SchemaType(map[string]string{
				dialect.MySQL:    "char(256)", // Override MySQL.
				dialect.Postgres: "char(256)", // Override Postgres.
				dialect.SQLite:   "char(256)", // Override Postgres.
			}),

		field.Int("ord").Default(0),

		field.String("memo").
			MaxLen(64).
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.Time("start_time").Immutable().Default(time.Now),
		field.Time("end_time").Immutable().Default(x.AddDay(1)),

		field.Int("status").Default(1),
	}
}

func (Adv) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Adv) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
		index.Fields("status"),
		index.Fields("ord"),
		index.Fields("start_time"),
		index.Fields("end_time"),
	}
}

func (Adv) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_adv"},
	}
}
