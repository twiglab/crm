// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/twiglab/crm/member/orm/ent/member"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 1)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: member.FieldID,
			},
		},
		Type: "Member",
		Fields: map[string]*sqlgraph.FieldSpec{
			member.FieldCreateTime: {Type: field.TypeTime, Column: member.FieldCreateTime},
			member.FieldUpdateTime: {Type: field.TypeTime, Column: member.FieldUpdateTime},
			member.FieldCode:       {Type: field.TypeString, Column: member.FieldCode},
			member.FieldPhone:      {Type: field.TypeString, Column: member.FieldPhone},
			member.FieldNickname:   {Type: field.TypeString, Column: member.FieldNickname},
			member.FieldWxOpenID:   {Type: field.TypeString, Column: member.FieldWxOpenID},
			member.FieldStatus:     {Type: field.TypeInt, Column: member.FieldStatus},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (mq *MemberQuery) addPredicate(pred func(s *sql.Selector)) {
	mq.predicates = append(mq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the MemberQuery builder.
func (mq *MemberQuery) Filter() *MemberFilter {
	return &MemberFilter{config: mq.config, predicateAdder: mq}
}

// addPredicate implements the predicateAdder interface.
func (m *MemberMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the MemberMutation builder.
func (m *MemberMutation) Filter() *MemberFilter {
	return &MemberFilter{config: m.config, predicateAdder: m}
}

// MemberFilter provides a generic filtering capability at runtime for MemberQuery.
type MemberFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *MemberFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *MemberFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(member.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *MemberFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(member.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *MemberFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(member.FieldUpdateTime))
}

// WhereCode applies the entql string predicate on the code field.
func (f *MemberFilter) WhereCode(p entql.StringP) {
	f.Where(p.Field(member.FieldCode))
}

// WherePhone applies the entql string predicate on the phone field.
func (f *MemberFilter) WherePhone(p entql.StringP) {
	f.Where(p.Field(member.FieldPhone))
}

// WhereNickname applies the entql string predicate on the nickname field.
func (f *MemberFilter) WhereNickname(p entql.StringP) {
	f.Where(p.Field(member.FieldNickname))
}

// WhereWxOpenID applies the entql string predicate on the wx_open_id field.
func (f *MemberFilter) WhereWxOpenID(p entql.StringP) {
	f.Where(p.Field(member.FieldWxOpenID))
}

// WhereStatus applies the entql int predicate on the status field.
func (f *MemberFilter) WhereStatus(p entql.IntP) {
	f.Where(p.Field(member.FieldStatus))
}
