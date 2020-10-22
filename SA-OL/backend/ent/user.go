// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/Teeth/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Foodmenus holds the value of the foodmenus edge.
	Foodmenus []*Foodmenu
	// Mealplans holds the value of the mealplans edge.
	Mealplans []*Mealplan
	// Eatinghistorys holds the value of the eatinghistorys edge.
	Eatinghistorys []*Eatinghistory
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// FoodmenusOrErr returns the Foodmenus value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FoodmenusOrErr() ([]*Foodmenu, error) {
	if e.loadedTypes[0] {
		return e.Foodmenus, nil
	}
	return nil, &NotLoadedError{edge: "foodmenus"}
}

// MealplansOrErr returns the Mealplans value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MealplansOrErr() ([]*Mealplan, error) {
	if e.loadedTypes[1] {
		return e.Mealplans, nil
	}
	return nil, &NotLoadedError{edge: "mealplans"}
}

// EatinghistorysOrErr returns the Eatinghistorys value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) EatinghistorysOrErr() ([]*Eatinghistory, error) {
	if e.loadedTypes[2] {
		return e.Eatinghistorys, nil
	}
	return nil, &NotLoadedError{edge: "eatinghistorys"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // email
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		u.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[1])
	} else if value.Valid {
		u.Email = value.String
	}
	return nil
}

// QueryFoodmenus queries the foodmenus edge of the User.
func (u *User) QueryFoodmenus() *FoodmenuQuery {
	return (&UserClient{config: u.config}).QueryFoodmenus(u)
}

// QueryMealplans queries the mealplans edge of the User.
func (u *User) QueryMealplans() *MealplanQuery {
	return (&UserClient{config: u.config}).QueryMealplans(u)
}

// QueryEatinghistorys queries the eatinghistorys edge of the User.
func (u *User) QueryEatinghistorys() *EatinghistoryQuery {
	return (&UserClient{config: u.config}).QueryEatinghistorys(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}