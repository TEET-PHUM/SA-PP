// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/Teeth/app/ent/migrate"

	"github.com/Teeth/app/ent/eatinghistory"
	"github.com/Teeth/app/ent/foodmenu"
	"github.com/Teeth/app/ent/mealplan"
	"github.com/Teeth/app/ent/taste"
	"github.com/Teeth/app/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Eatinghistory is the client for interacting with the Eatinghistory builders.
	Eatinghistory *EatinghistoryClient
	// Foodmenu is the client for interacting with the Foodmenu builders.
	Foodmenu *FoodmenuClient
	// Mealplan is the client for interacting with the Mealplan builders.
	Mealplan *MealplanClient
	// Taste is the client for interacting with the Taste builders.
	Taste *TasteClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Eatinghistory = NewEatinghistoryClient(c.config)
	c.Foodmenu = NewFoodmenuClient(c.config)
	c.Mealplan = NewMealplanClient(c.config)
	c.Taste = NewTasteClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		Eatinghistory: NewEatinghistoryClient(cfg),
		Foodmenu:      NewFoodmenuClient(cfg),
		Mealplan:      NewMealplanClient(cfg),
		Taste:         NewTasteClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:        cfg,
		Eatinghistory: NewEatinghistoryClient(cfg),
		Foodmenu:      NewFoodmenuClient(cfg),
		Mealplan:      NewMealplanClient(cfg),
		Taste:         NewTasteClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Eatinghistory.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Eatinghistory.Use(hooks...)
	c.Foodmenu.Use(hooks...)
	c.Mealplan.Use(hooks...)
	c.Taste.Use(hooks...)
	c.User.Use(hooks...)
}

// EatinghistoryClient is a client for the Eatinghistory schema.
type EatinghistoryClient struct {
	config
}

// NewEatinghistoryClient returns a client for the Eatinghistory from the given config.
func NewEatinghistoryClient(c config) *EatinghistoryClient {
	return &EatinghistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `eatinghistory.Hooks(f(g(h())))`.
func (c *EatinghistoryClient) Use(hooks ...Hook) {
	c.hooks.Eatinghistory = append(c.hooks.Eatinghistory, hooks...)
}

// Create returns a create builder for Eatinghistory.
func (c *EatinghistoryClient) Create() *EatinghistoryCreate {
	mutation := newEatinghistoryMutation(c.config, OpCreate)
	return &EatinghistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Eatinghistory.
func (c *EatinghistoryClient) Update() *EatinghistoryUpdate {
	mutation := newEatinghistoryMutation(c.config, OpUpdate)
	return &EatinghistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EatinghistoryClient) UpdateOne(e *Eatinghistory) *EatinghistoryUpdateOne {
	mutation := newEatinghistoryMutation(c.config, OpUpdateOne, withEatinghistory(e))
	return &EatinghistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EatinghistoryClient) UpdateOneID(id int) *EatinghistoryUpdateOne {
	mutation := newEatinghistoryMutation(c.config, OpUpdateOne, withEatinghistoryID(id))
	return &EatinghistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Eatinghistory.
func (c *EatinghistoryClient) Delete() *EatinghistoryDelete {
	mutation := newEatinghistoryMutation(c.config, OpDelete)
	return &EatinghistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EatinghistoryClient) DeleteOne(e *Eatinghistory) *EatinghistoryDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EatinghistoryClient) DeleteOneID(id int) *EatinghistoryDeleteOne {
	builder := c.Delete().Where(eatinghistory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EatinghistoryDeleteOne{builder}
}

// Create returns a query builder for Eatinghistory.
func (c *EatinghistoryClient) Query() *EatinghistoryQuery {
	return &EatinghistoryQuery{config: c.config}
}

// Get returns a Eatinghistory entity by its id.
func (c *EatinghistoryClient) Get(ctx context.Context, id int) (*Eatinghistory, error) {
	return c.Query().Where(eatinghistory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EatinghistoryClient) GetX(ctx context.Context, id int) *Eatinghistory {
	e, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return e
}

// QueryMealplan queries the mealplan edge of a Eatinghistory.
func (c *EatinghistoryClient) QueryMealplan(e *Eatinghistory) *MealplanQuery {
	query := &MealplanQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(eatinghistory.Table, eatinghistory.FieldID, id),
			sqlgraph.To(mealplan.Table, mealplan.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, eatinghistory.MealplanTable, eatinghistory.MealplanColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFoodmenu queries the foodmenu edge of a Eatinghistory.
func (c *EatinghistoryClient) QueryFoodmenu(e *Eatinghistory) *FoodmenuQuery {
	query := &FoodmenuQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(eatinghistory.Table, eatinghistory.FieldID, id),
			sqlgraph.To(foodmenu.Table, foodmenu.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, eatinghistory.FoodmenuTable, eatinghistory.FoodmenuColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTaste queries the taste edge of a Eatinghistory.
func (c *EatinghistoryClient) QueryTaste(e *Eatinghistory) *TasteQuery {
	query := &TasteQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(eatinghistory.Table, eatinghistory.FieldID, id),
			sqlgraph.To(taste.Table, taste.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, eatinghistory.TasteTable, eatinghistory.TasteColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Eatinghistory.
func (c *EatinghistoryClient) QueryUser(e *Eatinghistory) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(eatinghistory.Table, eatinghistory.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, eatinghistory.UserTable, eatinghistory.UserColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EatinghistoryClient) Hooks() []Hook {
	return c.hooks.Eatinghistory
}

// FoodmenuClient is a client for the Foodmenu schema.
type FoodmenuClient struct {
	config
}

// NewFoodmenuClient returns a client for the Foodmenu from the given config.
func NewFoodmenuClient(c config) *FoodmenuClient {
	return &FoodmenuClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `foodmenu.Hooks(f(g(h())))`.
func (c *FoodmenuClient) Use(hooks ...Hook) {
	c.hooks.Foodmenu = append(c.hooks.Foodmenu, hooks...)
}

// Create returns a create builder for Foodmenu.
func (c *FoodmenuClient) Create() *FoodmenuCreate {
	mutation := newFoodmenuMutation(c.config, OpCreate)
	return &FoodmenuCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Foodmenu.
func (c *FoodmenuClient) Update() *FoodmenuUpdate {
	mutation := newFoodmenuMutation(c.config, OpUpdate)
	return &FoodmenuUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FoodmenuClient) UpdateOne(f *Foodmenu) *FoodmenuUpdateOne {
	mutation := newFoodmenuMutation(c.config, OpUpdateOne, withFoodmenu(f))
	return &FoodmenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FoodmenuClient) UpdateOneID(id int) *FoodmenuUpdateOne {
	mutation := newFoodmenuMutation(c.config, OpUpdateOne, withFoodmenuID(id))
	return &FoodmenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Foodmenu.
func (c *FoodmenuClient) Delete() *FoodmenuDelete {
	mutation := newFoodmenuMutation(c.config, OpDelete)
	return &FoodmenuDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FoodmenuClient) DeleteOne(f *Foodmenu) *FoodmenuDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FoodmenuClient) DeleteOneID(id int) *FoodmenuDeleteOne {
	builder := c.Delete().Where(foodmenu.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FoodmenuDeleteOne{builder}
}

// Create returns a query builder for Foodmenu.
func (c *FoodmenuClient) Query() *FoodmenuQuery {
	return &FoodmenuQuery{config: c.config}
}

// Get returns a Foodmenu entity by its id.
func (c *FoodmenuClient) Get(ctx context.Context, id int) (*Foodmenu, error) {
	return c.Query().Where(foodmenu.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FoodmenuClient) GetX(ctx context.Context, id int) *Foodmenu {
	f, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return f
}

// QueryOwner queries the owner edge of a Foodmenu.
func (c *FoodmenuClient) QueryOwner(f *Foodmenu) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(foodmenu.Table, foodmenu.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, foodmenu.OwnerTable, foodmenu.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEatinghistorys queries the eatinghistorys edge of a Foodmenu.
func (c *FoodmenuClient) QueryEatinghistorys(f *Foodmenu) *EatinghistoryQuery {
	query := &EatinghistoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(foodmenu.Table, foodmenu.FieldID, id),
			sqlgraph.To(eatinghistory.Table, eatinghistory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, foodmenu.EatinghistorysTable, foodmenu.EatinghistorysColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FoodmenuClient) Hooks() []Hook {
	return c.hooks.Foodmenu
}

// MealplanClient is a client for the Mealplan schema.
type MealplanClient struct {
	config
}

// NewMealplanClient returns a client for the Mealplan from the given config.
func NewMealplanClient(c config) *MealplanClient {
	return &MealplanClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mealplan.Hooks(f(g(h())))`.
func (c *MealplanClient) Use(hooks ...Hook) {
	c.hooks.Mealplan = append(c.hooks.Mealplan, hooks...)
}

// Create returns a create builder for Mealplan.
func (c *MealplanClient) Create() *MealplanCreate {
	mutation := newMealplanMutation(c.config, OpCreate)
	return &MealplanCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Mealplan.
func (c *MealplanClient) Update() *MealplanUpdate {
	mutation := newMealplanMutation(c.config, OpUpdate)
	return &MealplanUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MealplanClient) UpdateOne(m *Mealplan) *MealplanUpdateOne {
	mutation := newMealplanMutation(c.config, OpUpdateOne, withMealplan(m))
	return &MealplanUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MealplanClient) UpdateOneID(id int) *MealplanUpdateOne {
	mutation := newMealplanMutation(c.config, OpUpdateOne, withMealplanID(id))
	return &MealplanUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Mealplan.
func (c *MealplanClient) Delete() *MealplanDelete {
	mutation := newMealplanMutation(c.config, OpDelete)
	return &MealplanDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MealplanClient) DeleteOne(m *Mealplan) *MealplanDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MealplanClient) DeleteOneID(id int) *MealplanDeleteOne {
	builder := c.Delete().Where(mealplan.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MealplanDeleteOne{builder}
}

// Create returns a query builder for Mealplan.
func (c *MealplanClient) Query() *MealplanQuery {
	return &MealplanQuery{config: c.config}
}

// Get returns a Mealplan entity by its id.
func (c *MealplanClient) Get(ctx context.Context, id int) (*Mealplan, error) {
	return c.Query().Where(mealplan.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MealplanClient) GetX(ctx context.Context, id int) *Mealplan {
	m, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return m
}

// QueryOwner queries the owner edge of a Mealplan.
func (c *MealplanClient) QueryOwner(m *Mealplan) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(mealplan.Table, mealplan.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, mealplan.OwnerTable, mealplan.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEatinghistorys queries the eatinghistorys edge of a Mealplan.
func (c *MealplanClient) QueryEatinghistorys(m *Mealplan) *EatinghistoryQuery {
	query := &EatinghistoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(mealplan.Table, mealplan.FieldID, id),
			sqlgraph.To(eatinghistory.Table, eatinghistory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, mealplan.EatinghistorysTable, mealplan.EatinghistorysColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MealplanClient) Hooks() []Hook {
	return c.hooks.Mealplan
}

// TasteClient is a client for the Taste schema.
type TasteClient struct {
	config
}

// NewTasteClient returns a client for the Taste from the given config.
func NewTasteClient(c config) *TasteClient {
	return &TasteClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `taste.Hooks(f(g(h())))`.
func (c *TasteClient) Use(hooks ...Hook) {
	c.hooks.Taste = append(c.hooks.Taste, hooks...)
}

// Create returns a create builder for Taste.
func (c *TasteClient) Create() *TasteCreate {
	mutation := newTasteMutation(c.config, OpCreate)
	return &TasteCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Taste.
func (c *TasteClient) Update() *TasteUpdate {
	mutation := newTasteMutation(c.config, OpUpdate)
	return &TasteUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TasteClient) UpdateOne(t *Taste) *TasteUpdateOne {
	mutation := newTasteMutation(c.config, OpUpdateOne, withTaste(t))
	return &TasteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TasteClient) UpdateOneID(id int) *TasteUpdateOne {
	mutation := newTasteMutation(c.config, OpUpdateOne, withTasteID(id))
	return &TasteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Taste.
func (c *TasteClient) Delete() *TasteDelete {
	mutation := newTasteMutation(c.config, OpDelete)
	return &TasteDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TasteClient) DeleteOne(t *Taste) *TasteDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TasteClient) DeleteOneID(id int) *TasteDeleteOne {
	builder := c.Delete().Where(taste.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TasteDeleteOne{builder}
}

// Create returns a query builder for Taste.
func (c *TasteClient) Query() *TasteQuery {
	return &TasteQuery{config: c.config}
}

// Get returns a Taste entity by its id.
func (c *TasteClient) Get(ctx context.Context, id int) (*Taste, error) {
	return c.Query().Where(taste.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TasteClient) GetX(ctx context.Context, id int) *Taste {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// QueryEatinghistorys queries the eatinghistorys edge of a Taste.
func (c *TasteClient) QueryEatinghistorys(t *Taste) *EatinghistoryQuery {
	query := &EatinghistoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(taste.Table, taste.FieldID, id),
			sqlgraph.To(eatinghistory.Table, eatinghistory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, taste.EatinghistorysTable, taste.EatinghistorysColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TasteClient) Hooks() []Hook {
	return c.hooks.Taste
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryFoodmenus queries the foodmenus edge of a User.
func (c *UserClient) QueryFoodmenus(u *User) *FoodmenuQuery {
	query := &FoodmenuQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(foodmenu.Table, foodmenu.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.FoodmenusTable, user.FoodmenusColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMealplans queries the mealplans edge of a User.
func (c *UserClient) QueryMealplans(u *User) *MealplanQuery {
	query := &MealplanQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(mealplan.Table, mealplan.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.MealplansTable, user.MealplansColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEatinghistorys queries the eatinghistorys edge of a User.
func (c *UserClient) QueryEatinghistorys(u *User) *EatinghistoryQuery {
	query := &EatinghistoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(eatinghistory.Table, eatinghistory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.EatinghistorysTable, user.EatinghistorysColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}