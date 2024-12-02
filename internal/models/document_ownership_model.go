package models

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DocumentOwnership struct {
	UserID     uuid.UUID `json:"user_id" validate:"required,len=36"`
	DocumentID uuid.UUID `json:"document_id" validate:"required,len=36"`
}

func (d *DocumentOwnership) Insert(pool *pgxpool.Pool, resultChan chan<- ResultChan[string]) {
	query := `INSERT INTO public.document_ownerships (user_id, document_id) VALUES ($1, $2) RETURNING document_id`
	var docID string
	err := pool.QueryRow(context.Background(), query, d.UserID, d.DocumentID).Scan(&docID)
	if err != nil {
		resultChan <- ResultChan[string]{Error: err}
		return
	}
	resultChan <- ResultChan[string]{Data: "success"}
}

func (d *DocumentOwnership) Update(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `UPDATE public.document_ownerships SET user_id = $1 WHERE document_id = $2`
	_, err := pool.Exec(context.Background(), query, d.UserID, d.DocumentID)
	resultChan <- ResultChan[error]{Error: err}
}

func (d *DocumentOwnership) Delete(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `DELETE FROM public.document_ownerships WHERE document_id = $1 AND user_id = $2`
	_, err := pool.Exec(context.Background(), query, d.DocumentID, d.UserID)
	resultChan <- ResultChan[error]{Error: err}
}

func (d *DocumentOwnership) Query(pool *pgxpool.Pool, resultChan chan<- ResultChan[*DocumentOwnership]) {
	var doc DocumentOwnership
	query := `SELECT user_id, document_id FROM public.document_ownerships WHERE document_id = $1 AND user_id = $2`
	err := pool.QueryRow(context.Background(), query, d.DocumentID, d.UserID).Scan(&doc.UserID, &doc.DocumentID)
	if err != nil {
		resultChan <- ResultChan[*DocumentOwnership]{Error: err}
		return
	}
	resultChan <- ResultChan[*DocumentOwnership]{Data: &doc}
}
func (d *DocumentOwnership) QueryAllByUser(ctx context.Context, pool *pgxpool.Pool, resultChan chan<- ResultChan[[]uuid.UUID]) {
	defer close(resultChan) // Ensure the channel is closed when the function completes

	rows, err := pool.Query(ctx, `SELECT document_id FROM public.document_ownerships WHERE user_id = $1 ORDER BY updated_at DESC`, d.UserID)
	if err != nil {
		resultChan <- ResultChan[[]uuid.UUID]{Error: err}
		return
	}
	defer rows.Close()

	var docs []uuid.UUID
	for rows.Next() {
		var doc uuid.UUID
		if err := rows.Scan(&doc); err != nil {
			resultChan <- ResultChan[[]uuid.UUID]{Error: err}
			return
		}
		docs = append(docs, doc)
	}

	// Check for row iteration errors
	if err := rows.Err(); err != nil {
		resultChan <- ResultChan[[]uuid.UUID]{Error: err}
		return
	}

	// Send the collected document IDs through the channel
	resultChan <- ResultChan[[]uuid.UUID]{Data: docs}
}
