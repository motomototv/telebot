// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/godcong/telebot/database/ent/predicate"
	"github.com/godcong/telebot/database/ent/statistic"
)

// StatisticQuery is the builder for querying Statistic entities.
type StatisticQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Statistic
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StatisticQuery builder.
func (sq *StatisticQuery) Where(ps ...predicate.Statistic) *StatisticQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *StatisticQuery) Limit(limit int) *StatisticQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *StatisticQuery) Offset(offset int) *StatisticQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *StatisticQuery) Unique(unique bool) *StatisticQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *StatisticQuery) Order(o ...OrderFunc) *StatisticQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// First returns the first Statistic entity from the query.
// Returns a *NotFoundError when no Statistic was found.
func (sq *StatisticQuery) First(ctx context.Context) (*Statistic, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{statistic.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *StatisticQuery) FirstX(ctx context.Context) *Statistic {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Statistic ID from the query.
// Returns a *NotFoundError when no Statistic ID was found.
func (sq *StatisticQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{statistic.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *StatisticQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Statistic entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Statistic entity is not found.
// Returns a *NotFoundError when no Statistic entities are found.
func (sq *StatisticQuery) Only(ctx context.Context) (*Statistic, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{statistic.Label}
	default:
		return nil, &NotSingularError{statistic.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *StatisticQuery) OnlyX(ctx context.Context) *Statistic {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Statistic ID in the query.
// Returns a *NotSingularError when exactly one Statistic ID is not found.
// Returns a *NotFoundError when no entities are found.
func (sq *StatisticQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{statistic.Label}
	default:
		err = &NotSingularError{statistic.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *StatisticQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Statistics.
func (sq *StatisticQuery) All(ctx context.Context) ([]*Statistic, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *StatisticQuery) AllX(ctx context.Context) []*Statistic {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Statistic IDs.
func (sq *StatisticQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sq.Select(statistic.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *StatisticQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *StatisticQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *StatisticQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *StatisticQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *StatisticQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StatisticQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *StatisticQuery) Clone() *StatisticQuery {
	if sq == nil {
		return nil
	}
	return &StatisticQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]OrderFunc{}, sq.order...),
		predicates: append([]predicate.Statistic{}, sq.predicates...),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FirstName string `json:"first_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Statistic.Query().
//		GroupBy(statistic.FieldFirstName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sq *StatisticQuery) GroupBy(field string, fields ...string) *StatisticGroupBy {
	group := &StatisticGroupBy{config: sq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FirstName string `json:"first_name,omitempty"`
//	}
//
//	client.Statistic.Query().
//		Select(statistic.FieldFirstName).
//		Scan(ctx, &v)
//
func (sq *StatisticQuery) Select(fields ...string) *StatisticSelect {
	sq.fields = append(sq.fields, fields...)
	return &StatisticSelect{StatisticQuery: sq}
}

func (sq *StatisticQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !statistic.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *StatisticQuery) sqlAll(ctx context.Context) ([]*Statistic, error) {
	var (
		nodes = []*Statistic{}
		_spec = sq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Statistic{config: sq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sq *StatisticQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *StatisticQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sq *StatisticQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statistic.Table,
			Columns: statistic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statistic.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, statistic.FieldID)
		for i := range fields {
			if fields[i] != statistic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *StatisticQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(statistic.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = statistic.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StatisticGroupBy is the group-by builder for Statistic entities.
type StatisticGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *StatisticGroupBy) Aggregate(fns ...AggregateFunc) *StatisticGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *StatisticGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgb *StatisticGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgb *StatisticGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: StatisticGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgb *StatisticGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgb *StatisticGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statistic.Label}
	default:
		err = fmt.Errorf("ent: StatisticGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sgb *StatisticGroupBy) StringX(ctx context.Context) string {
	v, err := sgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgb *StatisticGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: StatisticGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgb *StatisticGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgb *StatisticGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statistic.Label}
	default:
		err = fmt.Errorf("ent: StatisticGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sgb *StatisticGroupBy) IntX(ctx context.Context) int {
	v, err := sgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgb *StatisticGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: StatisticGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgb *StatisticGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgb *StatisticGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statistic.Label}
	default:
		err = fmt.Errorf("ent: StatisticGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sgb *StatisticGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgb *StatisticGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: StatisticGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgb *StatisticGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgb *StatisticGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statistic.Label}
	default:
		err = fmt.Errorf("ent: StatisticGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sgb *StatisticGroupBy) BoolX(ctx context.Context) bool {
	v, err := sgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgb *StatisticGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sgb.fields {
		if !statistic.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *StatisticGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// StatisticSelect is the builder for selecting fields of Statistic entities.
type StatisticSelect struct {
	*StatisticQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ss *StatisticSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.StatisticQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ss *StatisticSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ss *StatisticSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: StatisticSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ss *StatisticSelect) StringsX(ctx context.Context) []string {
	v, err := ss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ss *StatisticSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statistic.Label}
	default:
		err = fmt.Errorf("ent: StatisticSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ss *StatisticSelect) StringX(ctx context.Context) string {
	v, err := ss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ss *StatisticSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: StatisticSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ss *StatisticSelect) IntsX(ctx context.Context) []int {
	v, err := ss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ss *StatisticSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statistic.Label}
	default:
		err = fmt.Errorf("ent: StatisticSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ss *StatisticSelect) IntX(ctx context.Context) int {
	v, err := ss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ss *StatisticSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: StatisticSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ss *StatisticSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ss *StatisticSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statistic.Label}
	default:
		err = fmt.Errorf("ent: StatisticSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ss *StatisticSelect) Float64X(ctx context.Context) float64 {
	v, err := ss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ss *StatisticSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: StatisticSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ss *StatisticSelect) BoolsX(ctx context.Context) []bool {
	v, err := ss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ss *StatisticSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statistic.Label}
	default:
		err = fmt.Errorf("ent: StatisticSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ss *StatisticSelect) BoolX(ctx context.Context) bool {
	v, err := ss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ss *StatisticSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
