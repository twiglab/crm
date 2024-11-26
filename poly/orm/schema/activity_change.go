package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/twiglab/crm/adv/orm/schema/x"
)

type ActivityChange struct {
	ent.Schema
}

func (ActivityChange) Fields() []ent.Field {
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
		field.String("activity_code").
			MaxLen(36).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		// 经办人
		field.String("operator").
			MaxLen(36).
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),
		// 变动提交时间
		field.Time("submit_time").Immutable().Default(time.Now), // 默认值
		// 审核人
		field.String("approver").
			MaxLen(36).
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),
		// 审批时间
		field.Time("approve_time"),
		// 审批结果 1:待审批、2:已批准、3:已拒绝
		field.Int("status").Default(1),
		// 变动简介
		field.String("change_summary").NotEmpty(),
		// 变动原因
		field.String("change_reason").NotEmpty(),

		// 变动记录
		field.Text("change_record").NotEmpty(),
	}
}
