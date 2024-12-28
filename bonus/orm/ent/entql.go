// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/twiglab/crm/bonus/orm/ent/bonusacc"
	"github.com/twiglab/crm/bonus/orm/ent/bonusitem"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   bonusacc.Table,
			Columns: bonusacc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bonusacc.FieldID,
			},
		},
		Type: "BonusAcc",
		Fields: map[string]*sqlgraph.FieldSpec{
			bonusacc.FieldCreateTime:       {Type: field.TypeTime, Column: bonusacc.FieldCreateTime},
			bonusacc.FieldUpdateTime:       {Type: field.TypeTime, Column: bonusacc.FieldUpdateTime},
			bonusacc.FieldMemberCode:       {Type: field.TypeString, Column: bonusacc.FieldMemberCode},
			bonusacc.FieldLastCleanBalance: {Type: field.TypeInt, Column: bonusacc.FieldLastCleanBalance},
			bonusacc.FieldLastCleanTs:      {Type: field.TypeInt64, Column: bonusacc.FieldLastCleanTs},
			bonusacc.FieldLastCleanTime:    {Type: field.TypeTime, Column: bonusacc.FieldLastCleanTime},
			bonusacc.FieldStatus:           {Type: field.TypeInt, Column: bonusacc.FieldStatus},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   bonusitem.Table,
			Columns: bonusitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bonusitem.FieldID,
			},
		},
		Type: "BonusItem",
		Fields: map[string]*sqlgraph.FieldSpec{
			bonusitem.FieldCreateTime:       {Type: field.TypeTime, Column: bonusitem.FieldCreateTime},
			bonusitem.FieldUpdateTime:       {Type: field.TypeTime, Column: bonusitem.FieldUpdateTime},
			bonusitem.FieldCode:             {Type: field.TypeString, Column: bonusitem.FieldCode},
			bonusitem.FieldMchID:            {Type: field.TypeString, Column: bonusitem.FieldMchID},
			bonusitem.FieldMallCode:         {Type: field.TypeString, Column: bonusitem.FieldMallCode},
			bonusitem.FieldMallName:         {Type: field.TypeString, Column: bonusitem.FieldMallName},
			bonusitem.FieldShopCode:         {Type: field.TypeString, Column: bonusitem.FieldShopCode},
			bonusitem.FieldShopName:         {Type: field.TypeString, Column: bonusitem.FieldShopName},
			bonusitem.FieldMemberCode:       {Type: field.TypeString, Column: bonusitem.FieldMemberCode},
			bonusitem.FieldWxOpenID:         {Type: field.TypeString, Column: bonusitem.FieldWxOpenID},
			bonusitem.FieldBcmbNotifyTime:   {Type: field.TypeTime, Column: bonusitem.FieldBcmbNotifyTime},
			bonusitem.FieldBcmbNotifyID:     {Type: field.TypeString, Column: bonusitem.FieldBcmbNotifyID},
			bonusitem.FieldBcmbTransCode:    {Type: field.TypeString, Column: bonusitem.FieldBcmbTransCode},
			bonusitem.FieldAmount:           {Type: field.TypeInt, Column: bonusitem.FieldAmount},
			bonusitem.FieldBcmbTransTime:    {Type: field.TypeTime, Column: bonusitem.FieldBcmbTransTime},
			bonusitem.FieldCreateTs:         {Type: field.TypeInt64, Column: bonusitem.FieldCreateTs},
			bonusitem.FieldBcmbTransPayCode: {Type: field.TypeString, Column: bonusitem.FieldBcmbTransPayCode},
			bonusitem.FieldBcmbTransType:    {Type: field.TypeInt, Column: bonusitem.FieldBcmbTransType},
			bonusitem.FieldBonus:            {Type: field.TypeInt, Column: bonusitem.FieldBonus},
			bonusitem.FieldBonusRate:        {Type: field.TypeInt, Column: bonusitem.FieldBonusRate},
			bonusitem.FieldStatus:           {Type: field.TypeInt, Column: bonusitem.FieldStatus},
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
func (baq *BonusAccQuery) addPredicate(pred func(s *sql.Selector)) {
	baq.predicates = append(baq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the BonusAccQuery builder.
func (baq *BonusAccQuery) Filter() *BonusAccFilter {
	return &BonusAccFilter{config: baq.config, predicateAdder: baq}
}

// addPredicate implements the predicateAdder interface.
func (m *BonusAccMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the BonusAccMutation builder.
func (m *BonusAccMutation) Filter() *BonusAccFilter {
	return &BonusAccFilter{config: m.config, predicateAdder: m}
}

// BonusAccFilter provides a generic filtering capability at runtime for BonusAccQuery.
type BonusAccFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *BonusAccFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *BonusAccFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(bonusacc.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *BonusAccFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(bonusacc.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *BonusAccFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(bonusacc.FieldUpdateTime))
}

// WhereMemberCode applies the entql string predicate on the member_code field.
func (f *BonusAccFilter) WhereMemberCode(p entql.StringP) {
	f.Where(p.Field(bonusacc.FieldMemberCode))
}

// WhereLastCleanBalance applies the entql int predicate on the last_clean_balance field.
func (f *BonusAccFilter) WhereLastCleanBalance(p entql.IntP) {
	f.Where(p.Field(bonusacc.FieldLastCleanBalance))
}

// WhereLastCleanTs applies the entql int64 predicate on the last_clean_ts field.
func (f *BonusAccFilter) WhereLastCleanTs(p entql.Int64P) {
	f.Where(p.Field(bonusacc.FieldLastCleanTs))
}

// WhereLastCleanTime applies the entql time.Time predicate on the last_clean_time field.
func (f *BonusAccFilter) WhereLastCleanTime(p entql.TimeP) {
	f.Where(p.Field(bonusacc.FieldLastCleanTime))
}

// WhereStatus applies the entql int predicate on the status field.
func (f *BonusAccFilter) WhereStatus(p entql.IntP) {
	f.Where(p.Field(bonusacc.FieldStatus))
}

// addPredicate implements the predicateAdder interface.
func (biq *BonusItemQuery) addPredicate(pred func(s *sql.Selector)) {
	biq.predicates = append(biq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the BonusItemQuery builder.
func (biq *BonusItemQuery) Filter() *BonusItemFilter {
	return &BonusItemFilter{config: biq.config, predicateAdder: biq}
}

// addPredicate implements the predicateAdder interface.
func (m *BonusItemMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the BonusItemMutation builder.
func (m *BonusItemMutation) Filter() *BonusItemFilter {
	return &BonusItemFilter{config: m.config, predicateAdder: m}
}

// BonusItemFilter provides a generic filtering capability at runtime for BonusItemQuery.
type BonusItemFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *BonusItemFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *BonusItemFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(bonusitem.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *BonusItemFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(bonusitem.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *BonusItemFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(bonusitem.FieldUpdateTime))
}

// WhereCode applies the entql string predicate on the code field.
func (f *BonusItemFilter) WhereCode(p entql.StringP) {
	f.Where(p.Field(bonusitem.FieldCode))
}

// WhereMchID applies the entql string predicate on the mch_id field.
func (f *BonusItemFilter) WhereMchID(p entql.StringP) {
	f.Where(p.Field(bonusitem.FieldMchID))
}

// WhereMallCode applies the entql string predicate on the mall_code field.
func (f *BonusItemFilter) WhereMallCode(p entql.StringP) {
	f.Where(p.Field(bonusitem.FieldMallCode))
}

// WhereMallName applies the entql string predicate on the mall_name field.
func (f *BonusItemFilter) WhereMallName(p entql.StringP) {
	f.Where(p.Field(bonusitem.FieldMallName))
}

// WhereShopCode applies the entql string predicate on the shop_code field.
func (f *BonusItemFilter) WhereShopCode(p entql.StringP) {
	f.Where(p.Field(bonusitem.FieldShopCode))
}

// WhereShopName applies the entql string predicate on the shop_name field.
func (f *BonusItemFilter) WhereShopName(p entql.StringP) {
	f.Where(p.Field(bonusitem.FieldShopName))
}

// WhereMemberCode applies the entql string predicate on the member_code field.
func (f *BonusItemFilter) WhereMemberCode(p entql.StringP) {
	f.Where(p.Field(bonusitem.FieldMemberCode))
}

// WhereWxOpenID applies the entql string predicate on the wx_open_id field.
func (f *BonusItemFilter) WhereWxOpenID(p entql.StringP) {
	f.Where(p.Field(bonusitem.FieldWxOpenID))
}

// WhereBcmbNotifyTime applies the entql time.Time predicate on the bcmb_notify_time field.
func (f *BonusItemFilter) WhereBcmbNotifyTime(p entql.TimeP) {
	f.Where(p.Field(bonusitem.FieldBcmbNotifyTime))
}

// WhereBcmbNotifyID applies the entql string predicate on the bcmb_notify_id field.
func (f *BonusItemFilter) WhereBcmbNotifyID(p entql.StringP) {
	f.Where(p.Field(bonusitem.FieldBcmbNotifyID))
}

// WhereBcmbTransCode applies the entql string predicate on the bcmb_trans_code field.
func (f *BonusItemFilter) WhereBcmbTransCode(p entql.StringP) {
	f.Where(p.Field(bonusitem.FieldBcmbTransCode))
}

// WhereAmount applies the entql int predicate on the amount field.
func (f *BonusItemFilter) WhereAmount(p entql.IntP) {
	f.Where(p.Field(bonusitem.FieldAmount))
}

// WhereBcmbTransTime applies the entql time.Time predicate on the bcmb_trans_time field.
func (f *BonusItemFilter) WhereBcmbTransTime(p entql.TimeP) {
	f.Where(p.Field(bonusitem.FieldBcmbTransTime))
}

// WhereCreateTs applies the entql int64 predicate on the create_ts field.
func (f *BonusItemFilter) WhereCreateTs(p entql.Int64P) {
	f.Where(p.Field(bonusitem.FieldCreateTs))
}

// WhereBcmbTransPayCode applies the entql string predicate on the bcmb_trans_pay_code field.
func (f *BonusItemFilter) WhereBcmbTransPayCode(p entql.StringP) {
	f.Where(p.Field(bonusitem.FieldBcmbTransPayCode))
}

// WhereBcmbTransType applies the entql int predicate on the bcmb_trans_type field.
func (f *BonusItemFilter) WhereBcmbTransType(p entql.IntP) {
	f.Where(p.Field(bonusitem.FieldBcmbTransType))
}

// WhereBonus applies the entql int predicate on the bonus field.
func (f *BonusItemFilter) WhereBonus(p entql.IntP) {
	f.Where(p.Field(bonusitem.FieldBonus))
}

// WhereBonusRate applies the entql int predicate on the bonus_rate field.
func (f *BonusItemFilter) WhereBonusRate(p entql.IntP) {
	f.Where(p.Field(bonusitem.FieldBonusRate))
}

// WhereStatus applies the entql int predicate on the status field.
func (f *BonusItemFilter) WhereStatus(p entql.IntP) {
	f.Where(p.Field(bonusitem.FieldStatus))
}