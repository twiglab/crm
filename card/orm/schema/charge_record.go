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

type ChargeRecord struct {
	ent.Schema
}

func (ChargeRecord) Fields() []ent.Field {
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

		// pay code
		field.String("pay_code").
			MaxLen(128).
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		// pay timestamp
		field.Int64("pay_ts").Default(0),

		// 扣缴, 单位分
		field.Int64("deduct").Default(0),

		// card code
		field.String("card_code").
			MaxLen(36).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		// 状态 0:未支付 1:已支付 2:已清算
		field.Int("status").Default(0),
	}
}

func (ChargeRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (ChargeRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
	}
}

func (ChargeRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_charge_record"},
	}
}
