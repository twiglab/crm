package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"

	"github.com/twiglab/crm/member/orm/schema/x"
)

type Member struct {
	ent.Schema
}

func (Member) Fields() []ent.Field {
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

		field.String("phone").
			MaxLen(64).
			Unique().
			Validate(x.CheckMobile).
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("nickname").
			MaxLen(64).
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("wx_open_id").
			MaxLen(256).
			Unique().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(256)", // Override MySQL.
				dialect.Postgres: "varchar(256)", // Override Postgres.
				dialect.SQLite:   "varchar(256)", // Override Postgres.
			}),

		field.Int("status").Default(0),
	}
}

func (Member) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Member) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
		index.Fields("wx_open_id"),
		index.Fields("phone"),
	}
}

func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_member"},
	}
}
