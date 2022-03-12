// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Zubayear/song-ql/ent/song"
	"github.com/google/uuid"
)

// Song is the model entity for the Song schema.
type Song struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration float64 `json:"duration,omitempty"`
	// LyricsExits holds the value of the "lyricsExits" field.
	LyricsExits bool `json:"lyricsExits,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SongQuery when eager-loading is set.
	Edges SongEdges `json:"edges"`
}

// SongEdges holds the relations/edges for other nodes in the graph.
type SongEdges struct {
	// Artists holds the value of the artists edge.
	Artists []*Artist `json:"artists,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ArtistsOrErr returns the Artists value or an error if the edge
// was not loaded in eager-loading.
func (e SongEdges) ArtistsOrErr() ([]*Artist, error) {
	if e.loadedTypes[0] {
		return e.Artists, nil
	}
	return nil, &NotLoadedError{edge: "artists"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Song) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case song.FieldLyricsExits:
			values[i] = new(sql.NullBool)
		case song.FieldDuration:
			values[i] = new(sql.NullFloat64)
		case song.FieldTitle:
			values[i] = new(sql.NullString)
		case song.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Song", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Song fields.
func (s *Song) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case song.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case song.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case song.FieldDuration:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				s.Duration = value.Float64
			}
		case song.FieldLyricsExits:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field lyricsExits", values[i])
			} else if value.Valid {
				s.LyricsExits = value.Bool
			}
		}
	}
	return nil
}

// QueryArtists queries the "artists" edge of the Song entity.
func (s *Song) QueryArtists() *ArtistQuery {
	return (&SongClient{config: s.config}).QueryArtists(s)
}

// Update returns a builder for updating this Song.
// Note that you need to call Song.Unwrap() before calling this method if this Song
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Song) Update() *SongUpdateOne {
	return (&SongClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Song entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Song) Unwrap() *Song {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Song is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Song) String() string {
	var builder strings.Builder
	builder.WriteString("Song(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", title=")
	builder.WriteString(s.Title)
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", s.Duration))
	builder.WriteString(", lyricsExits=")
	builder.WriteString(fmt.Sprintf("%v", s.LyricsExits))
	builder.WriteByte(')')
	return builder.String()
}

// Songs is a parsable slice of Song.
type Songs []*Song

func (s Songs) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
