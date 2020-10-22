// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/Teeth/app/ent/eatinghistory"
	"github.com/Teeth/app/ent/foodmenu"
	"github.com/Teeth/app/ent/mealplan"
	"github.com/Teeth/app/ent/predicate"
	"github.com/Teeth/app/ent/taste"
	"github.com/Teeth/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// EatinghistoryQuery is the builder for querying Eatinghistory entities.
type EatinghistoryQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Eatinghistory
	// eager-loading edges.
	withMealplan *MealplanQuery
	withFoodmenu *FoodmenuQuery
	withTaste    *TasteQuery
	withUser     *UserQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (eq *EatinghistoryQuery) Where(ps ...predicate.Eatinghistory) *EatinghistoryQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit adds a limit step to the query.
func (eq *EatinghistoryQuery) Limit(limit int) *EatinghistoryQuery {
	eq.limit = &limit
	return eq
}

// Offset adds an offset step to the query.
func (eq *EatinghistoryQuery) Offset(offset int) *EatinghistoryQuery {
	eq.offset = &offset
	return eq
}

// Order adds an order step to the query.
func (eq *EatinghistoryQuery) Order(o ...OrderFunc) *EatinghistoryQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryMealplan chains the current query on the mealplan edge.
func (eq *EatinghistoryQuery) QueryMealplan() *MealplanQuery {
	query := &MealplanQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(eatinghistory.Table, eatinghistory.FieldID, eq.sqlQuery()),
			sqlgraph.To(mealplan.Table, mealplan.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, eatinghistory.MealplanTable, eatinghistory.MealplanColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFoodmenu chains the current query on the foodmenu edge.
func (eq *EatinghistoryQuery) QueryFoodmenu() *FoodmenuQuery {
	query := &FoodmenuQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(eatinghistory.Table, eatinghistory.FieldID, eq.sqlQuery()),
			sqlgraph.To(foodmenu.Table, foodmenu.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, eatinghistory.FoodmenuTable, eatinghistory.FoodmenuColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTaste chains the current query on the taste edge.
func (eq *EatinghistoryQuery) QueryTaste() *TasteQuery {
	query := &TasteQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(eatinghistory.Table, eatinghistory.FieldID, eq.sqlQuery()),
			sqlgraph.To(taste.Table, taste.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, eatinghistory.TasteTable, eatinghistory.TasteColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the user edge.
func (eq *EatinghistoryQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(eatinghistory.Table, eatinghistory.FieldID, eq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, eatinghistory.UserTable, eatinghistory.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Eatinghistory entity in the query. Returns *NotFoundError when no eatinghistory was found.
func (eq *EatinghistoryQuery) First(ctx context.Context) (*Eatinghistory, error) {
	es, err := eq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(es) == 0 {
		return nil, &NotFoundError{eatinghistory.Label}
	}
	return es[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EatinghistoryQuery) FirstX(ctx context.Context) *Eatinghistory {
	e, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return e
}

// FirstID returns the first Eatinghistory id in the query. Returns *NotFoundError when no id was found.
func (eq *EatinghistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{eatinghistory.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (eq *EatinghistoryQuery) FirstXID(ctx context.Context) int {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Eatinghistory entity in the query, returns an error if not exactly one entity was returned.
func (eq *EatinghistoryQuery) Only(ctx context.Context) (*Eatinghistory, error) {
	es, err := eq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(es) {
	case 1:
		return es[0], nil
	case 0:
		return nil, &NotFoundError{eatinghistory.Label}
	default:
		return nil, &NotSingularError{eatinghistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EatinghistoryQuery) OnlyX(ctx context.Context) *Eatinghistory {
	e, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return e
}

// OnlyID returns the only Eatinghistory id in the query, returns an error if not exactly one id was returned.
func (eq *EatinghistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{eatinghistory.Label}
	default:
		err = &NotSingularError{eatinghistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EatinghistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Eatinghistories.
func (eq *EatinghistoryQuery) All(ctx context.Context) ([]*Eatinghistory, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return eq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eq *EatinghistoryQuery) AllX(ctx context.Context) []*Eatinghistory {
	es, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return es
}

// IDs executes the query and returns a list of Eatinghistory ids.
func (eq *EatinghistoryQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := eq.Select(eatinghistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EatinghistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EatinghistoryQuery) Count(ctx context.Context) (int, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return eq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EatinghistoryQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EatinghistoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return eq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EatinghistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EatinghistoryQuery) Clone() *EatinghistoryQuery {
	return &EatinghistoryQuery{
		config:     eq.config,
		limit:      eq.limit,
		offset:     eq.offset,
		order:      append([]OrderFunc{}, eq.order...),
		unique:     append([]string{}, eq.unique...),
		predicates: append([]predicate.Eatinghistory{}, eq.predicates...),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

//  WithMealplan tells the query-builder to eager-loads the nodes that are connected to
// the "mealplan" edge. The optional arguments used to configure the query builder of the edge.
func (eq *EatinghistoryQuery) WithMealplan(opts ...func(*MealplanQuery)) *EatinghistoryQuery {
	query := &MealplanQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withMealplan = query
	return eq
}

//  WithFoodmenu tells the query-builder to eager-loads the nodes that are connected to
// the "foodmenu" edge. The optional arguments used to configure the query builder of the edge.
func (eq *EatinghistoryQuery) WithFoodmenu(opts ...func(*FoodmenuQuery)) *EatinghistoryQuery {
	query := &FoodmenuQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withFoodmenu = query
	return eq
}

//  WithTaste tells the query-builder to eager-loads the nodes that are connected to
// the "taste" edge. The optional arguments used to configure the query builder of the edge.
func (eq *EatinghistoryQuery) WithTaste(opts ...func(*TasteQuery)) *EatinghistoryQuery {
	query := &TasteQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withTaste = query
	return eq
}

//  WithUser tells the query-builder to eager-loads the nodes that are connected to
// the "user" edge. The optional arguments used to configure the query builder of the edge.
func (eq *EatinghistoryQuery) WithUser(opts ...func(*UserQuery)) *EatinghistoryQuery {
	query := &UserQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withUser = query
	return eq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AddedTime time.Time `json:"added_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Eatinghistory.Query().
//		GroupBy(eatinghistory.FieldAddedTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (eq *EatinghistoryQuery) GroupBy(field string, fields ...string) *EatinghistoryGroupBy {
	group := &EatinghistoryGroupBy{config: eq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		AddedTime time.Time `json:"added_time,omitempty"`
//	}
//
//	client.Eatinghistory.Query().
//		Select(eatinghistory.FieldAddedTime).
//		Scan(ctx, &v)
//
func (eq *EatinghistoryQuery) Select(field string, fields ...string) *EatinghistorySelect {
	selector := &EatinghistorySelect{config: eq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eq.sqlQuery(), nil
	}
	return selector
}

func (eq *EatinghistoryQuery) prepareQuery(ctx context.Context) error {
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EatinghistoryQuery) sqlAll(ctx context.Context) ([]*Eatinghistory, error) {
	var (
		nodes       = []*Eatinghistory{}
		withFKs     = eq.withFKs
		_spec       = eq.querySpec()
		loadedTypes = [4]bool{
			eq.withMealplan != nil,
			eq.withFoodmenu != nil,
			eq.withTaste != nil,
			eq.withUser != nil,
		}
	)
	if eq.withMealplan != nil || eq.withFoodmenu != nil || eq.withTaste != nil || eq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, eatinghistory.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Eatinghistory{config: eq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := eq.withMealplan; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Eatinghistory)
		for i := range nodes {
			if fk := nodes[i].mealplan_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(mealplan.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "mealplan_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Mealplan = n
			}
		}
	}

	if query := eq.withFoodmenu; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Eatinghistory)
		for i := range nodes {
			if fk := nodes[i].foodmenu_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(foodmenu.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "foodmenu_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Foodmenu = n
			}
		}
	}

	if query := eq.withTaste; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Eatinghistory)
		for i := range nodes {
			if fk := nodes[i].taste_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(taste.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "taste_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Taste = n
			}
		}
	}

	if query := eq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Eatinghistory)
		for i := range nodes {
			if fk := nodes[i].user_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (eq *EatinghistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EatinghistoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := eq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (eq *EatinghistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eatinghistory.Table,
			Columns: eatinghistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eatinghistory.FieldID,
			},
		},
		From:   eq.sql,
		Unique: true,
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EatinghistoryQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(eatinghistory.Table)
	selector := builder.Select(t1.Columns(eatinghistory.Columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(eatinghistory.Columns...)...)
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EatinghistoryGroupBy is the builder for group-by Eatinghistory entities.
type EatinghistoryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EatinghistoryGroupBy) Aggregate(fns ...AggregateFunc) *EatinghistoryGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the group-by query and scan the result into the given value.
func (egb *EatinghistoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := egb.path(ctx)
	if err != nil {
		return err
	}
	egb.sql = query
	return egb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (egb *EatinghistoryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := egb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (egb *EatinghistoryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EatinghistoryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (egb *EatinghistoryGroupBy) StringsX(ctx context.Context) []string {
	v, err := egb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (egb *EatinghistoryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = egb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eatinghistory.Label}
	default:
		err = fmt.Errorf("ent: EatinghistoryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (egb *EatinghistoryGroupBy) StringX(ctx context.Context) string {
	v, err := egb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (egb *EatinghistoryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EatinghistoryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (egb *EatinghistoryGroupBy) IntsX(ctx context.Context) []int {
	v, err := egb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (egb *EatinghistoryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = egb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eatinghistory.Label}
	default:
		err = fmt.Errorf("ent: EatinghistoryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (egb *EatinghistoryGroupBy) IntX(ctx context.Context) int {
	v, err := egb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (egb *EatinghistoryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EatinghistoryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (egb *EatinghistoryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := egb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (egb *EatinghistoryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = egb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eatinghistory.Label}
	default:
		err = fmt.Errorf("ent: EatinghistoryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (egb *EatinghistoryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := egb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (egb *EatinghistoryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EatinghistoryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (egb *EatinghistoryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := egb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (egb *EatinghistoryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = egb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eatinghistory.Label}
	default:
		err = fmt.Errorf("ent: EatinghistoryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (egb *EatinghistoryGroupBy) BoolX(ctx context.Context) bool {
	v, err := egb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (egb *EatinghistoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := egb.sqlQuery().Query()
	if err := egb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (egb *EatinghistoryGroupBy) sqlQuery() *sql.Selector {
	selector := egb.sql
	columns := make([]string, 0, len(egb.fields)+len(egb.fns))
	columns = append(columns, egb.fields...)
	for _, fn := range egb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(egb.fields...)
}

// EatinghistorySelect is the builder for select fields of Eatinghistory entities.
type EatinghistorySelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (es *EatinghistorySelect) Scan(ctx context.Context, v interface{}) error {
	query, err := es.path(ctx)
	if err != nil {
		return err
	}
	es.sql = query
	return es.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (es *EatinghistorySelect) ScanX(ctx context.Context, v interface{}) {
	if err := es.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (es *EatinghistorySelect) Strings(ctx context.Context) ([]string, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EatinghistorySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (es *EatinghistorySelect) StringsX(ctx context.Context) []string {
	v, err := es.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (es *EatinghistorySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = es.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eatinghistory.Label}
	default:
		err = fmt.Errorf("ent: EatinghistorySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (es *EatinghistorySelect) StringX(ctx context.Context) string {
	v, err := es.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (es *EatinghistorySelect) Ints(ctx context.Context) ([]int, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EatinghistorySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (es *EatinghistorySelect) IntsX(ctx context.Context) []int {
	v, err := es.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (es *EatinghistorySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = es.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eatinghistory.Label}
	default:
		err = fmt.Errorf("ent: EatinghistorySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (es *EatinghistorySelect) IntX(ctx context.Context) int {
	v, err := es.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (es *EatinghistorySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EatinghistorySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (es *EatinghistorySelect) Float64sX(ctx context.Context) []float64 {
	v, err := es.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (es *EatinghistorySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = es.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eatinghistory.Label}
	default:
		err = fmt.Errorf("ent: EatinghistorySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (es *EatinghistorySelect) Float64X(ctx context.Context) float64 {
	v, err := es.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (es *EatinghistorySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EatinghistorySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (es *EatinghistorySelect) BoolsX(ctx context.Context) []bool {
	v, err := es.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (es *EatinghistorySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = es.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eatinghistory.Label}
	default:
		err = fmt.Errorf("ent: EatinghistorySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (es *EatinghistorySelect) BoolX(ctx context.Context) bool {
	v, err := es.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (es *EatinghistorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := es.sqlQuery().Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (es *EatinghistorySelect) sqlQuery() sql.Querier {
	selector := es.sql
	selector.Select(selector.Columns(es.fields...)...)
	return selector
}
