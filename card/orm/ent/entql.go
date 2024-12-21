// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/twiglab/crm/card/orm/ent/card"

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
			Table:   card.Table,
			Columns: card.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: card.FieldID,
			},
		},
		Type: "Card",
		Fields: map[string]*sqlgraph.FieldSpec{
			card.FieldCreateTime: {Type: field.TypeTime, Column: card.FieldCreateTime},
			card.FieldUpdateTime: {Type: field.TypeTime, Column: card.FieldUpdateTime},
			card.FieldCode:       {Type: field.TypeString, Column: card.FieldCode},
			card.FieldCardCode:   {Type: field.TypeString, Column: card.FieldCardCode},
			card.FieldMemberCode: {Type: field.TypeString, Column: card.FieldMemberCode},
			card.FieldType:       {Type: field.TypeInt, Column: card.FieldType},
			card.FieldBalance:    {Type: field.TypeInt, Column: card.FieldBalance},
			card.FieldAmount:     {Type: field.TypeInt, Column: card.FieldAmount},
			card.FieldStatus:     {Type: field.TypeInt, Column: card.FieldStatus},
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
func (cq *CardQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CardQuery builder.
func (cq *CardQuery) Filter() *CardFilter {
	return &CardFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CardMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CardMutation builder.
func (m *CardMutation) Filter() *CardFilter {
	return &CardFilter{config: m.config, predicateAdder: m}
}

// CardFilter provides a generic filtering capability at runtime for CardQuery.
type CardFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CardFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *CardFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(card.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *CardFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(card.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *CardFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(card.FieldUpdateTime))
}

// WhereCode applies the entql string predicate on the code field.
func (f *CardFilter) WhereCode(p entql.StringP) {
	f.Where(p.Field(card.FieldCode))
}

// WhereCardCode applies the entql string predicate on the card_code field.
func (f *CardFilter) WhereCardCode(p entql.StringP) {
	f.Where(p.Field(card.FieldCardCode))
}

// WhereMemberCode applies the entql string predicate on the member_code field.
func (f *CardFilter) WhereMemberCode(p entql.StringP) {
	f.Where(p.Field(card.FieldMemberCode))
}

// WhereType applies the entql int predicate on the type field.
func (f *CardFilter) WhereType(p entql.IntP) {
	f.Where(p.Field(card.FieldType))
}

// WhereBalance applies the entql int predicate on the balance field.
func (f *CardFilter) WhereBalance(p entql.IntP) {
	f.Where(p.Field(card.FieldBalance))
}

// WhereAmount applies the entql int predicate on the amount field.
func (f *CardFilter) WhereAmount(p entql.IntP) {
	f.Where(p.Field(card.FieldAmount))
}

// WhereStatus applies the entql int predicate on the status field.
func (f *CardFilter) WhereStatus(p entql.IntP) {
	f.Where(p.Field(card.FieldStatus))
}