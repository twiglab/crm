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
)

// 大数据脚本处理，应用只有查询
// 主要用于清算
type BonusAcc struct {
	ent.Schema
}

func (BonusAcc) Fields() []ent.Field {
	return []ent.Field{
		field.String("member_code").
			MaxLen(36).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		/*
			field.String("wx_open_id").
				MaxLen(256).
				Immutable().
				SchemaType(map[string]string{
					dialect.MySQL:    "varchar(256)", // Override MySQL.
					dialect.Postgres: "varchar(256)", // Override Postgres.
					dialect.SQLite:   "varchar(256)", // Override Postgres.
				}),
		*/

		// 清算后的余额，单位分，清算脚本设置
		field.Int("last_clean_balance").Immutable().Default(0),
		// 最后一次清算时间戳,到1/1000秒
		field.Int64("last_clean_ts").Immutable().Default(0),

		field.Time("last_clean_time").Immutable().Default(time.Now),

		field.Int("status").Default(1),
	}
}

func (BonusAcc) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (BonusAcc) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("member_code"),
	}
}

func (BonusAcc) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_bonus_acc"},
	}
}
