package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type SignInTask struct {
	ent.Schema
}

// Fields of the SignInTask.
func (SignInTask) Fields() []ent.Field {
	return []ent.Field{
		field.Int("task_id").Unique().Comment("任务ID"),
		field.String("description").Comment("任务描述"),
		field.Time("start_time").Comment("开始时间"),
		field.Time("end_time").Comment("结束时间"),
		field.Int("reward_info").Comment("奖励信息json"),
	}
}

// Edges of the SignInTask.
func (SignInTask) Edges() []ent.Edge {
	return nil
}

func (SignInTask) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("task_id"),
	}
}
