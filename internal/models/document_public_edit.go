package models

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DocumentPublicEdit represents the document access for public edit.
type DocumentPublicEdit struct {
	AccessID   uuid.UUID `json:"access_id" validate:"required,len=36"`
	DocumentID uuid.UUID `json:"document_id" validate:"required,len=36"`
}

// Insert inserts a new DocumentPublicEdit record into the database.
func (d *DocumentPublicEdit) Insert(pool *pgxpool.Pool, resultChan chan<- ResultChan[string]) {
	query := `INSERT INTO public.document_access_public_edit (document_id) VALUES ($1) RETURNING access_id`
	var accessID uuid.UUID
	err := pool.QueryRow(context.Background(), query, d.DocumentID).Scan(&accessID)
	if err != nil {
		resultChan <- ResultChan[string]{Error: err}
		return
	}
	resultChan <- ResultChan[string]{Data: "success"}
}

// Update updates an existing DocumentPublicEdit record in the database.
func (d *DocumentPublicEdit) Update(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `UPDATE public.document_access_public_edit SET document_id = $1 WHERE access_id = $2`
	_, err := pool.Exec(context.Background(), query, d.DocumentID, d.AccessID)
	resultChan <- ResultChan[error]{Error: err}
}

// Delete deletes a DocumentPublicEdit record from the database.
func (d *DocumentPublicEdit) Delete(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `DELETE FROM public.document_access_public_edit WHERE access_id = $1`
	_, err := pool.Exec(context.Background(), query, d.AccessID)
	resultChan <- ResultChan[error]{Error: err}
}

// Query retrieves a DocumentPublicEdit record from the database.
func (d *DocumentPublicEdit) Query(pool *pgxpool.Pool, resultChan chan<- ResultChan[*DocumentPublicEdit]) {
	var doc DocumentPublicEdit
	query := `SELECT access_id, document_id FROM public.document_access_public_edit WHERE access_id = $1`
	err := pool.QueryRow(context.Background(), query, d.AccessID).Scan(&doc.AccessID, &doc.DocumentID)
	if err != nil {
		resultChan <- ResultChan[*DocumentPublicEdit]{Error: err}
		return
	}
	resultChan <- ResultChan[*DocumentPublicEdit]{Data: &doc}
}
