// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/Zubayear/song-ql/ent/migrate"
	"github.com/google/uuid"

	"github.com/Zubayear/song-ql/ent/artist"
	"github.com/Zubayear/song-ql/ent/song"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Artist is the client for interacting with the Artist builders.
	Artist *ArtistClient
	// Song is the client for interacting with the Song builders.
	Song *SongClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Artist = NewArtistClient(c.config)
	c.Song = NewSongClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Artist: NewArtistClient(cfg),
		Song:   NewSongClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:    ctx,
		config: cfg,
		Artist: NewArtistClient(cfg),
		Song:   NewSongClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Artist.
//		Query().
//		Count(ctx)
//
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
	c.Artist.Use(hooks...)
	c.Song.Use(hooks...)
}

// ArtistClient is a client for the Artist schema.
type ArtistClient struct {
	config
}

// NewArtistClient returns a client for the Artist from the given config.
func NewArtistClient(c config) *ArtistClient {
	return &ArtistClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `artist.Hooks(f(g(h())))`.
func (c *ArtistClient) Use(hooks ...Hook) {
	c.hooks.Artist = append(c.hooks.Artist, hooks...)
}

// Create returns a create builder for Artist.
func (c *ArtistClient) Create() *ArtistCreate {
	mutation := newArtistMutation(c.config, OpCreate)
	return &ArtistCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Artist entities.
func (c *ArtistClient) CreateBulk(builders ...*ArtistCreate) *ArtistCreateBulk {
	return &ArtistCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Artist.
func (c *ArtistClient) Update() *ArtistUpdate {
	mutation := newArtistMutation(c.config, OpUpdate)
	return &ArtistUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArtistClient) UpdateOne(a *Artist) *ArtistUpdateOne {
	mutation := newArtistMutation(c.config, OpUpdateOne, withArtist(a))
	return &ArtistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArtistClient) UpdateOneID(id uuid.UUID) *ArtistUpdateOne {
	mutation := newArtistMutation(c.config, OpUpdateOne, withArtistID(id))
	return &ArtistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Artist.
func (c *ArtistClient) Delete() *ArtistDelete {
	mutation := newArtistMutation(c.config, OpDelete)
	return &ArtistDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ArtistClient) DeleteOne(a *Artist) *ArtistDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ArtistClient) DeleteOneID(id uuid.UUID) *ArtistDeleteOne {
	builder := c.Delete().Where(artist.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArtistDeleteOne{builder}
}

// Query returns a query builder for Artist.
func (c *ArtistClient) Query() *ArtistQuery {
	return &ArtistQuery{
		config: c.config,
	}
}

// Get returns a Artist entity by its id.
func (c *ArtistClient) Get(ctx context.Context, id uuid.UUID) (*Artist, error) {
	return c.Query().Where(artist.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArtistClient) GetX(ctx context.Context, id uuid.UUID) *Artist {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySongs queries the songs edge of a Artist.
func (c *ArtistClient) QuerySongs(a *Artist) *SongQuery {
	query := &SongQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(artist.Table, artist.FieldID, id),
			sqlgraph.To(song.Table, song.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, artist.SongsTable, artist.SongsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArtistClient) Hooks() []Hook {
	return c.hooks.Artist
}

// SongClient is a client for the Song schema.
type SongClient struct {
	config
}

// NewSongClient returns a client for the Song from the given config.
func NewSongClient(c config) *SongClient {
	return &SongClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `song.Hooks(f(g(h())))`.
func (c *SongClient) Use(hooks ...Hook) {
	c.hooks.Song = append(c.hooks.Song, hooks...)
}

// Create returns a create builder for Song.
func (c *SongClient) Create() *SongCreate {
	mutation := newSongMutation(c.config, OpCreate)
	return &SongCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Song entities.
func (c *SongClient) CreateBulk(builders ...*SongCreate) *SongCreateBulk {
	return &SongCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Song.
func (c *SongClient) Update() *SongUpdate {
	mutation := newSongMutation(c.config, OpUpdate)
	return &SongUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SongClient) UpdateOne(s *Song) *SongUpdateOne {
	mutation := newSongMutation(c.config, OpUpdateOne, withSong(s))
	return &SongUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SongClient) UpdateOneID(id uuid.UUID) *SongUpdateOne {
	mutation := newSongMutation(c.config, OpUpdateOne, withSongID(id))
	return &SongUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Song.
func (c *SongClient) Delete() *SongDelete {
	mutation := newSongMutation(c.config, OpDelete)
	return &SongDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SongClient) DeleteOne(s *Song) *SongDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SongClient) DeleteOneID(id uuid.UUID) *SongDeleteOne {
	builder := c.Delete().Where(song.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SongDeleteOne{builder}
}

// Query returns a query builder for Song.
func (c *SongClient) Query() *SongQuery {
	return &SongQuery{
		config: c.config,
	}
}

// Get returns a Song entity by its id.
func (c *SongClient) Get(ctx context.Context, id uuid.UUID) (*Song, error) {
	return c.Query().Where(song.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SongClient) GetX(ctx context.Context, id uuid.UUID) *Song {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryArtists queries the artists edge of a Song.
func (c *SongClient) QueryArtists(s *Song) *ArtistQuery {
	query := &ArtistQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(song.Table, song.FieldID, id),
			sqlgraph.To(artist.Table, artist.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, song.ArtistsTable, song.ArtistsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SongClient) Hooks() []Hook {
	return c.hooks.Song
}
