package models

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DocumentPublicView represents the document access for public view.
type DocumentPublicView struct {
	AccessID   uuid.UUID `json:"access_id" validate:"required,len=36"`
	DocumentID uuid.UUID `json:"document_id" validate:"required,len=36"`
}

// Insert inserts a new DocumentPublicView record into the database.
func (d *DocumentPublicView) Insert(pool *pgxpool.Pool, resultChan chan<- ResultChan[string]) {
	query := `INSERT INTO public.document_access_public_view (document_id) VALUES ($1) RETURNING access_id`
	var accessID uuid.UUID
	err := pool.QueryRow(context.Background(), query, d.DocumentID).Scan(&accessID)
	if err != nil {
		resultChan <- ResultChan[string]{Error: err}
		return
	}
	resultChan <- ResultChan[string]{Data: "success"}
}

// Update updates an existing DocumentPublicView record in the database.
func (d *DocumentPublicView) Update(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `UPDATE public.document_access_public_view SET document_id = $1 WHERE access_id = $2`
	_, err := pool.Exec(context.Background(), query, d.DocumentID, d.AccessID)
	resultChan <- ResultChan[error]{Error: err}
}

// Delete deletes a DocumentPublicView record from the database.
func (d *DocumentPublicView) Delete(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `DELETE FROM public.document_access_public_view WHERE access_id = $1`
	_, err := pool.Exec(context.Background(), query, d.AccessID)
	resultChan <- ResultChan[error]{Error: err}
}

// Query retrieves a DocumentPublicView record from the database.
func (d *DocumentPublicView) Query(pool *pgxpool.Pool, resultChan chan<- ResultChan[*DocumentPublicView]) {
	var doc DocumentPublicView
	query := `SELECT access_id, document_id FROM public.document_access_public_view WHERE access_id = $1`
	err := pool.QueryRow(context.Background(), query, d.AccessID).Scan(&doc.AccessID, &doc.DocumentID)
	if err != nil {
		resultChan <- ResultChan[*DocumentPublicView]{Error: err}
		return
	}
	resultChan <- ResultChan[*DocumentPublicView]{Data: &doc}
}
