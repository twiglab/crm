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
	"github.com/twiglab/crm/bonus/orm/schema/internal/x"
)

type BonusItem struct {
	ent.Schema
}

func (BonusItem) Fields() []ent.Field {
	return []ent.Field{

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

		field.String("mch_id").
			MaxLen(36).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("mall_code").
			MaxLen(36).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("mall_name").
			MaxLen(256).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(256)", // Override MySQL.
				dialect.Postgres: "varchar(256)", // Override Postgres.
				dialect.SQLite:   "varchar(256)", // Override Postgres.
			}),

		field.String("shop_code").
			MaxLen(36).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("shop_name").
			MaxLen(256).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(256)", // Override MySQL.
				dialect.Postgres: "varchar(256)", // Override Postgres.
				dialect.SQLite:   "varchar(256)", // Override Postgres.
			}),

		field.String("member_code").
			MaxLen(36).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("wx_open_id").
			MaxLen(256).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(256)", // Override MySQL.
				dialect.Postgres: "varchar(256)", // Override Postgres.
				dialect.SQLite:   "varchar(256)", // Override Postgres.
			}),

		field.Time("bcmb_notify_time").Immutable().Default(time.Now),

		// 商圈会员消息id，用于幂等
		field.String("bcmb_notify_id").
			MaxLen(64).
			Unique().
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("bcmb_trans_code").
			MaxLen(64).
			Unique().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(32)", // Override MySQL.
				dialect.Postgres: "varchar(32)", // Override Postgres.
				dialect.SQLite:   "varchar(32)", // Override Postgres.
			}),

		field.Int("amount").Default(0).Immutable(),

		field.Time("bcmb_trans_time").Immutable().Default(time.Now),

		field.Int64("create_ts").Immutable().Default(0),

		// 如果是退款，记录原始交易单号
		field.String("bcmb_trans_pay_code").
			MaxLen(64).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(32)", // Override MySQL.
				dialect.Postgres: "varchar(32)", // Override Postgres.
				dialect.SQLite:   "varchar(32)", // Override Postgres.
			}),
		// 1 = 退款 ， 0 = 支付
		field.Int("bcmb_trans_type").Immutable().Default(0),

		field.Int("bonus").Default(0).Immutable(),

		field.Int("bonus_rate").Default(100).Comment("积分比例"),

		field.Int("status").Default(1),
	}
}

func (BonusItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (BonusItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
		index.Fields("wx_open_id"),
	}
}

func (BonusItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_bonus_list"},
	}
}
