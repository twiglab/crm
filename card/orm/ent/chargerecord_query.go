// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/twiglab/crm/card/orm/ent/chargerecord"
	"github.com/twiglab/crm/card/orm/ent/predicate"
)

// ChargeRecordQuery is the builder for querying ChargeRecord entities.
type ChargeRecordQuery struct {
	config
	ctx        *QueryContext
	order      []chargerecord.OrderOption
	inters     []Interceptor
	predicates []predicate.ChargeRecord
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChargeRecordQuery builder.
func (crq *ChargeRecordQuery) Where(ps ...predicate.ChargeRecord) *ChargeRecordQuery {
	crq.predicates = append(crq.predicates, ps...)
	return crq
}

// Limit the number of records to be returned by this query.
func (crq *ChargeRecordQuery) Limit(limit int) *ChargeRecordQuery {
	crq.ctx.Limit = &limit
	return crq
}

// Offset to start from.
func (crq *ChargeRecordQuery) Offset(offset int) *ChargeRecordQuery {
	crq.ctx.Offset = &offset
	return crq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (crq *ChargeRecordQuery) Unique(unique bool) *ChargeRecordQuery {
	crq.ctx.Unique = &unique
	return crq
}

// Order specifies how the records should be ordered.
func (crq *ChargeRecordQuery) Order(o ...chargerecord.OrderOption) *ChargeRecordQuery {
	crq.order = append(crq.order, o...)
	return crq
}

// First returns the first ChargeRecord entity from the query.
// Returns a *NotFoundError when no ChargeRecord was found.
func (crq *ChargeRecordQuery) First(ctx context.Context) (*ChargeRecord, error) {
	nodes, err := crq.Limit(1).All(setContextOp(ctx, crq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{chargerecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (crq *ChargeRecordQuery) FirstX(ctx context.Context) *ChargeRecord {
	node, err := crq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ChargeRecord ID from the query.
// Returns a *NotFoundError when no ChargeRecord ID was found.
func (crq *ChargeRecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(1).IDs(setContextOp(ctx, crq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{chargerecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (crq *ChargeRecordQuery) FirstIDX(ctx context.Context) int {
	id, err := crq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ChargeRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ChargeRecord entity is found.
// Returns a *NotFoundError when no ChargeRecord entities are found.
func (crq *ChargeRecordQuery) Only(ctx context.Context) (*ChargeRecord, error) {
	nodes, err := crq.Limit(2).All(setContextOp(ctx, crq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{chargerecord.Label}
	default:
		return nil, &NotSingularError{chargerecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (crq *ChargeRecordQuery) OnlyX(ctx context.Context) *ChargeRecord {
	node, err := crq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ChargeRecord ID in the query.
// Returns a *NotSingularError when more than one ChargeRecord ID is found.
// Returns a *NotFoundError when no entities are found.
func (crq *ChargeRecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(2).IDs(setContextOp(ctx, crq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{chargerecord.Label}
	default:
		err = &NotSingularError{chargerecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (crq *ChargeRecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := crq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ChargeRecords.
func (crq *ChargeRecordQuery) All(ctx context.Context) ([]*ChargeRecord, error) {
	ctx = setContextOp(ctx, crq.ctx, ent.OpQueryAll)
	if err := crq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ChargeRecord, *ChargeRecordQuery]()
	return withInterceptors[[]*ChargeRecord](ctx, crq, qr, crq.inters)
}

// AllX is like All, but panics if an error occurs.
func (crq *ChargeRecordQuery) AllX(ctx context.Context) []*ChargeRecord {
	nodes, err := crq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ChargeRecord IDs.
func (crq *ChargeRecordQuery) IDs(ctx context.Context) (ids []int, err error) {
	if crq.ctx.Unique == nil && crq.path != nil {
		crq.Unique(true)
	}
	ctx = setContextOp(ctx, crq.ctx, ent.OpQueryIDs)
	if err = crq.Select(chargerecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (crq *ChargeRecordQuery) IDsX(ctx context.Context) []int {
	ids, err := crq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (crq *ChargeRecordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, crq.ctx, ent.OpQueryCount)
	if err := crq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, crq, querierCount[*ChargeRecordQuery](), crq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (crq *ChargeRecordQuery) CountX(ctx context.Context) int {
	count, err := crq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (crq *ChargeRecordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, crq.ctx, ent.OpQueryExist)
	switch _, err := crq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (crq *ChargeRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := crq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChargeRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (crq *ChargeRecordQuery) Clone() *ChargeRecordQuery {
	if crq == nil {
		return nil
	}
	return &ChargeRecordQuery{
		config:     crq.config,
		ctx:        crq.ctx.Clone(),
		order:      append([]chargerecord.OrderOption{}, crq.order...),
		inters:     append([]Interceptor{}, crq.inters...),
		predicates: append([]predicate.ChargeRecord{}, crq.predicates...),
		// clone intermediate query.
		sql:  crq.sql.Clone(),
		path: crq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ChargeRecord.Query().
//		GroupBy(chargerecord.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (crq *ChargeRecordQuery) GroupBy(field string, fields ...string) *ChargeRecordGroupBy {
	crq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChargeRecordGroupBy{build: crq}
	grbuild.flds = &crq.ctx.Fields
	grbuild.label = chargerecord.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.ChargeRecord.Query().
//		Select(chargerecord.FieldCreateTime).
//		Scan(ctx, &v)
func (crq *ChargeRecordQuery) Select(fields ...string) *ChargeRecordSelect {
	crq.ctx.Fields = append(crq.ctx.Fields, fields...)
	sbuild := &ChargeRecordSelect{ChargeRecordQuery: crq}
	sbuild.label = chargerecord.Label
	sbuild.flds, sbuild.scan = &crq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChargeRecordSelect configured with the given aggregations.
func (crq *ChargeRecordQuery) Aggregate(fns ...AggregateFunc) *ChargeRecordSelect {
	return crq.Select().Aggregate(fns...)
}

func (crq *ChargeRecordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range crq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, crq); err != nil {
				return err
			}
		}
	}
	for _, f := range crq.ctx.Fields {
		if !chargerecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if crq.path != nil {
		prev, err := crq.path(ctx)
		if err != nil {
			return err
		}
		crq.sql = prev
	}
	return nil
}

func (crq *ChargeRecordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ChargeRecord, error) {
	var (
		nodes = []*ChargeRecord{}
		_spec = crq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ChargeRecord).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ChargeRecord{config: crq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, crq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (crq *ChargeRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := crq.querySpec()
	_spec.Node.Columns = crq.ctx.Fields
	if len(crq.ctx.Fields) > 0 {
		_spec.Unique = crq.ctx.Unique != nil && *crq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, crq.driver, _spec)
}

func (crq *ChargeRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(chargerecord.Table, chargerecord.Columns, sqlgraph.NewFieldSpec(chargerecord.FieldID, field.TypeInt))
	_spec.From = crq.sql
	if unique := crq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if crq.path != nil {
		_spec.Unique = true
	}
	if fields := crq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chargerecord.FieldID)
		for i := range fields {
			if fields[i] != chargerecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := crq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := crq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := crq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := crq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (crq *ChargeRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(crq.driver.Dialect())
	t1 := builder.Table(chargerecord.Table)
	columns := crq.ctx.Fields
	if len(columns) == 0 {
		columns = chargerecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if crq.sql != nil {
		selector = crq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if crq.ctx.Unique != nil && *crq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range crq.predicates {
		p(selector)
	}
	for _, p := range crq.order {
		p(selector)
	}
	if offset := crq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := crq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChargeRecordGroupBy is the group-by builder for ChargeRecord entities.
type ChargeRecordGroupBy struct {
	selector
	build *ChargeRecordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (crgb *ChargeRecordGroupBy) Aggregate(fns ...AggregateFunc) *ChargeRecordGroupBy {
	crgb.fns = append(crgb.fns, fns...)
	return crgb
}

// Scan applies the selector query and scans the result into the given value.
func (crgb *ChargeRecordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crgb.build.ctx, ent.OpQueryGroupBy)
	if err := crgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChargeRecordQuery, *ChargeRecordGroupBy](ctx, crgb.build, crgb, crgb.build.inters, v)
}

func (crgb *ChargeRecordGroupBy) sqlScan(ctx context.Context, root *ChargeRecordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(crgb.fns))
	for _, fn := range crgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*crgb.flds)+len(crgb.fns))
		for _, f := range *crgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*crgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChargeRecordSelect is the builder for selecting fields of ChargeRecord entities.
type ChargeRecordSelect struct {
	*ChargeRecordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (crs *ChargeRecordSelect) Aggregate(fns ...AggregateFunc) *ChargeRecordSelect {
	crs.fns = append(crs.fns, fns...)
	return crs
}

// Scan applies the selector query and scans the result into the given value.
func (crs *ChargeRecordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crs.ctx, ent.OpQuerySelect)
	if err := crs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChargeRecordQuery, *ChargeRecordSelect](ctx, crs.ChargeRecordQuery, crs, crs.inters, v)
}

func (crs *ChargeRecordSelect) sqlScan(ctx context.Context, root *ChargeRecordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(crs.fns))
	for _, fn := range crs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*crs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
