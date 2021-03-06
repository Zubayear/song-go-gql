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

// SongCreate is the builder for creating a Song entity.
type SongCreate struct {
	config
	mutation *SongMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (sc *SongCreate) SetTitle(s string) *SongCreate {
	sc.mutation.SetTitle(s)
	return sc
}

// SetDuration sets the "duration" field.
func (sc *SongCreate) SetDuration(f float64) *SongCreate {
	sc.mutation.SetDuration(f)
	return sc
}

// SetLyricsExits sets the "lyricsExits" field.
func (sc *SongCreate) SetLyricsExits(b bool) *SongCreate {
	sc.mutation.SetLyricsExits(b)
	return sc
}

// SetNillableLyricsExits sets the "lyricsExits" field if the given value is not nil.
func (sc *SongCreate) SetNillableLyricsExits(b *bool) *SongCreate {
	if b != nil {
		sc.SetLyricsExits(*b)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SongCreate) SetID(u uuid.UUID) *SongCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SongCreate) SetNillableID(u *uuid.UUID) *SongCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (sc *SongCreate) AddArtistIDs(ids ...uuid.UUID) *SongCreate {
	sc.mutation.AddArtistIDs(ids...)
	return sc
}

// AddArtists adds the "artists" edges to the Artist entity.
func (sc *SongCreate) AddArtists(a ...*Artist) *SongCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddArtistIDs(ids...)
}

// Mutation returns the SongMutation object of the builder.
func (sc *SongCreate) Mutation() *SongMutation {
	return sc.mutation
}

// Save creates the Song in the database.
func (sc *SongCreate) Save(ctx context.Context) (*Song, error) {
	var (
		err  error
		node *Song
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SongMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SongCreate) SaveX(ctx context.Context) *Song {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SongCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SongCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SongCreate) defaults() {
	if _, ok := sc.mutation.LyricsExits(); !ok {
		v := song.DefaultLyricsExits
		sc.mutation.SetLyricsExits(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := song.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SongCreate) check() error {
	if _, ok := sc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Song.title"`)}
	}
	if _, ok := sc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`ent: missing required field "Song.duration"`)}
	}
	if _, ok := sc.mutation.LyricsExits(); !ok {
		return &ValidationError{Name: "lyricsExits", err: errors.New(`ent: missing required field "Song.lyricsExits"`)}
	}
	return nil
}

func (sc *SongCreate) sqlSave(ctx context.Context) (*Song, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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

func (sc *SongCreate) createSpec() (*Song, *sqlgraph.CreateSpec) {
	var (
		_node = &Song{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: song.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: song.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := sc.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: song.FieldDuration,
		})
		_node.Duration = value
	}
	if value, ok := sc.mutation.LyricsExits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: song.FieldLyricsExits,
		})
		_node.LyricsExits = value
	}
	if nodes := sc.mutation.ArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   song.ArtistsTable,
			Columns: []string{song.ArtistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SongCreateBulk is the builder for creating many Song entities in bulk.
type SongCreateBulk struct {
	config
	builders []*SongCreate
}

// Save creates the Song entities in the database.
func (scb *SongCreateBulk) Save(ctx context.Context) ([]*Song, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Song, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SongMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SongCreateBulk) SaveX(ctx context.Context) []*Song {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SongCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SongCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
