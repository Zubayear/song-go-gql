// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Zubayear/song-ql/ent/artist"
	"github.com/Zubayear/song-ql/ent/predicate"
	"github.com/Zubayear/song-ql/ent/song"
	"github.com/google/uuid"
)

// ArtistUpdate is the builder for updating Artist entities.
type ArtistUpdate struct {
	config
	hooks    []Hook
	mutation *ArtistMutation
}

// Where appends a list predicates to the ArtistUpdate builder.
func (au *ArtistUpdate) Where(ps ...predicate.Artist) *ArtistUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *ArtistUpdate) SetName(s string) *ArtistUpdate {
	au.mutation.SetName(s)
	return au
}

// SetAge sets the "age" field.
func (au *ArtistUpdate) SetAge(i int) *ArtistUpdate {
	au.mutation.ResetAge()
	au.mutation.SetAge(i)
	return au
}

// AddAge adds i to the "age" field.
func (au *ArtistUpdate) AddAge(i int) *ArtistUpdate {
	au.mutation.AddAge(i)
	return au
}

// SetSongsID sets the "songs" edge to the Song entity by ID.
func (au *ArtistUpdate) SetSongsID(id uuid.UUID) *ArtistUpdate {
	au.mutation.SetSongsID(id)
	return au
}

// SetNillableSongsID sets the "songs" edge to the Song entity by ID if the given value is not nil.
func (au *ArtistUpdate) SetNillableSongsID(id *uuid.UUID) *ArtistUpdate {
	if id != nil {
		au = au.SetSongsID(*id)
	}
	return au
}

// SetSongs sets the "songs" edge to the Song entity.
func (au *ArtistUpdate) SetSongs(s *Song) *ArtistUpdate {
	return au.SetSongsID(s.ID)
}

// Mutation returns the ArtistMutation object of the builder.
func (au *ArtistUpdate) Mutation() *ArtistMutation {
	return au.mutation
}

// ClearSongs clears the "songs" edge to the Song entity.
func (au *ArtistUpdate) ClearSongs() *ArtistUpdate {
	au.mutation.ClearSongs()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArtistUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArtistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArtistUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArtistUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArtistUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ArtistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   artist.Table,
			Columns: artist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: artist.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: artist.FieldName,
		})
	}
	if value, ok := au.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: artist.FieldAge,
		})
	}
	if value, ok := au.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: artist.FieldAge,
		})
	}
	if au.mutation.SongsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.SongsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ArtistUpdateOne is the builder for updating a single Artist entity.
type ArtistUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArtistMutation
}

// SetName sets the "name" field.
func (auo *ArtistUpdateOne) SetName(s string) *ArtistUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetAge sets the "age" field.
func (auo *ArtistUpdateOne) SetAge(i int) *ArtistUpdateOne {
	auo.mutation.ResetAge()
	auo.mutation.SetAge(i)
	return auo
}

// AddAge adds i to the "age" field.
func (auo *ArtistUpdateOne) AddAge(i int) *ArtistUpdateOne {
	auo.mutation.AddAge(i)
	return auo
}

// SetSongsID sets the "songs" edge to the Song entity by ID.
func (auo *ArtistUpdateOne) SetSongsID(id uuid.UUID) *ArtistUpdateOne {
	auo.mutation.SetSongsID(id)
	return auo
}

// SetNillableSongsID sets the "songs" edge to the Song entity by ID if the given value is not nil.
func (auo *ArtistUpdateOne) SetNillableSongsID(id *uuid.UUID) *ArtistUpdateOne {
	if id != nil {
		auo = auo.SetSongsID(*id)
	}
	return auo
}

// SetSongs sets the "songs" edge to the Song entity.
func (auo *ArtistUpdateOne) SetSongs(s *Song) *ArtistUpdateOne {
	return auo.SetSongsID(s.ID)
}

// Mutation returns the ArtistMutation object of the builder.
func (auo *ArtistUpdateOne) Mutation() *ArtistMutation {
	return auo.mutation
}

// ClearSongs clears the "songs" edge to the Song entity.
func (auo *ArtistUpdateOne) ClearSongs() *ArtistUpdateOne {
	auo.mutation.ClearSongs()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArtistUpdateOne) Select(field string, fields ...string) *ArtistUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Artist entity.
func (auo *ArtistUpdateOne) Save(ctx context.Context) (*Artist, error) {
	var (
		err  error
		node *Artist
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArtistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArtistUpdateOne) SaveX(ctx context.Context) *Artist {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArtistUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArtistUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ArtistUpdateOne) sqlSave(ctx context.Context) (_node *Artist, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   artist.Table,
			Columns: artist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: artist.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Artist.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, artist.FieldID)
		for _, f := range fields {
			if !artist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != artist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: artist.FieldName,
		})
	}
	if value, ok := auo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: artist.FieldAge,
		})
	}
	if value, ok := auo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: artist.FieldAge,
		})
	}
	if auo.mutation.SongsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.SongsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Artist{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
