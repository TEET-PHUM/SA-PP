// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Teeth/app/ent/eatinghistory"
	"github.com/Teeth/app/ent/predicate"
	"github.com/Teeth/app/ent/taste"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// TasteUpdate is the builder for updating Taste entities.
type TasteUpdate struct {
	config
	hooks      []Hook
	mutation   *TasteMutation
	predicates []predicate.Taste
}

// Where adds a new predicate for the builder.
func (tu *TasteUpdate) Where(ps ...predicate.Taste) *TasteUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetTaste sets the taste field.
func (tu *TasteUpdate) SetTaste(s string) *TasteUpdate {
	tu.mutation.SetTaste(s)
	return tu
}

// AddEatinghistoryIDs adds the eatinghistorys edge to Eatinghistory by ids.
func (tu *TasteUpdate) AddEatinghistoryIDs(ids ...int) *TasteUpdate {
	tu.mutation.AddEatinghistoryIDs(ids...)
	return tu
}

// AddEatinghistorys adds the eatinghistorys edges to Eatinghistory.
func (tu *TasteUpdate) AddEatinghistorys(e ...*Eatinghistory) *TasteUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tu.AddEatinghistoryIDs(ids...)
}

// Mutation returns the TasteMutation object of the builder.
func (tu *TasteUpdate) Mutation() *TasteMutation {
	return tu.mutation
}

// RemoveEatinghistoryIDs removes the eatinghistorys edge to Eatinghistory by ids.
func (tu *TasteUpdate) RemoveEatinghistoryIDs(ids ...int) *TasteUpdate {
	tu.mutation.RemoveEatinghistoryIDs(ids...)
	return tu
}

// RemoveEatinghistorys removes eatinghistorys edges to Eatinghistory.
func (tu *TasteUpdate) RemoveEatinghistorys(e ...*Eatinghistory) *TasteUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tu.RemoveEatinghistoryIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TasteUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := tu.mutation.Taste(); ok {
		if err := taste.TasteValidator(v); err != nil {
			return 0, &ValidationError{Name: "taste", err: fmt.Errorf("ent: validator failed for field \"taste\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TasteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TasteUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TasteUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TasteUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TasteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taste.Table,
			Columns: taste.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: taste.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Taste(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taste.FieldTaste,
		})
	}
	if nodes := tu.mutation.RemovedEatinghistorysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taste.EatinghistorysTable,
			Columns: []string{taste.EatinghistorysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: eatinghistory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.EatinghistorysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taste.EatinghistorysTable,
			Columns: []string{taste.EatinghistorysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: eatinghistory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taste.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TasteUpdateOne is the builder for updating a single Taste entity.
type TasteUpdateOne struct {
	config
	hooks    []Hook
	mutation *TasteMutation
}

// SetTaste sets the taste field.
func (tuo *TasteUpdateOne) SetTaste(s string) *TasteUpdateOne {
	tuo.mutation.SetTaste(s)
	return tuo
}

// AddEatinghistoryIDs adds the eatinghistorys edge to Eatinghistory by ids.
func (tuo *TasteUpdateOne) AddEatinghistoryIDs(ids ...int) *TasteUpdateOne {
	tuo.mutation.AddEatinghistoryIDs(ids...)
	return tuo
}

// AddEatinghistorys adds the eatinghistorys edges to Eatinghistory.
func (tuo *TasteUpdateOne) AddEatinghistorys(e ...*Eatinghistory) *TasteUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tuo.AddEatinghistoryIDs(ids...)
}

// Mutation returns the TasteMutation object of the builder.
func (tuo *TasteUpdateOne) Mutation() *TasteMutation {
	return tuo.mutation
}

// RemoveEatinghistoryIDs removes the eatinghistorys edge to Eatinghistory by ids.
func (tuo *TasteUpdateOne) RemoveEatinghistoryIDs(ids ...int) *TasteUpdateOne {
	tuo.mutation.RemoveEatinghistoryIDs(ids...)
	return tuo
}

// RemoveEatinghistorys removes eatinghistorys edges to Eatinghistory.
func (tuo *TasteUpdateOne) RemoveEatinghistorys(e ...*Eatinghistory) *TasteUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tuo.RemoveEatinghistoryIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (tuo *TasteUpdateOne) Save(ctx context.Context) (*Taste, error) {
	if v, ok := tuo.mutation.Taste(); ok {
		if err := taste.TasteValidator(v); err != nil {
			return nil, &ValidationError{Name: "taste", err: fmt.Errorf("ent: validator failed for field \"taste\": %w", err)}
		}
	}

	var (
		err  error
		node *Taste
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TasteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TasteUpdateOne) SaveX(ctx context.Context) *Taste {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TasteUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TasteUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TasteUpdateOne) sqlSave(ctx context.Context) (t *Taste, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taste.Table,
			Columns: taste.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: taste.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Taste.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := tuo.mutation.Taste(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taste.FieldTaste,
		})
	}
	if nodes := tuo.mutation.RemovedEatinghistorysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taste.EatinghistorysTable,
			Columns: []string{taste.EatinghistorysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: eatinghistory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.EatinghistorysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taste.EatinghistorysTable,
			Columns: []string{taste.EatinghistorysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: eatinghistory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	t = &Taste{config: tuo.config}
	_spec.Assign = t.assignValues
	_spec.ScanValues = t.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taste.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}