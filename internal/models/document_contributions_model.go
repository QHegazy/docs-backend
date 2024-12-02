package models

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DocumentContribution struct {
	UserID     uuid.UUID `json:"user_id" validate:"required,len=36"`
	DocumentID uuid.UUID `json:"document_id" validate:"required,len=36"`
	Role       string    `json:"role" validate:"required,oneof=viewer editor contributor"`
}

// Insert a new document contribution record into the table.
func (d *DocumentContribution) Insert(pool *pgxpool.Pool, resultChan chan<- ResultChan[string]) {
	query := `INSERT INTO public.document_contributions (user_id, document_id, role) VALUES ($1, $2, $3)`

	result, err := pool.Exec(context.Background(), query, d.UserID, d.DocumentID, d.Role)
	if err != nil {
		resultChan <- ResultChan[string]{Error: err}
		return
	}
	fmt.Println(result.RowsAffected())
	resultChan <- ResultChan[string]{Data: "success"}
}

func (d *DocumentContribution) Update(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `UPDATE public.document_contributions SET role = $1 WHERE user_id = $2 AND document_id = $3`
	_, err := pool.Exec(context.Background(), query, d.Role, d.UserID, d.DocumentID)
	resultChan <- ResultChan[error]{Error: err}
}

func (d *DocumentContribution) Delete(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `UPDATE public.document_contributions SET deleted_at = CURRENT_TIMESTAMP WHERE user_id = $1 AND document_id = $2`
	_, err := pool.Exec(context.Background(), query, d.UserID, d.DocumentID)
	resultChan <- ResultChan[error]{Error: err}
}

func (d *DocumentContribution) Query(pool *pgxpool.Pool, resultChan chan<- ResultChan[*DocumentContribution]) {
	var doc DocumentContribution
	query := `SELECT user_id, document_id, role FROM public.document_contributions WHERE user_id = $1 AND document_id = $2`
	err := pool.QueryRow(context.Background(), query, d.UserID, d.DocumentID).Scan(&doc.UserID, &doc.DocumentID, &doc.Role)
	if err != nil {
		resultChan <- ResultChan[*DocumentContribution]{Error: err}
		return
	}
	resultChan <- ResultChan[*DocumentContribution]{Data: &doc}
}
