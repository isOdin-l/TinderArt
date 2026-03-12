package postgresql

import (
	"context"

	config "github.com/isOdin-l/TinderArt/services/auth/configs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	conn *pgxpool.Pool
}

func NewPostgresDB(cfg *config.Config) (*PostgresDB, error) {
	conn, err := pgxpool.New(context.Background(), cfg.DSNPsql())
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &PostgresDB{conn}, nil
}

func (ps *PostgresDB) Exec(ctx context.Context, sql string, values ...interface{}) error {
	_, err := ps.conn.Exec(ctx, sql, values...)
	return err
}

func (ps *PostgresDB) QueryRow(ctx context.Context, sql string, values ...interface{}) pgx.Row {
	return ps.conn.QueryRow(ctx, sql, values...)
}

func (ps *PostgresDB) Scan(row pgx.Row, dest ...any) error {
	return row.Scan(dest...)
}

func (ps *PostgresDB) Close() {
	ps.conn.Close()
}
