// Code generated by entc, DO NOT EDIT.

package song

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Zubayear/song-ql/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v float64) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDuration), v))
	})
}

// LyricsExits applies equality check predicate on the "lyricsExits" field. It's identical to LyricsExitsEQ.
func LyricsExits(v bool) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLyricsExits), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Song {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Song(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Song {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Song(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v float64) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDuration), v))
	})
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v float64) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDuration), v))
	})
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...float64) predicate.Song {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Song(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDuration), v...))
	})
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...float64) predicate.Song {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Song(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDuration), v...))
	})
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v float64) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDuration), v))
	})
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v float64) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDuration), v))
	})
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v float64) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDuration), v))
	})
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v float64) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDuration), v))
	})
}

// LyricsExitsEQ applies the EQ predicate on the "lyricsExits" field.
func LyricsExitsEQ(v bool) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLyricsExits), v))
	})
}

// LyricsExitsNEQ applies the NEQ predicate on the "lyricsExits" field.
func LyricsExitsNEQ(v bool) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLyricsExits), v))
	})
}

// HasArtists applies the HasEdge predicate on the "artists" edge.
func HasArtists() predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArtistsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ArtistsTable, ArtistsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArtistsWith applies the HasEdge predicate on the "artists" edge with a given conditions (other predicates).
func HasArtistsWith(preds ...predicate.Artist) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArtistsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ArtistsTable, ArtistsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Song) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Song) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Song) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		p(s.Not())
	})
}