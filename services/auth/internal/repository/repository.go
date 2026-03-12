package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type IDatabase interface {
	Exec(ctx context.Context, sql string, args ...interface{}) error
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Scan(row pgx.Row, dest ...any) error
	Close()
}

type AuthRepository struct {
	db IDatabase
}

func NewRepository(db IDatabase) *AuthRepository {
	return &AuthRepository{db: db}
}
