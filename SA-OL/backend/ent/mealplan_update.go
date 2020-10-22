// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Teeth/app/ent/eatinghistory"
	"github.com/Teeth/app/ent/mealplan"
	"github.com/Teeth/app/ent/predicate"
	"github.com/Teeth/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// MealplanUpdate is the builder for updating Mealplan entities.
type MealplanUpdate struct {
	config
	hooks      []Hook
	mutation   *MealplanMutation
	predicates []predicate.Mealplan
}

// Where adds a new predicate for the builder.
func (mu *MealplanUpdate) Where(ps ...predicate.Mealplan) *MealplanUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetDate sets the date field.
func (mu *MealplanUpdate) SetDate(s string) *MealplanUpdate {
	mu.mutation.SetDate(s)
	return mu
}

// SetFoodID sets the food_id field.
func (mu *MealplanUpdate) SetFoodID(i int) *MealplanUpdate {
	mu.mutation.ResetFoodID()
	mu.mutation.SetFoodID(i)
	return mu
}

// AddFoodID adds i to food_id.
func (mu *MealplanUpdate) AddFoodID(i int) *MealplanUpdate {
	mu.mutation.AddFoodID(i)
	return mu
}

// SetMealID sets the meal_id field.
func (mu *MealplanUpdate) SetMealID(i int) *MealplanUpdate {
	mu.mutation.ResetMealID()
	mu.mutation.SetMealID(i)
	return mu
}

// AddMealID adds i to meal_id.
func (mu *MealplanUpdate) AddMealID(i int) *MealplanUpdate {
	mu.mutation.AddMealID(i)
	return mu
}

// SetOwnerID sets the owner edge to User by id.
func (mu *MealplanUpdate) SetOwnerID(id int) *MealplanUpdate {
	mu.mutation.SetOwnerID(id)
	return mu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (mu *MealplanUpdate) SetNillableOwnerID(id *int) *MealplanUpdate {
	if id != nil {
		mu = mu.SetOwnerID(*id)
	}
	return mu
}

// SetOwner sets the owner edge to User.
func (mu *MealplanUpdate) SetOwner(u *User) *MealplanUpdate {
	return mu.SetOwnerID(u.ID)
}

// AddEatinghistoryIDs adds the eatinghistorys edge to Eatinghistory by ids.
func (mu *MealplanUpdate) AddEatinghistoryIDs(ids ...int) *MealplanUpdate {
	mu.mutation.AddEatinghistoryIDs(ids...)
	return mu
}

// AddEatinghistorys adds the eatinghistorys edges to Eatinghistory.
func (mu *MealplanUpdate) AddEatinghistorys(e ...*Eatinghistory) *MealplanUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mu.AddEatinghistoryIDs(ids...)
}

// Mutation returns the MealplanMutation object of the builder.
func (mu *MealplanUpdate) Mutation() *MealplanMutation {
	return mu.mutation
}

// ClearOwner clears the owner edge to User.
func (mu *MealplanUpdate) ClearOwner() *MealplanUpdate {
	mu.mutation.ClearOwner()
	return mu
}

// RemoveEatinghistoryIDs removes the eatinghistorys edge to Eatinghistory by ids.
func (mu *MealplanUpdate) RemoveEatinghistoryIDs(ids ...int) *MealplanUpdate {
	mu.mutation.RemoveEatinghistoryIDs(ids...)
	return mu
}

// RemoveEatinghistorys removes eatinghistorys edges to Eatinghistory.
func (mu *MealplanUpdate) RemoveEatinghistorys(e ...*Eatinghistory) *MealplanUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mu.RemoveEatinghistoryIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *MealplanUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := mu.mutation.Date(); ok {
		if err := mealplan.DateValidator(v); err != nil {
			return 0, &ValidationError{Name: "date", err: fmt.Errorf("ent: validator failed for field \"date\": %w", err)}
		}
	}
	if v, ok := mu.mutation.FoodID(); ok {
		if err := mealplan.FoodIDValidator(v); err != nil {
			return 0, &ValidationError{Name: "food_id", err: fmt.Errorf("ent: validator failed for field \"food_id\": %w", err)}
		}
	}
	if v, ok := mu.mutation.MealID(); ok {
		if err := mealplan.MealIDValidator(v); err != nil {
			return 0, &ValidationError{Name: "meal_id", err: fmt.Errorf("ent: validator failed for field \"meal_id\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MealplanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MealplanUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MealplanUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MealplanUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MealplanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mealplan.Table,
			Columns: mealplan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mealplan.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mealplan.FieldDate,
		})
	}
	if value, ok := mu.mutation.FoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mealplan.FieldFoodID,
		})
	}
	if value, ok := mu.mutation.AddedFoodID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mealplan.FieldFoodID,
		})
	}
	if value, ok := mu.mutation.MealID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mealplan.FieldMealID,
		})
	}
	if value, ok := mu.mutation.AddedMealID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mealplan.FieldMealID,
		})
	}
	if mu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mealplan.OwnerTable,
			Columns: []string{mealplan.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mealplan.OwnerTable,
			Columns: []string{mealplan.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := mu.mutation.RemovedEatinghistorysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mealplan.EatinghistorysTable,
			Columns: []string{mealplan.EatinghistorysColumn},
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
	if nodes := mu.mutation.EatinghistorysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mealplan.EatinghistorysTable,
			Columns: []string{mealplan.EatinghistorysColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mealplan.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MealplanUpdateOne is the builder for updating a single Mealplan entity.
type MealplanUpdateOne struct {
	config
	hooks    []Hook
	mutation *MealplanMutation
}

// SetDate sets the date field.
func (muo *MealplanUpdateOne) SetDate(s string) *MealplanUpdateOne {
	muo.mutation.SetDate(s)
	return muo
}

// SetFoodID sets the food_id field.
func (muo *MealplanUpdateOne) SetFoodID(i int) *MealplanUpdateOne {
	muo.mutation.ResetFoodID()
	muo.mutation.SetFoodID(i)
	return muo
}

// AddFoodID adds i to food_id.
func (muo *MealplanUpdateOne) AddFoodID(i int) *MealplanUpdateOne {
	muo.mutation.AddFoodID(i)
	return muo
}

// SetMealID sets the meal_id field.
func (muo *MealplanUpdateOne) SetMealID(i int) *MealplanUpdateOne {
	muo.mutation.ResetMealID()
	muo.mutation.SetMealID(i)
	return muo
}

// AddMealID adds i to meal_id.
func (muo *MealplanUpdateOne) AddMealID(i int) *MealplanUpdateOne {
	muo.mutation.AddMealID(i)
	return muo
}

// SetOwnerID sets the owner edge to User by id.
func (muo *MealplanUpdateOne) SetOwnerID(id int) *MealplanUpdateOne {
	muo.mutation.SetOwnerID(id)
	return muo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (muo *MealplanUpdateOne) SetNillableOwnerID(id *int) *MealplanUpdateOne {
	if id != nil {
		muo = muo.SetOwnerID(*id)
	}
	return muo
}

// SetOwner sets the owner edge to User.
func (muo *MealplanUpdateOne) SetOwner(u *User) *MealplanUpdateOne {
	return muo.SetOwnerID(u.ID)
}

// AddEatinghistoryIDs adds the eatinghistorys edge to Eatinghistory by ids.
func (muo *MealplanUpdateOne) AddEatinghistoryIDs(ids ...int) *MealplanUpdateOne {
	muo.mutation.AddEatinghistoryIDs(ids...)
	return muo
}

// AddEatinghistorys adds the eatinghistorys edges to Eatinghistory.
func (muo *MealplanUpdateOne) AddEatinghistorys(e ...*Eatinghistory) *MealplanUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return muo.AddEatinghistoryIDs(ids...)
}

// Mutation returns the MealplanMutation object of the builder.
func (muo *MealplanUpdateOne) Mutation() *MealplanMutation {
	return muo.mutation
}

// ClearOwner clears the owner edge to User.
func (muo *MealplanUpdateOne) ClearOwner() *MealplanUpdateOne {
	muo.mutation.ClearOwner()
	return muo
}

// RemoveEatinghistoryIDs removes the eatinghistorys edge to Eatinghistory by ids.
func (muo *MealplanUpdateOne) RemoveEatinghistoryIDs(ids ...int) *MealplanUpdateOne {
	muo.mutation.RemoveEatinghistoryIDs(ids...)
	return muo
}

// RemoveEatinghistorys removes eatinghistorys edges to Eatinghistory.
func (muo *MealplanUpdateOne) RemoveEatinghistorys(e ...*Eatinghistory) *MealplanUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return muo.RemoveEatinghistoryIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *MealplanUpdateOne) Save(ctx context.Context) (*Mealplan, error) {
	if v, ok := muo.mutation.Date(); ok {
		if err := mealplan.DateValidator(v); err != nil {
			return nil, &ValidationError{Name: "date", err: fmt.Errorf("ent: validator failed for field \"date\": %w", err)}
		}
	}
	if v, ok := muo.mutation.FoodID(); ok {
		if err := mealplan.FoodIDValidator(v); err != nil {
			return nil, &ValidationError{Name: "food_id", err: fmt.Errorf("ent: validator failed for field \"food_id\": %w", err)}
		}
	}
	if v, ok := muo.mutation.MealID(); ok {
		if err := mealplan.MealIDValidator(v); err != nil {
			return nil, &ValidationError{Name: "meal_id", err: fmt.Errorf("ent: validator failed for field \"meal_id\": %w", err)}
		}
	}

	var (
		err  error
		node *Mealplan
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MealplanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MealplanUpdateOne) SaveX(ctx context.Context) *Mealplan {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *MealplanUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MealplanUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MealplanUpdateOne) sqlSave(ctx context.Context) (m *Mealplan, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mealplan.Table,
			Columns: mealplan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mealplan.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Mealplan.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mealplan.FieldDate,
		})
	}
	if value, ok := muo.mutation.FoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mealplan.FieldFoodID,
		})
	}
	if value, ok := muo.mutation.AddedFoodID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mealplan.FieldFoodID,
		})
	}
	if value, ok := muo.mutation.MealID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mealplan.FieldMealID,
		})
	}
	if value, ok := muo.mutation.AddedMealID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mealplan.FieldMealID,
		})
	}
	if muo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mealplan.OwnerTable,
			Columns: []string{mealplan.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mealplan.OwnerTable,
			Columns: []string{mealplan.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := muo.mutation.RemovedEatinghistorysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mealplan.EatinghistorysTable,
			Columns: []string{mealplan.EatinghistorysColumn},
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
	if nodes := muo.mutation.EatinghistorysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mealplan.EatinghistorysTable,
			Columns: []string{mealplan.EatinghistorysColumn},
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
	m = &Mealplan{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mealplan.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}
