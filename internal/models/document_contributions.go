package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DocumentContribution struct {
	UserID     string `json:"user_id" validate:"required,len=36"`
	DocumentID string `json:"document_id" validate:"required,len=36"`
}

func (d *DocumentContribution) Insert(pool *pgxpool.Pool, resultChan chan<- ResultChan[string]) {
	query := `INSERT INTO public.document_contributions (user_id, document_id) VALUES ($1, $2) RETURNING document_id`

	var docID string
	err := pool.QueryRow(context.Background(), query, d.UserID, d.DocumentID).Scan(&docID)
	if err != nil {
		resultChan <- ResultChan[string]{Error: err}
		return
	}
	resultChan <- ResultChan[string]{Data: "success"}
}

func (d *DocumentContribution) Update(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `UPDATE public.document_contributions SET document_id = $1 WHERE user_id = $2`
	_, err := pool.Exec(context.Background(), query, d.DocumentID, d.UserID)
	resultChan <- ResultChan[error]{Error: err}
}
func (d *DocumentContribution) Delete(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `DELETE FROM public.document_contributions WHERE user_id = $1`
	_, err := pool.Exec(context.Background(), query, d.UserID)
	resultChan <- ResultChan[error]{Error: err}
}

func (d *DocumentContribution) Query(pool *pgxpool.Pool, resultChan chan<- ResultChan[*DocumentContribution]) {
	var doc DocumentContribution
	query := `SELECT user_id, document_id FROM public.document_contributions WHERE user_id = $1`
	err := pool.QueryRow(context.Background(), query, d.UserID).Scan(&doc.UserID, &doc.DocumentID)
	if err != nil {
		resultChan <- ResultChan[*DocumentContribution]{Error: err}
		return
	}
	resultChan <- ResultChan[*DocumentContribution]{Data: &doc}
}
