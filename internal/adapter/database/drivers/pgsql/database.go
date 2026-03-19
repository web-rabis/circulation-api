package pgsql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/web-rabis/circulation-api/internal/adapter/database/drivers"
	"github.com/web-rabis/circulation-api/internal/adapter/database/drivers/pgsql/order"
)

const (
	connectionTimeout = 10 * time.Second
	ensureIdxTimeout  = 10 * time.Second
)

type PgSql struct {
	connURL string
	dbName  string

	client *pgconn.PgConn
	pool   *pgxpool.Pool
	config *pgxpool.Config

	orderRepo         drivers.OrderRepository
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
func (m *PgSql) Order() drivers.OrderRepository {
	if m.orderRepo == nil {
		m.orderRepo = order.New(m.pool)
	}

	return m.orderRepo
}
