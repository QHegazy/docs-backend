package models

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Operations interface {
	Insert(pool *pgxpool.Pool) error
	Update(pool *pgxpool.Pool) error
	Delete(pool *pgxpool.Pool) error
	Query(pool *pgxpool.Pool) (User, error)
}
