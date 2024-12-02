package models

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Block represents a row in the blocklist table
type Block struct {
	BlockerID  uuid.UUID `json:"blocker_id" validate:"required"`
	BlockedID  uuid.UUID `json:"blocked_id" validate:"required"`
	DocumentID uuid.UUID `json:"document_id" validate:"required"`
}

// Insert adds a new block record to the blocklist table
func (b *Block) Insert(pool *pgxpool.Pool, resultChan chan<- ResultChan[uuid.UUID]) {
	query := `
		INSERT INTO public.blocklist (blocker_id, blocked_id, document_id, created_at, updated_at)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING blocker_id
	`

	var blockerID uuid.UUID
	err := pool.QueryRow(context.Background(), query, b.BlockerID, b.BlockedID, b.DocumentID).Scan(&blockerID)
	if err != nil {
		resultChan <- ResultChan[uuid.UUID]{Error: err}
		return
	}
	resultChan <- ResultChan[uuid.UUID]{Data: blockerID}
}

// Update modifies an existing block record's document association and timestamps
func (b *Block) Update(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `
		UPDATE public.blocklist
		SET document_id = $1, updated_at = CURRENT_TIMESTAMP
		WHERE blocker_id = $2 AND blocked_id = $3
	`

	_, err := pool.Exec(context.Background(), query, b.DocumentID, b.BlockerID, b.BlockedID)
	resultChan <- ResultChan[error]{Error: err}
}

// Delete removes a block record from the blocklist table
func (b *Block) Delete(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `DELETE FROM public.blocklist WHERE blocker_id = $1 AND blocked_id = $2`

	_, err := pool.Exec(context.Background(), query, b.BlockerID, b.BlockedID)
	resultChan <- ResultChan[error]{Error: err}
}

// Query fetches a block record by blocker_id and blocked_id
func (b *Block) Query(pool *pgxpool.Pool, resultChan chan<- ResultChan[*Block]) {
	var block Block
	query := `
		SELECT blocker_id, blocked_id, document_id, created_at, updated_at, deleted_at
		FROM public.blocklist
		WHERE blocker_id = $1 AND blocked_id = $2
	`

	err := pool.QueryRow(context.Background(), query, b.BlockerID, b.BlockedID).Scan(
		&block.BlockerID,
		&block.BlockedID,
		&block.DocumentID,
	)
	if err != nil {
		resultChan <- ResultChan[*Block]{Error: err}
		return
	}
	resultChan <- ResultChan[*Block]{Data: &block}
}
