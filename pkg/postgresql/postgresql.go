package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/AleksK1NG/nats-streaming/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

const (
	maxConn           = 50
	healthCheckPeriod = 3 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConns          = 10
	lazyConnect       = false
)

// NewPgxConn pool
func NewPgxConn(cfg *config.Config) (*pgxpool.Pool, error) {
	ctx := context.Background()
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.PostgreSQL.PostgresqlHost,
		cfg.PostgreSQL.PostgresqlPort,
		cfg.PostgreSQL.PostgresqlUser,
		cfg.PostgreSQL.PostgresqlDBName,
		cfg.PostgreSQL.PostgresqlSSLMode,
		cfg.PostgreSQL.PostgresqlPassword,
	)

	poolCfg, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, err
	}

	poolCfg.MaxConns = maxConn
	poolCfg.HealthCheckPeriod = healthCheckPeriod
	poolCfg.MaxConnIdleTime = maxConnIdleTime
	poolCfg.MaxConnLifetime = maxConnLifetime
	poolCfg.MinConns = minConns
	poolCfg.LazyConnect = lazyConnect

	connPool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return nil, errors.Wrap(err, "pgx.ConnectConfig")
	}

	return connPool, nil
}
