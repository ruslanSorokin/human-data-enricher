// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen/person"
)

// PersonCreate is the builder for creating a Person entity.
type PersonCreate struct {
	config
	mutation *PersonMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *PersonCreate) SetCreatedAt(t time.Time) *PersonCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PersonCreate) SetNillableCreatedAt(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PersonCreate) SetUpdatedAt(t time.Time) *PersonCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUpdatedAt(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *PersonCreate) SetDeletedAt(st sql.NullTime) *PersonCreate {
	pc.mutation.SetDeletedAt(st)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *PersonCreate) SetNillableDeletedAt(st *sql.NullTime) *PersonCreate {
	if st != nil {
		pc.SetDeletedAt(*st)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PersonCreate) SetName(s string) *PersonCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetSurname sets the "surname" field.
func (pc *PersonCreate) SetSurname(s string) *PersonCreate {
	pc.mutation.SetSurname(s)
	return pc
}

// SetMiddleName sets the "middle_name" field.
func (pc *PersonCreate) SetMiddleName(ss sql.NullString) *PersonCreate {
	pc.mutation.SetMiddleName(ss)
	return pc
}

// SetNillableMiddleName sets the "middle_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillableMiddleName(ss *sql.NullString) *PersonCreate {
	if ss != nil {
		pc.SetMiddleName(*ss)
	}
	return pc
}

// SetAge sets the "age" field.
func (pc *PersonCreate) SetAge(i int) *PersonCreate {
	pc.mutation.SetAge(i)
	return pc
}

// SetGender sets the "gender" field.
func (pc *PersonCreate) SetGender(s string) *PersonCreate {
	pc.mutation.SetGender(s)
	return pc
}

// SetNationality sets the "nationality" field.
func (pc *PersonCreate) SetNationality(s string) *PersonCreate {
	pc.mutation.SetNationality(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PersonCreate) SetID(u uuid.UUID) *PersonCreate {
	pc.mutation.SetID(u)
	return pc
}

// Mutation returns the PersonMutation object of the builder.
func (pc *PersonCreate) Mutation() *PersonMutation {
	return pc.mutation
}

// Save creates the Person in the database.
func (pc *PersonCreate) Save(ctx context.Context) (*Person, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PersonCreate) SaveX(ctx context.Context) *Person {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PersonCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PersonCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PersonCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := person.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := person.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.DeletedAt(); !ok {
		v := person.DefaultDeletedAt()
		pc.mutation.SetDeletedAt(v)
	}
	if _, ok := pc.mutation.MiddleName(); !ok {
		v := person.DefaultMiddleName()
		pc.mutation.SetMiddleName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PersonCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`gen: missing required field "Person.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := person.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`gen: validator failed for field "Person.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Surname(); !ok {
		return &ValidationError{Name: "surname", err: errors.New(`gen: missing required field "Person.surname"`)}
	}
	if v, ok := pc.mutation.Surname(); ok {
		if err := person.SurnameValidator(v); err != nil {
			return &ValidationError{Name: "surname", err: fmt.Errorf(`gen: validator failed for field "Person.surname": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`gen: missing required field "Person.age"`)}
	}
	if _, ok := pc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`gen: missing required field "Person.gender"`)}
	}
	if v, ok := pc.mutation.Gender(); ok {
		if err := person.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`gen: validator failed for field "Person.gender": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Nationality(); !ok {
		return &ValidationError{Name: "nationality", err: errors.New(`gen: missing required field "Person.nationality"`)}
	}
	if v, ok := pc.mutation.Nationality(); ok {
		if err := person.NationalityValidator(v); err != nil {
			return &ValidationError{Name: "nationality", err: fmt.Errorf(`gen: validator failed for field "Person.nationality": %w`, err)}
		}
	}
	return nil
}

func (pc *PersonCreate) sqlSave(ctx context.Context) (*Person, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PersonCreate) createSpec() (*Person, *sqlgraph.CreateSpec) {
	var (
		_node = &Person{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(person.Table, sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(person.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(person.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.SetField(person.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(person.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Surname(); ok {
		_spec.SetField(person.FieldSurname, field.TypeString, value)
		_node.Surname = value
	}
	if value, ok := pc.mutation.MiddleName(); ok {
		_spec.SetField(person.FieldMiddleName, field.TypeString, value)
		_node.MiddleName = value
	}
	if value, ok := pc.mutation.Age(); ok {
		_spec.SetField(person.FieldAge, field.TypeInt, value)
		_node.Age = value
	}
	if value, ok := pc.mutation.Gender(); ok {
		_spec.SetField(person.FieldGender, field.TypeString, value)
		_node.Gender = value
	}
	if value, ok := pc.mutation.Nationality(); ok {
		_spec.SetField(person.FieldNationality, field.TypeString, value)
		_node.Nationality = value
	}
	return _node, _spec
}

// PersonCreateBulk is the builder for creating many Person entities in bulk.
type PersonCreateBulk struct {
	config
	err      error
	builders []*PersonCreate
}

// Save creates the Person entities in the database.
func (pcb *PersonCreateBulk) Save(ctx context.Context) ([]*Person, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Person, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PersonCreateBulk) SaveX(ctx context.Context) []*Person {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PersonCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PersonCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
