package models

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Document struct {
	DocumentName string    `json:"document_name" validate:"required,max=255"`
	MongoID      string    `json:"mongo_id" validate:"required,len=24"`
	DocumentID   uuid.UUID `json:"document_id"`
}

func (d *Document) Insert(pool *pgxpool.Pool, resultChan chan<- ResultChan[uuid.UUID]) {
	query := `INSERT INTO public.documents (document_name, mongo_id) VALUES ($1, $2) RETURNING document_id`

	var docID uuid.UUID
	err := pool.QueryRow(context.Background(), query, d.DocumentName, d.MongoID).Scan(&docID)
	if err != nil {
		resultChan <- ResultChan[uuid.UUID]{Error: err}
		return
	}
	resultChan <- ResultChan[uuid.UUID]{Data: docID}
}

func (d *Document) Update(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `UPDATE documents SET document_name = $1 WHERE mongo_id = $2`
	_, err := pool.Exec(context.Background(), query, d.DocumentName, d.MongoID)
	resultChan <- ResultChan[error]{Error: err}
}

func (d *Document) Delete(pool *pgxpool.Pool, resultChan chan<- ResultChan[error]) {
	query := `DELETE FROM documents WHERE mongo_id = $1`
	_, err := pool.Exec(context.Background(), query, d.MongoID)
	resultChan <- ResultChan[error]{Error: err}
}

func (d *Document) Query(pool *pgxpool.Pool, resultChan chan<- ResultChan[*Document]) {
	var doc Document
	query := `SELECT document_name, mongo_id FROM public.documents WHERE document_id = $1`
	err := pool.QueryRow(context.Background(), query, d.DocumentID).Scan(&doc.DocumentName, &doc.MongoID)
	if err != nil {
		resultChan <- ResultChan[*Document]{Error: err}
		return
	}
	resultChan <- ResultChan[*Document]{Data: &doc}
}

// func (d *Document) QueryAll(pool *pgxpool.Pool, resultChan chan<- ResultChan[*[]Document]) {
// 	rows, err := pool.Query(context.Background(), `SELECT document_name, mongo_id FROM documents ORDER BY document_name ASC`)
// 	if err != nil {
// 		resultChan <- ResultChan[*[]Document]{Error: err}
// 		return
// 	}
// 	defer rows.Close()

// 	var docs []Document
// 	for rows.Next() {
// 		var doc Document
// 		err := rows.Scan(&doc.DocumentName, &doc.MongoID)
// 		if err != nil {
// 			resultChan <- ResultChan[*[]Document]{Error: err}
// 			return
// 		}
// 		docs = append(docs, doc)
// 	}

// 	if err := rows.Err(); err != nil {
// 		resultChan <- ResultChan[*[]Document]{Error: err}
// 		return
// 	}

// 	resultChan <- ResultChan[*[]Document]{Data: &docs}
// }
