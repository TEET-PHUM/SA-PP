// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Teeth/app/ent/eatinghistory"
	"github.com/Teeth/app/ent/foodmenu"
	"github.com/Teeth/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// FoodmenuCreate is the builder for creating a Foodmenu entity.
type FoodmenuCreate struct {
	config
	mutation *FoodmenuMutation
	hooks    []Hook
}

// SetName sets the name field.
func (fc *FoodmenuCreate) SetName(s string) *FoodmenuCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetMenuing sets the menuing field.
func (fc *FoodmenuCreate) SetMenuing(s string) *FoodmenuCreate {
	fc.mutation.SetMenuing(s)
	return fc
}

// SetAddedtime sets the addedtime field.
func (fc *FoodmenuCreate) SetAddedtime(s string) *FoodmenuCreate {
	fc.mutation.SetAddedtime(s)
	return fc
}

// SetGroupid sets the groupid field.
func (fc *FoodmenuCreate) SetGroupid(i int) *FoodmenuCreate {
	fc.mutation.SetGroupid(i)
	return fc
}

// SetTypeid sets the typeid field.
func (fc *FoodmenuCreate) SetTypeid(i int) *FoodmenuCreate {
	fc.mutation.SetTypeid(i)
	return fc
}

// SetCalories sets the calories field.
func (fc *FoodmenuCreate) SetCalories(i int) *FoodmenuCreate {
	fc.mutation.SetCalories(i)
	return fc
}

// SetOwnerID sets the owner edge to User by id.
func (fc *FoodmenuCreate) SetOwnerID(id int) *FoodmenuCreate {
	fc.mutation.SetOwnerID(id)
	return fc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (fc *FoodmenuCreate) SetNillableOwnerID(id *int) *FoodmenuCreate {
	if id != nil {
		fc = fc.SetOwnerID(*id)
	}
	return fc
}

// SetOwner sets the owner edge to User.
func (fc *FoodmenuCreate) SetOwner(u *User) *FoodmenuCreate {
	return fc.SetOwnerID(u.ID)
}

// AddEatinghistoryIDs adds the eatinghistorys edge to Eatinghistory by ids.
func (fc *FoodmenuCreate) AddEatinghistoryIDs(ids ...int) *FoodmenuCreate {
	fc.mutation.AddEatinghistoryIDs(ids...)
	return fc
}

// AddEatinghistorys adds the eatinghistorys edges to Eatinghistory.
func (fc *FoodmenuCreate) AddEatinghistorys(e ...*Eatinghistory) *FoodmenuCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fc.AddEatinghistoryIDs(ids...)
}

// Mutation returns the FoodmenuMutation object of the builder.
func (fc *FoodmenuCreate) Mutation() *FoodmenuMutation {
	return fc.mutation
}

// Save creates the Foodmenu in the database.
func (fc *FoodmenuCreate) Save(ctx context.Context) (*Foodmenu, error) {
	if _, ok := fc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := fc.mutation.Name(); ok {
		if err := foodmenu.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := fc.mutation.Menuing(); !ok {
		return nil, &ValidationError{Name: "menuing", err: errors.New("ent: missing required field \"menuing\"")}
	}
	if v, ok := fc.mutation.Menuing(); ok {
		if err := foodmenu.MenuingValidator(v); err != nil {
			return nil, &ValidationError{Name: "menuing", err: fmt.Errorf("ent: validator failed for field \"menuing\": %w", err)}
		}
	}
	if _, ok := fc.mutation.Addedtime(); !ok {
		return nil, &ValidationError{Name: "addedtime", err: errors.New("ent: missing required field \"addedtime\"")}
	}
	if v, ok := fc.mutation.Addedtime(); ok {
		if err := foodmenu.AddedtimeValidator(v); err != nil {
			return nil, &ValidationError{Name: "addedtime", err: fmt.Errorf("ent: validator failed for field \"addedtime\": %w", err)}
		}
	}
	if _, ok := fc.mutation.Groupid(); !ok {
		return nil, &ValidationError{Name: "groupid", err: errors.New("ent: missing required field \"groupid\"")}
	}
	if v, ok := fc.mutation.Groupid(); ok {
		if err := foodmenu.GroupidValidator(v); err != nil {
			return nil, &ValidationError{Name: "groupid", err: fmt.Errorf("ent: validator failed for field \"groupid\": %w", err)}
		}
	}
	if _, ok := fc.mutation.Typeid(); !ok {
		return nil, &ValidationError{Name: "typeid", err: errors.New("ent: missing required field \"typeid\"")}
	}
	if v, ok := fc.mutation.Typeid(); ok {
		if err := foodmenu.TypeidValidator(v); err != nil {
			return nil, &ValidationError{Name: "typeid", err: fmt.Errorf("ent: validator failed for field \"typeid\": %w", err)}
		}
	}
	if _, ok := fc.mutation.Calories(); !ok {
		return nil, &ValidationError{Name: "calories", err: errors.New("ent: missing required field \"calories\"")}
	}
	if v, ok := fc.mutation.Calories(); ok {
		if err := foodmenu.CaloriesValidator(v); err != nil {
			return nil, &ValidationError{Name: "calories", err: fmt.Errorf("ent: validator failed for field \"calories\": %w", err)}
		}
	}
	var (
		err  error
		node *Foodmenu
	)
	if len(fc.hooks) == 0 {
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FoodmenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FoodmenuCreate) SaveX(ctx context.Context) *Foodmenu {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FoodmenuCreate) sqlSave(ctx context.Context) (*Foodmenu, error) {
	f, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	f.ID = int(id)
	return f, nil
}

func (fc *FoodmenuCreate) createSpec() (*Foodmenu, *sqlgraph.CreateSpec) {
	var (
		f     = &Foodmenu{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: foodmenu.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: foodmenu.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: foodmenu.FieldName,
		})
		f.Name = value
	}
	if value, ok := fc.mutation.Menuing(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: foodmenu.FieldMenuing,
		})
		f.Menuing = value
	}
	if value, ok := fc.mutation.Addedtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: foodmenu.FieldAddedtime,
		})
		f.Addedtime = value
	}
	if value, ok := fc.mutation.Groupid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: foodmenu.FieldGroupid,
		})
		f.Groupid = value
	}
	if value, ok := fc.mutation.Typeid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: foodmenu.FieldTypeid,
		})
		f.Typeid = value
	}
	if value, ok := fc.mutation.Calories(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: foodmenu.FieldCalories,
		})
		f.Calories = value
	}
	if nodes := fc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodmenu.OwnerTable,
			Columns: []string{foodmenu.OwnerColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.EatinghistorysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   foodmenu.EatinghistorysTable,
			Columns: []string{foodmenu.EatinghistorysColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return f, _spec
}