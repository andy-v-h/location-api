// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/location-api/internal/ent/generated/location"
	"go.infratographer.com/location-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeLocation = "Location"
)

// LocationMutation represents an operation that mutates the Location nodes in the graph.
type LocationMutation struct {
	config
	op            Op
	typ           string
	id            *gidx.PrefixedID
	created_at    *time.Time
	updated_at    *time.Time
	name          *string
	description   *string
	owner_id      *gidx.PrefixedID
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Location, error)
	predicates    []predicate.Location
}

var _ ent.Mutation = (*LocationMutation)(nil)

// locationOption allows management of the mutation configuration using functional options.
type locationOption func(*LocationMutation)

// newLocationMutation creates new mutation for the Location entity.
func newLocationMutation(c config, op Op, opts ...locationOption) *LocationMutation {
	m := &LocationMutation{
		config:        c,
		op:            op,
		typ:           TypeLocation,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLocationID sets the ID field of the mutation.
func withLocationID(id gidx.PrefixedID) locationOption {
	return func(m *LocationMutation) {
		var (
			err   error
			once  sync.Once
			value *Location
		)
		m.oldValue = func(ctx context.Context) (*Location, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Location.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLocation sets the old Location of the mutation.
func withLocation(node *Location) locationOption {
	return func(m *LocationMutation) {
		m.oldValue = func(context.Context) (*Location, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LocationMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LocationMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("generated: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Location entities.
func (m *LocationMutation) SetID(id gidx.PrefixedID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *LocationMutation) ID() (id gidx.PrefixedID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *LocationMutation) IDs(ctx context.Context) ([]gidx.PrefixedID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []gidx.PrefixedID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Location.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *LocationMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *LocationMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Location entity.
// If the Location object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LocationMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *LocationMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *LocationMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *LocationMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Location entity.
// If the Location object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LocationMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *LocationMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetName sets the "name" field.
func (m *LocationMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *LocationMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Location entity.
// If the Location object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LocationMutation) OldName(ctx context.Context) (v string, err error) {
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
func (m *LocationMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *LocationMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *LocationMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Location entity.
// If the Location object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LocationMutation) OldDescription(ctx context.Context) (v string, err error) {
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

// ClearDescription clears the value of the "description" field.
func (m *LocationMutation) ClearDescription() {
	m.description = nil
	m.clearedFields[location.FieldDescription] = struct{}{}
}

// DescriptionCleared returns if the "description" field was cleared in this mutation.
func (m *LocationMutation) DescriptionCleared() bool {
	_, ok := m.clearedFields[location.FieldDescription]
	return ok
}

// ResetDescription resets all changes to the "description" field.
func (m *LocationMutation) ResetDescription() {
	m.description = nil
	delete(m.clearedFields, location.FieldDescription)
}

// SetOwnerID sets the "owner_id" field.
func (m *LocationMutation) SetOwnerID(gi gidx.PrefixedID) {
	m.owner_id = &gi
}

// OwnerID returns the value of the "owner_id" field in the mutation.
func (m *LocationMutation) OwnerID() (r gidx.PrefixedID, exists bool) {
	v := m.owner_id
	if v == nil {
		return
	}
	return *v, true
}

// OldOwnerID returns the old "owner_id" field's value of the Location entity.
// If the Location object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LocationMutation) OldOwnerID(ctx context.Context) (v gidx.PrefixedID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOwnerID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOwnerID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwnerID: %w", err)
	}
	return oldValue.OwnerID, nil
}

// ResetOwnerID resets all changes to the "owner_id" field.
func (m *LocationMutation) ResetOwnerID() {
	m.owner_id = nil
}

// Where appends a list predicates to the LocationMutation builder.
func (m *LocationMutation) Where(ps ...predicate.Location) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the LocationMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *LocationMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Location, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *LocationMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *LocationMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Location).
func (m *LocationMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *LocationMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, location.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, location.FieldUpdatedAt)
	}
	if m.name != nil {
		fields = append(fields, location.FieldName)
	}
	if m.description != nil {
		fields = append(fields, location.FieldDescription)
	}
	if m.owner_id != nil {
		fields = append(fields, location.FieldOwnerID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *LocationMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case location.FieldCreatedAt:
		return m.CreatedAt()
	case location.FieldUpdatedAt:
		return m.UpdatedAt()
	case location.FieldName:
		return m.Name()
	case location.FieldDescription:
		return m.Description()
	case location.FieldOwnerID:
		return m.OwnerID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *LocationMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case location.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case location.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case location.FieldName:
		return m.OldName(ctx)
	case location.FieldDescription:
		return m.OldDescription(ctx)
	case location.FieldOwnerID:
		return m.OldOwnerID(ctx)
	}
	return nil, fmt.Errorf("unknown Location field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LocationMutation) SetField(name string, value ent.Value) error {
	switch name {
	case location.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case location.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case location.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case location.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case location.FieldOwnerID:
		v, ok := value.(gidx.PrefixedID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwnerID(v)
		return nil
	}
	return fmt.Errorf("unknown Location field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *LocationMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *LocationMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LocationMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Location numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *LocationMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(location.FieldDescription) {
		fields = append(fields, location.FieldDescription)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *LocationMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *LocationMutation) ClearField(name string) error {
	switch name {
	case location.FieldDescription:
		m.ClearDescription()
		return nil
	}
	return fmt.Errorf("unknown Location nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *LocationMutation) ResetField(name string) error {
	switch name {
	case location.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case location.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case location.FieldName:
		m.ResetName()
		return nil
	case location.FieldDescription:
		m.ResetDescription()
		return nil
	case location.FieldOwnerID:
		m.ResetOwnerID()
		return nil
	}
	return fmt.Errorf("unknown Location field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *LocationMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *LocationMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *LocationMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *LocationMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *LocationMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *LocationMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *LocationMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Location unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *LocationMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Location edge %s", name)
}
