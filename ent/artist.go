// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Zubayear/song-ql/ent/artist"
	"github.com/google/uuid"
)

// Artist is the model entity for the Artist schema.
type Artist struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArtistQuery when eager-loading is set.
	Edges ArtistEdges `json:"edges"`
}

// ArtistEdges holds the relations/edges for other nodes in the graph.
type ArtistEdges struct {
	// Songs holds the value of the songs edge.
	Songs []*Song `json:"songs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SongsOrErr returns the Songs value or an error if the edge
// was not loaded in eager-loading.
func (e ArtistEdges) SongsOrErr() ([]*Song, error) {
	if e.loadedTypes[0] {
		return e.Songs, nil
	}
	return nil, &NotLoadedError{edge: "songs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Artist) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case artist.FieldAge:
			values[i] = new(sql.NullInt64)
		case artist.FieldName:
			values[i] = new(sql.NullString)
		case artist.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Artist", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Artist fields.
func (a *Artist) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case artist.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case artist.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case artist.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				a.Age = int(value.Int64)
			}
		}
	}
	return nil
}

// QuerySongs queries the "songs" edge of the Artist entity.
func (a *Artist) QuerySongs() *SongQuery {
	return (&ArtistClient{config: a.config}).QuerySongs(a)
}

// Update returns a builder for updating this Artist.
// Note that you need to call Artist.Unwrap() before calling this method if this Artist
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Artist) Update() *ArtistUpdateOne {
	return (&ArtistClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Artist entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Artist) Unwrap() *Artist {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Artist is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Artist) String() string {
	var builder strings.Builder
	builder.WriteString("Artist(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", a.Age))
	builder.WriteByte(')')
	return builder.String()
}

// Artists is a parsable slice of Artist.
type Artists []*Artist

func (a Artists) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
