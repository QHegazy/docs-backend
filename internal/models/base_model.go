package models

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// ResultChan is a generic struct to encapsulate the result of operations.
type ResultChan[T any] struct {
	Error error
	Data  T
}

// Operations defines the CRUD operations that can be performed on a model.
type Operations[T any] interface {
	Insert(pool *pgxpool.Pool, resultChan chan<- ResultChan[T])
	Update(pool *pgxpool.Pool, resultChan chan<- ResultChan[T])
	Delete(pool *pgxpool.Pool, resultChan chan<- ResultChan[T])
	Query(pool *pgxpool.Pool, resultChan chan<- ResultChan[T])
}
