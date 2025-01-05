// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/twiglab/crm/card/orm/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/twiglab/crm/card/orm/ent/card"
	"github.com/twiglab/crm/card/orm/ent/chargerecord"

	stdsql "database/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Card is the client for interacting with the Card builders.
	Card *CardClient
	// ChargeRecord is the client for interacting with the ChargeRecord builders.
	ChargeRecord *ChargeRecordClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Card = NewCardClient(c.config)
	c.ChargeRecord = NewChargeRecordClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Card:         NewCardClient(cfg),
		ChargeRecord: NewChargeRecordClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Card:         NewCardClient(cfg),
		ChargeRecord: NewChargeRecordClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Card.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.Card.Use(hooks...)
	c.ChargeRecord.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Card.Intercept(interceptors...)
	c.ChargeRecord.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CardMutation:
		return c.Card.mutate(ctx, m)
	case *ChargeRecordMutation:
		return c.ChargeRecord.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CardClient is a client for the Card schema.
type CardClient struct {
	config
}

// NewCardClient returns a client for the Card from the given config.
func NewCardClient(c config) *CardClient {
	return &CardClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `card.Hooks(f(g(h())))`.
func (c *CardClient) Use(hooks ...Hook) {
	c.hooks.Card = append(c.hooks.Card, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `card.Intercept(f(g(h())))`.
func (c *CardClient) Intercept(interceptors ...Interceptor) {
	c.inters.Card = append(c.inters.Card, interceptors...)
}

// Create returns a builder for creating a Card entity.
func (c *CardClient) Create() *CardCreate {
	mutation := newCardMutation(c.config, OpCreate)
	return &CardCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Card entities.
func (c *CardClient) CreateBulk(builders ...*CardCreate) *CardCreateBulk {
	return &CardCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CardClient) MapCreateBulk(slice any, setFunc func(*CardCreate, int)) *CardCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CardCreateBulk{err: fmt.Errorf("calling to CardClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CardCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CardCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Card.
func (c *CardClient) Update() *CardUpdate {
	mutation := newCardMutation(c.config, OpUpdate)
	return &CardUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CardClient) UpdateOne(ca *Card) *CardUpdateOne {
	mutation := newCardMutation(c.config, OpUpdateOne, withCard(ca))
	return &CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CardClient) UpdateOneID(id int) *CardUpdateOne {
	mutation := newCardMutation(c.config, OpUpdateOne, withCardID(id))
	return &CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Card.
func (c *CardClient) Delete() *CardDelete {
	mutation := newCardMutation(c.config, OpDelete)
	return &CardDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CardClient) DeleteOne(ca *Card) *CardDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CardClient) DeleteOneID(id int) *CardDeleteOne {
	builder := c.Delete().Where(card.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CardDeleteOne{builder}
}

// Query returns a query builder for Card.
func (c *CardClient) Query() *CardQuery {
	return &CardQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCard},
		inters: c.Interceptors(),
	}
}

// Get returns a Card entity by its id.
func (c *CardClient) Get(ctx context.Context, id int) (*Card, error) {
	return c.Query().Where(card.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CardClient) GetX(ctx context.Context, id int) *Card {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CardClient) Hooks() []Hook {
	return c.hooks.Card
}

// Interceptors returns the client interceptors.
func (c *CardClient) Interceptors() []Interceptor {
	return c.inters.Card
}

func (c *CardClient) mutate(ctx context.Context, m *CardMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CardCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CardUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CardDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Card mutation op: %q", m.Op())
	}
}

// ChargeRecordClient is a client for the ChargeRecord schema.
type ChargeRecordClient struct {
	config
}

// NewChargeRecordClient returns a client for the ChargeRecord from the given config.
func NewChargeRecordClient(c config) *ChargeRecordClient {
	return &ChargeRecordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chargerecord.Hooks(f(g(h())))`.
func (c *ChargeRecordClient) Use(hooks ...Hook) {
	c.hooks.ChargeRecord = append(c.hooks.ChargeRecord, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chargerecord.Intercept(f(g(h())))`.
func (c *ChargeRecordClient) Intercept(interceptors ...Interceptor) {
	c.inters.ChargeRecord = append(c.inters.ChargeRecord, interceptors...)
}

// Create returns a builder for creating a ChargeRecord entity.
func (c *ChargeRecordClient) Create() *ChargeRecordCreate {
	mutation := newChargeRecordMutation(c.config, OpCreate)
	return &ChargeRecordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChargeRecord entities.
func (c *ChargeRecordClient) CreateBulk(builders ...*ChargeRecordCreate) *ChargeRecordCreateBulk {
	return &ChargeRecordCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ChargeRecordClient) MapCreateBulk(slice any, setFunc func(*ChargeRecordCreate, int)) *ChargeRecordCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ChargeRecordCreateBulk{err: fmt.Errorf("calling to ChargeRecordClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ChargeRecordCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ChargeRecordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChargeRecord.
func (c *ChargeRecordClient) Update() *ChargeRecordUpdate {
	mutation := newChargeRecordMutation(c.config, OpUpdate)
	return &ChargeRecordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChargeRecordClient) UpdateOne(cr *ChargeRecord) *ChargeRecordUpdateOne {
	mutation := newChargeRecordMutation(c.config, OpUpdateOne, withChargeRecord(cr))
	return &ChargeRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChargeRecordClient) UpdateOneID(id int) *ChargeRecordUpdateOne {
	mutation := newChargeRecordMutation(c.config, OpUpdateOne, withChargeRecordID(id))
	return &ChargeRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChargeRecord.
func (c *ChargeRecordClient) Delete() *ChargeRecordDelete {
	mutation := newChargeRecordMutation(c.config, OpDelete)
	return &ChargeRecordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChargeRecordClient) DeleteOne(cr *ChargeRecord) *ChargeRecordDeleteOne {
	return c.DeleteOneID(cr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChargeRecordClient) DeleteOneID(id int) *ChargeRecordDeleteOne {
	builder := c.Delete().Where(chargerecord.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChargeRecordDeleteOne{builder}
}

// Query returns a query builder for ChargeRecord.
func (c *ChargeRecordClient) Query() *ChargeRecordQuery {
	return &ChargeRecordQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChargeRecord},
		inters: c.Interceptors(),
	}
}

// Get returns a ChargeRecord entity by its id.
func (c *ChargeRecordClient) Get(ctx context.Context, id int) (*ChargeRecord, error) {
	return c.Query().Where(chargerecord.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChargeRecordClient) GetX(ctx context.Context, id int) *ChargeRecord {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ChargeRecordClient) Hooks() []Hook {
	return c.hooks.ChargeRecord
}

// Interceptors returns the client interceptors.
func (c *ChargeRecordClient) Interceptors() []Interceptor {
	return c.inters.ChargeRecord
}

func (c *ChargeRecordClient) mutate(ctx context.Context, m *ChargeRecordMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChargeRecordCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChargeRecordUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChargeRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChargeRecordDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ChargeRecord mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Card, ChargeRecord []ent.Hook
	}
	inters struct {
		Card, ChargeRecord []ent.Interceptor
	}
)

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
