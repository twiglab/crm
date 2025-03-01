// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/twiglab/crm/poly/orm/ent/poly"

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
			Table:   poly.Table,
			Columns: poly.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: poly.FieldID,
			},
		},
		Type: "Poly",
		Fields: map[string]*sqlgraph.FieldSpec{
			poly.FieldCreateTime: {Type: field.TypeTime, Column: poly.FieldCreateTime},
			poly.FieldUpdateTime: {Type: field.TypeTime, Column: poly.FieldUpdateTime},
			poly.FieldCode:       {Type: field.TypeString, Column: poly.FieldCode},
			poly.FieldMallCode:   {Type: field.TypeString, Column: poly.FieldMallCode},
			poly.FieldName:       {Type: field.TypeString, Column: poly.FieldName},
			poly.FieldTitle:      {Type: field.TypeString, Column: poly.FieldTitle},
			poly.FieldMemo:       {Type: field.TypeString, Column: poly.FieldMemo},
			poly.FieldStartTime:  {Type: field.TypeTime, Column: poly.FieldStartTime},
			poly.FieldEndTime:    {Type: field.TypeTime, Column: poly.FieldEndTime},
			poly.FieldLogoPic:    {Type: field.TypeString, Column: poly.FieldLogoPic},
			poly.FieldPic1:       {Type: field.TypeString, Column: poly.FieldPic1},
			poly.FieldPic2:       {Type: field.TypeString, Column: poly.FieldPic2},
			poly.FieldPic3:       {Type: field.TypeString, Column: poly.FieldPic3},
			poly.FieldMenkan:     {Type: field.TypeString, Column: poly.FieldMenkan},
			poly.FieldFafang:     {Type: field.TypeString, Column: poly.FieldFafang},
			poly.FieldXiaoqi:     {Type: field.TypeString, Column: poly.FieldXiaoqi},
			poly.FieldShiyong:    {Type: field.TypeString, Column: poly.FieldShiyong},
			poly.FieldStatus:     {Type: field.TypeInt32, Column: poly.FieldStatus},
			poly.FieldType:       {Type: field.TypeInt32, Column: poly.FieldType},
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
func (pq *PolyQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PolyQuery builder.
func (pq *PolyQuery) Filter() *PolyFilter {
	return &PolyFilter{config: pq.config, predicateAdder: pq}
}

// addPredicate implements the predicateAdder interface.
func (m *PolyMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PolyMutation builder.
func (m *PolyMutation) Filter() *PolyFilter {
	return &PolyFilter{config: m.config, predicateAdder: m}
}

// PolyFilter provides a generic filtering capability at runtime for PolyQuery.
type PolyFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PolyFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *PolyFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(poly.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *PolyFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(poly.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *PolyFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(poly.FieldUpdateTime))
}

// WhereCode applies the entql string predicate on the code field.
func (f *PolyFilter) WhereCode(p entql.StringP) {
	f.Where(p.Field(poly.FieldCode))
}

// WhereMallCode applies the entql string predicate on the mall_code field.
func (f *PolyFilter) WhereMallCode(p entql.StringP) {
	f.Where(p.Field(poly.FieldMallCode))
}

// WhereName applies the entql string predicate on the name field.
func (f *PolyFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(poly.FieldName))
}

// WhereTitle applies the entql string predicate on the title field.
func (f *PolyFilter) WhereTitle(p entql.StringP) {
	f.Where(p.Field(poly.FieldTitle))
}

// WhereMemo applies the entql string predicate on the memo field.
func (f *PolyFilter) WhereMemo(p entql.StringP) {
	f.Where(p.Field(poly.FieldMemo))
}

// WhereStartTime applies the entql time.Time predicate on the start_time field.
func (f *PolyFilter) WhereStartTime(p entql.TimeP) {
	f.Where(p.Field(poly.FieldStartTime))
}

// WhereEndTime applies the entql time.Time predicate on the end_time field.
func (f *PolyFilter) WhereEndTime(p entql.TimeP) {
	f.Where(p.Field(poly.FieldEndTime))
}

// WhereLogoPic applies the entql string predicate on the logo_pic field.
func (f *PolyFilter) WhereLogoPic(p entql.StringP) {
	f.Where(p.Field(poly.FieldLogoPic))
}

// WherePic1 applies the entql string predicate on the pic1 field.
func (f *PolyFilter) WherePic1(p entql.StringP) {
	f.Where(p.Field(poly.FieldPic1))
}

// WherePic2 applies the entql string predicate on the pic2 field.
func (f *PolyFilter) WherePic2(p entql.StringP) {
	f.Where(p.Field(poly.FieldPic2))
}

// WherePic3 applies the entql string predicate on the pic3 field.
func (f *PolyFilter) WherePic3(p entql.StringP) {
	f.Where(p.Field(poly.FieldPic3))
}

// WhereMenkan applies the entql string predicate on the menkan field.
func (f *PolyFilter) WhereMenkan(p entql.StringP) {
	f.Where(p.Field(poly.FieldMenkan))
}

// WhereFafang applies the entql string predicate on the fafang field.
func (f *PolyFilter) WhereFafang(p entql.StringP) {
	f.Where(p.Field(poly.FieldFafang))
}

// WhereXiaoqi applies the entql string predicate on the xiaoqi field.
func (f *PolyFilter) WhereXiaoqi(p entql.StringP) {
	f.Where(p.Field(poly.FieldXiaoqi))
}

// WhereShiyong applies the entql string predicate on the shiyong field.
func (f *PolyFilter) WhereShiyong(p entql.StringP) {
	f.Where(p.Field(poly.FieldShiyong))
}

// WhereStatus applies the entql int32 predicate on the status field.
func (f *PolyFilter) WhereStatus(p entql.Int32P) {
	f.Where(p.Field(poly.FieldStatus))
}

// WhereType applies the entql int32 predicate on the type field.
func (f *PolyFilter) WhereType(p entql.Int32P) {
	f.Where(p.Field(poly.FieldType))
}
