package pgsql

import (
	"context"

	"github.com/web-rabis/db/internal/adapter/database/ebook/drivers"
	"github.com/web-rabis/db/internal/adapter/database/ebook/drivers/pgsql/ebook"
	"github.com/web-rabis/db/internal/adapter/database/ebook/drivers/pgsql/ebook_inv"

	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	connectionTimeout = 100 * time.Second
	ensureIdxTimeout  = 10 * time.Second
)

type PgSql struct {
	connURL string
	dbName  string

	client *pgconn.PgConn
	pool   *pgxpool.Pool
	config *pgxpool.Config

	ebook             drivers.Ebook
	ebookInv          drivers.EbookInv
	connectionTimeout time.Duration
	ensureIdxTimeout  time.Duration
}

func New(conf drivers.DataStoreConfig) (drivers.DataStore, error) {
	if conf.URL == "" {
		return nil, drivers.ErrInvalidConfigStruct
	}

	if conf.DataBaseName == "" {
		return nil, drivers.ErrInvalidConfigStruct
	}

	config, err := pgxpool.ParseConfig(conf.URL)
	if err != nil {
		return nil, err
	}
	config.MaxConns = 100
	return &PgSql{
		connURL:           conf.URL,
		dbName:            conf.DataBaseName,
		config:            config,
		connectionTimeout: connectionTimeout,
		ensureIdxTimeout:  ensureIdxTimeout,
	}, nil
}

func (m *PgSql) Name() string { return "PgSql" }

func (m *PgSql) Connect(ctx context.Context) error {
	ctxWT, cancel := context.WithTimeout(ctx, m.connectionTimeout)
	defer cancel()

	var err error
	m.pool, err = pgxpool.NewWithConfig(ctxWT, m.config)
	if err != nil {
		return err
	}
	return nil
}

func (m *PgSql) Ping(ctx context.Context) error {
	conn, err := m.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	return conn.Ping(ctx)
}

func (m *PgSql) Close(ctx context.Context) error {
	m.pool.Close()
	return nil
}

func (m *PgSql) Ebook() drivers.Ebook {
	if m.ebook == nil {
		m.ebook = ebook.New(m.pool)
	}

	return m.ebook
}

func (m *PgSql) EbookInv() drivers.EbookInv {
	if m.ebookInv == nil {
		m.ebookInv = ebook_inv.New(m.pool)
	}

	return m.ebookInv
}
