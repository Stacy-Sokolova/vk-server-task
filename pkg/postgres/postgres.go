package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPostgresDB(ctx context.Context, cfg Config) (*Postgres, error) {
	var url = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	pg := &Postgres{Pool: db}

	err = pg.Pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to ping to db: %w", err)
	}

	return pg, nil
}

func (p *Postgres) Close() {
	p.Pool.Close()
}
