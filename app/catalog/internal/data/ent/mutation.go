// Code generated by ent, DO NOT EDIT.

package ent

import (
	"beer-shop/app/catalog/internal/biz"
	"beer-shop/app/catalog/internal/data/ent/beer"
	"beer-shop/app/catalog/internal/data/ent/predicate"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeBeer = "Beer"
)

// BeerMutation represents an operation that mutates the Beer nodes in the graph.
type BeerMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	name          *string
	description   *string
	count         *int64
	addcount      *int64
	price         *int64
	addprice      *int64
	images        *[]biz.Image
	appendimages  []biz.Image
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Beer, error)
	predicates    []predicate.Beer
}

var _ ent.Mutation = (*BeerMutation)(nil)

// beerOption allows management of the mutation configuration using functional options.
type beerOption func(*BeerMutation)

// newBeerMutation creates new mutation for the Beer entity.
func newBeerMutation(c config, op Op, opts ...beerOption) *BeerMutation {
	m := &BeerMutation{
		config:        c,
		op:            op,
		typ:           TypeBeer,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withBeerID sets the ID field of the mutation.
func withBeerID(id int64) beerOption {
	return func(m *BeerMutation) {
		var (
			err   error
			once  sync.Once
			value *Beer
		)
		m.oldValue = func(ctx context.Context) (*Beer, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Beer.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withBeer sets the old Beer of the mutation.
func withBeer(node *Beer) beerOption {
	return func(m *BeerMutation) {
		m.oldValue = func(context.Context) (*Beer, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m BeerMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m BeerMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Beer entities.
func (m *BeerMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *BeerMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *BeerMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Beer.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *BeerMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *BeerMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Beer entity.
// If the Beer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BeerMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *BeerMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *BeerMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *BeerMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Beer entity.
// If the Beer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BeerMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *BeerMutation) ResetDescription() {
	m.description = nil
}

// SetCount sets the "count" field.
func (m *BeerMutation) SetCount(i int64) {
	m.count = &i
	m.addcount = nil
}

// Count returns the value of the "count" field in the mutation.
func (m *BeerMutation) Count() (r int64, exists bool) {
	v := m.count
	if v == nil {
		return
	}
	return *v, true
}

// OldCount returns the old "count" field's value of the Beer entity.
// If the Beer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BeerMutation) OldCount(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCount: %w", err)
	}
	return oldValue.Count, nil
}

// AddCount adds i to the "count" field.
func (m *BeerMutation) AddCount(i int64) {
	if m.addcount != nil {
		*m.addcount += i
	} else {
		m.addcount = &i
	}
}

// AddedCount returns the value that was added to the "count" field in this mutation.
func (m *BeerMutation) AddedCount() (r int64, exists bool) {
	v := m.addcount
	if v == nil {
		return
	}
	return *v, true
}

// ResetCount resets all changes to the "count" field.
func (m *BeerMutation) ResetCount() {
	m.count = nil
	m.addcount = nil
}

// SetPrice sets the "price" field.
func (m *BeerMutation) SetPrice(i int64) {
	m.price = &i
	m.addprice = nil
}

// Price returns the value of the "price" field in the mutation.
func (m *BeerMutation) Price() (r int64, exists bool) {
	v := m.price
	if v == nil {
		return
	}
	return *v, true
}

// OldPrice returns the old "price" field's value of the Beer entity.
// If the Beer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BeerMutation) OldPrice(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPrice is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPrice requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPrice: %w", err)
	}
	return oldValue.Price, nil
}

// AddPrice adds i to the "price" field.
func (m *BeerMutation) AddPrice(i int64) {
	if m.addprice != nil {
		*m.addprice += i
	} else {
		m.addprice = &i
	}
}

// AddedPrice returns the value that was added to the "price" field in this mutation.
func (m *BeerMutation) AddedPrice() (r int64, exists bool) {
	v := m.addprice
	if v == nil {
		return
	}
	return *v, true
}

// ResetPrice resets all changes to the "price" field.
func (m *BeerMutation) ResetPrice() {
	m.price = nil
	m.addprice = nil
}

// SetImages sets the "images" field.
func (m *BeerMutation) SetImages(b []biz.Image) {
	m.images = &b
	m.appendimages = nil
}

// Images returns the value of the "images" field in the mutation.
func (m *BeerMutation) Images() (r []biz.Image, exists bool) {
	v := m.images
	if v == nil {
		return
	}
	return *v, true
}

// OldImages returns the old "images" field's value of the Beer entity.
// If the Beer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BeerMutation) OldImages(ctx context.Context) (v []biz.Image, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldImages is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldImages requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldImages: %w", err)
	}
	return oldValue.Images, nil
}

// AppendImages adds b to the "images" field.
func (m *BeerMutation) AppendImages(b []biz.Image) {
	m.appendimages = append(m.appendimages, b...)
}

// AppendedImages returns the list of values that were appended to the "images" field in this mutation.
func (m *BeerMutation) AppendedImages() ([]biz.Image, bool) {
	if len(m.appendimages) == 0 {
		return nil, false
	}
	return m.appendimages, true
}

// ResetImages resets all changes to the "images" field.
func (m *BeerMutation) ResetImages() {
	m.images = nil
	m.appendimages = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *BeerMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *BeerMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Beer entity.
// If the Beer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BeerMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *BeerMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *BeerMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *BeerMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Beer entity.
// If the Beer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BeerMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *BeerMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the BeerMutation builder.
func (m *BeerMutation) Where(ps ...predicate.Beer) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the BeerMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *BeerMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Beer, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *BeerMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *BeerMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Beer).
func (m *BeerMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *BeerMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.name != nil {
		fields = append(fields, beer.FieldName)
	}
	if m.description != nil {
		fields = append(fields, beer.FieldDescription)
	}
	if m.count != nil {
		fields = append(fields, beer.FieldCount)
	}
	if m.price != nil {
		fields = append(fields, beer.FieldPrice)
	}
	if m.images != nil {
		fields = append(fields, beer.FieldImages)
	}
	if m.created_at != nil {
		fields = append(fields, beer.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, beer.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *BeerMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case beer.FieldName:
		return m.Name()
	case beer.FieldDescription:
		return m.Description()
	case beer.FieldCount:
		return m.Count()
	case beer.FieldPrice:
		return m.Price()
	case beer.FieldImages:
		return m.Images()
	case beer.FieldCreatedAt:
		return m.CreatedAt()
	case beer.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *BeerMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case beer.FieldName:
		return m.OldName(ctx)
	case beer.FieldDescription:
		return m.OldDescription(ctx)
	case beer.FieldCount:
		return m.OldCount(ctx)
	case beer.FieldPrice:
		return m.OldPrice(ctx)
	case beer.FieldImages:
		return m.OldImages(ctx)
	case beer.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case beer.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Beer field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BeerMutation) SetField(name string, value ent.Value) error {
	switch name {
	case beer.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case beer.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case beer.FieldCount:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCount(v)
		return nil
	case beer.FieldPrice:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPrice(v)
		return nil
	case beer.FieldImages:
		v, ok := value.([]biz.Image)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetImages(v)
		return nil
	case beer.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case beer.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Beer field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *BeerMutation) AddedFields() []string {
	var fields []string
	if m.addcount != nil {
		fields = append(fields, beer.FieldCount)
	}
	if m.addprice != nil {
		fields = append(fields, beer.FieldPrice)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *BeerMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case beer.FieldCount:
		return m.AddedCount()
	case beer.FieldPrice:
		return m.AddedPrice()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BeerMutation) AddField(name string, value ent.Value) error {
	switch name {
	case beer.FieldCount:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCount(v)
		return nil
	case beer.FieldPrice:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPrice(v)
		return nil
	}
	return fmt.Errorf("unknown Beer numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *BeerMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *BeerMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *BeerMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Beer nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *BeerMutation) ResetField(name string) error {
	switch name {
	case beer.FieldName:
		m.ResetName()
		return nil
	case beer.FieldDescription:
		m.ResetDescription()
		return nil
	case beer.FieldCount:
		m.ResetCount()
		return nil
	case beer.FieldPrice:
		m.ResetPrice()
		return nil
	case beer.FieldImages:
		m.ResetImages()
		return nil
	case beer.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case beer.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Beer field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *BeerMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *BeerMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *BeerMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *BeerMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *BeerMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *BeerMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *BeerMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Beer unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *BeerMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Beer edge %s", name)
}
