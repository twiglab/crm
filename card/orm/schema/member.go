package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/twiglab/crm/card/orm/schema/internal/x"
)

type Card struct {
	ent.Schema
}

func (Card) Fields() []ent.Field {
	return []ent.Field{

		// Code
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

		// card code
		field.String("card_code").
			MaxLen(16).
			NotEmpty().
			Unique().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(16)", // Override MySQL.
				dialect.Postgres: "char(16)", // Override Postgres.
				dialect.SQLite:   "char(16)", // Override Postgres.
			}),

		// Member code
		field.String("member_code").
			MaxLen(36).
			DefaultFunc(x.Code36).
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		// 种类
		field.Int("type").Default(0),

		// 额度
		field.Int("balance").Default(0),

		// 余额
		field.Int("amount").Default(0),

		field.Int("status").Default(1),
	}
}

func (Card) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Card) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
		index.Fields("member_code"),
	}
}

func (Card) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_card"},
	}
}
