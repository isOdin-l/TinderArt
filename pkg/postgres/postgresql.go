package postgresql

import (
	"context"

	"github.com/isOdin-l/TinderArt/pkg/configs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	TxContext = "txPOSTGIS"
)

type IExecutor interface {
	Exec(ctx context.Context, sql string, values ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, values ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, values ...any) pgx.Row
}

type PostgresDB struct {
	conn *pgxpool.Pool
}

func NewPostgresDB(cfg *configs.ConfigPostgreWithPostGIS) (*PostgresDB, error) {
	conn, err := pgxpool.New(context.Background(), cfg.DSN())
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &PostgresDB{conn}, nil
}

func (ps *PostgresDB) Exec(ctx context.Context, sql string, values ...any) (pgconn.CommandTag, error) {
	return ps.getExecutor(ctx).Exec(ctx, sql, values...)
}

func (ps *PostgresDB) Query(ctx context.Context, sql string, values ...any) (pgx.Rows, error) {
	return ps.getExecutor(ctx).Query(ctx, sql, values...)
}

func (ps *PostgresDB) QueryRow(ctx context.Context, sql string, values ...any) pgx.Row {
	return ps.getExecutor(ctx).QueryRow(ctx, sql, values...)
}

func (ps *PostgresDB) Scan(row pgx.Row, dest ...any) error {
	return row.Scan(dest...)
}

func (ps *PostgresDB) Close() {
	ps.conn.Close()
}

func (ps *PostgresDB) getExecutor(ctx context.Context) IExecutor {
	tx, ok := ctx.Value(TxContext).(pgx.Tx)
	if !ok {
		return ps.conn
	}
	return tx
}

func (ps *PostgresDB) WithTx(ctx context.Context, fn func(ctx context.Context) (any, error)) (any, error) {
	if _, ok := ctx.Value(TxContext).(pgx.Tx); ok {
		return fn(ctx)
	}

	tx, errTx := ps.conn.BeginTx(ctx, pgx.TxOptions{})
	if errTx != nil {
		return nil, errTx
	}
	defer tx.Rollback(ctx)

	ctx = context.WithValue(ctx, TxContext, tx)
	res, errFn := fn(ctx)
	if errFn != nil {
		return nil, errFn
	}

	return res, tx.Commit(ctx)
}
