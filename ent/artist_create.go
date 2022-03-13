// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Zubayear/song-ql/ent/artist"
	"github.com/Zubayear/song-ql/ent/song"
	"github.com/google/uuid"
)

// ArtistCreate is the builder for creating a Artist entity.
type ArtistCreate struct {
	config
	mutation *ArtistMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *ArtistCreate) SetName(s string) *ArtistCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetAge sets the "age" field.
func (ac *ArtistCreate) SetAge(i int) *ArtistCreate {
	ac.mutation.SetAge(i)
	return ac
}

// SetID sets the "id" field.
func (ac *ArtistCreate) SetID(u uuid.UUID) *ArtistCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *ArtistCreate) SetNillableID(u *uuid.UUID) *ArtistCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetSongsID sets the "songs" edge to the Song entity by ID.
func (ac *ArtistCreate) SetSongsID(id uuid.UUID) *ArtistCreate {
	ac.mutation.SetSongsID(id)
	return ac
}

// SetNillableSongsID sets the "songs" edge to the Song entity by ID if the given value is not nil.
func (ac *ArtistCreate) SetNillableSongsID(id *uuid.UUID) *ArtistCreate {
	if id != nil {
		ac = ac.SetSongsID(*id)
	}
	return ac
}

// SetSongs sets the "songs" edge to the Song entity.
func (ac *ArtistCreate) SetSongs(s *Song) *ArtistCreate {
	return ac.SetSongsID(s.ID)
}

// Mutation returns the ArtistMutation object of the builder.
func (ac *ArtistCreate) Mutation() *ArtistMutation {
	return ac.mutation
}

// Save creates the Artist in the database.
func (ac *ArtistCreate) Save(ctx context.Context) (*Artist, error) {
	var (
		err  error
		node *Artist
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArtistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArtistCreate) SaveX(ctx context.Context) *Artist {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArtistCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArtistCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArtistCreate) defaults() {
	if _, ok := ac.mutation.ID(); !ok {
		v := artist.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArtistCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Artist.name"`)}
	}
	if _, ok := ac.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Artist.age"`)}
	}
	return nil
}

func (ac *ArtistCreate) sqlSave(ctx context.Context) (*Artist, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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
	return _node, nil
}

func (ac *ArtistCreate) createSpec() (*Artist, *sqlgraph.CreateSpec) {
	var (
		_node = &Artist{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: artist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: artist.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: artist.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: artist.FieldAge,
		})
		_node.Age = value
	}
	if nodes := ac.mutation.SongsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artist.SongsTable,
			Columns: []string{artist.SongsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: song.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.song_artists = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ArtistCreateBulk is the builder for creating many Artist entities in bulk.
type ArtistCreateBulk struct {
	config
	builders []*ArtistCreate
}

// Save creates the Artist entities in the database.
func (acb *ArtistCreateBulk) Save(ctx context.Context) ([]*Artist, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Artist, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArtistMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArtistCreateBulk) SaveX(ctx context.Context) []*Artist {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArtistCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArtistCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
