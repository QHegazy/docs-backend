package models

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Session represents a user session entity
type Session struct {
	UserID    uuid.UUID `json:"user_id" validate:"required,uuid"`  // Associated user's UUID (foreign key to public.users)
	Token     string    `json:"token" validate:"required,max=255"` // Session token
	ExpiresAt time.Time `json:"expires_at" validate:"required"`    // Expiry time for the session
	Online    bool      `json:"online"`                            // Whether the user is currently online

}

// Validate validates the Session struct fields
func (s *Session) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}
func (s *Session) Insert(pool *pgxpool.Pool) error {
	insertSQL := "INSERT INTO auth.sessions (user_id, token, expires_at, online) VALUES ($1, $2, $3, $4)"
	result, err := pool.Exec(context.Background(), insertSQL, s.UserID, s.Token, s.ExpiresAt, s.Online)
	if err != nil {
		return fmt.Errorf("failed to insert session: %w", err)
	}

	rowsAffected := result.RowsAffected()
	fmt.Printf("Inserted %d session(s) successfully.\n", rowsAffected)
	return nil
}

func (s *Session) Update(pool *pgxpool.Pool) error {
	updateSQL := "UPDATE auth.sessions SET token = $1, expires_at = $2, online = $3 WHERE user_id = $4 AND token = $5"

	result, err := pool.Exec(context.Background(), updateSQL, s.Token, s.ExpiresAt, s.Online, s.UserID, s.Token)
	if err != nil {
		return fmt.Errorf("failed to update session: %w", err)
	}

	rowsAffected := result.RowsAffected()
	fmt.Printf("Updated %d session(s) successfully.\n", rowsAffected)
	return nil
}

func (s *Session) Delete(pool *pgxpool.Pool) error {
	deleteSQL := "DELETE FROM auth.sessions WHERE user_id = $1 AND token = $2"

	result, err := pool.Exec(context.Background(), deleteSQL, s.UserID, s.Token)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	rowsAffected := result.RowsAffected()
	fmt.Printf("Deleted %d session(s) successfully.\n", rowsAffected)
	return nil
}

func (s *Session) Query(pool *pgxpool.Pool, resultChan chan<- ResultChan[Session]) {
	querySQL := "SELECT token, expires_at FROM auth.sessions WHERE token = $1"
	row := pool.QueryRow(context.Background(), querySQL, s.Token)

	var session Session
	if err := row.Scan(&session.Token, &session.ExpiresAt); err != nil {
		resultChan <- ResultChan[Session]{Error: err}
		return
	}

	resultChan <- ResultChan[Session]{Data: session}

}
