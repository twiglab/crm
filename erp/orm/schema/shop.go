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
	"github.com/twiglab/crm/psdk/code"
)

type Shop struct {
	ent.Schema
}

func (Shop) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").
			MaxLen(36).
			NotEmpty().
			Unique().
			DefaultFunc(code.Code36).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		field.String("mall_code").
			MaxLen(36).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(36)", // Override MySQL.
				dialect.Postgres: "varchar(36)", // Override Postgres.
				dialect.SQLite:   "varchar(36)", // Override Postgres.
			}),

		field.String("mall_name").
			MaxLen(256).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(256)", // Override MySQL.
				dialect.Postgres: "varchar(256)", // Override Postgres.
				dialect.SQLite:   "varchar(256)", // Override Postgres.
			}),

		field.String("contract_code").
			MaxLen(64).
			Unique().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("pos_code").
			MaxLen(36).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(36)", // Override MySQL.
				dialect.Postgres: "varchar(36)", // Override Postgres.
				dialect.SQLite:   "varchar(36)", // Override Postgres.
			}),

		field.String("shop_code").
			MaxLen(36).
			Immutable().
			Unique().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(36)", // Override MySQL.
				dialect.Postgres: "varchar(36)", // Override Postgres.
				dialect.SQLite:   "varchar(36)", // Override Postgres.
			}),

		field.String("shop_name").
			MaxLen(256).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(256)", // Override MySQL.
				dialect.Postgres: "varchar(256)", // Override Postgres.
				dialect.SQLite:   "varchar(256)", // Override Postgres.
			}),

		field.String("biz_class_1").
			MaxLen(64).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("biz_class_name_1").
			MaxLen(64).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("biz_class_2").
			MaxLen(64).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("biz_class_name_2").
			MaxLen(64).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.Time("biz_begin_time").Immutable().Default(time.Now),
		field.Time("biz_end_time").Immutable().Default(time.Now),

		field.Int("status").Default(1),
	}
}

func (Shop) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Shop) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
	}
}

func (Shop) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_shop"},
	}
}
