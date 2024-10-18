package models

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

// Insert inserts a new session into the database
func (s *Session) Insert(pool *pgxpool.Pool, resultChan chan<- ResultChan[Session]) {
	insertSQL := "INSERT INTO auth.sessions (user_id, token, expires_at, online) VALUES ($1, $2, $3, $4)"
	result, err := pool.Exec(context.Background(), insertSQL, s.UserID, s.Token, s.ExpiresAt, s.Online)
	if err != nil {
		resultChan <- ResultChan[Session]{Error: fmt.Errorf("failed to insert session: %w", err)}
		return
	}

	rowsAffected := result.RowsAffected()
	fmt.Printf("Inserted %d session(s) successfully.\n", rowsAffected)
	resultChan <- ResultChan[Session]{Data: *s} // Send the inserted session data
}

// Update updates an existing session in the database
func (s *Session) Update(pool *pgxpool.Pool, resultChan chan<- ResultChan[Session]) {
	updateSQL := "UPDATE auth.sessions SET token = $1, expires_at = $2, online = $3 WHERE user_id = $4 AND token = $5"

	result, err := pool.Exec(context.Background(), updateSQL, s.Token, s.ExpiresAt, s.Online, s.UserID, s.Token)
	if err != nil {
		resultChan <- ResultChan[Session]{Error: fmt.Errorf("failed to update session: %w", err)}
		return
	}

	rowsAffected := result.RowsAffected()
	fmt.Printf("Updated %d session(s) successfully.\n", rowsAffected)
	resultChan <- ResultChan[Session]{Data: *s} // Send the updated session data
}

// Delete removes a session from the database
func (s *Session) Delete(pool *pgxpool.Pool, resultChan chan<- ResultChan[Session]) {
	deleteSQL := "DELETE FROM auth.sessions WHERE user_id = $1 AND token = $2"

	result, err := pool.Exec(context.Background(), deleteSQL, s.UserID, s.Token)
	if err != nil {
		resultChan <- ResultChan[Session]{Error: fmt.Errorf("failed to delete session: %w", err)}
		return
	}

	rowsAffected := result.RowsAffected()
	fmt.Printf("Deleted %d session(s) successfully.\n", rowsAffected)
	resultChan <- ResultChan[Session]{Data: *s} // Optionally return the deleted session data
}

// Query retrieves a session from the database
func (s *Session) Query(pool *pgxpool.Pool, resultChan chan<- ResultChan[Session]) {
	querySQL := "SELECT token, expires_at FROM auth.sessions WHERE token = $1"
	row := pool.QueryRow(context.Background(), querySQL, s.Token)

	var session Session
	if err := row.Scan(&session.Token, &session.ExpiresAt); err != nil {
		resultChan <- ResultChan[Session]{Error: err}
		return
	}

	resultChan <- ResultChan[Session]{Data: session} // Send the retrieved session data
}
func (s *Session) QueryUserId(pool *pgxpool.Pool, resultChan chan<- ResultChan[User]) {
	querySQL := `
		SELECT u.name, u.image_url, u.email
		FROM auth.sessions AS s
		LEFT JOIN public.users AS u ON s.user_id = u.user_id
		WHERE s.token = $1
	`

	row := pool.QueryRow(context.Background(), querySQL, s.Token)

	var user User
	err := row.Scan(&user.Name, &user.ImageURL, &user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("No user found for this session token.")
			resultChan <- ResultChan[User]{Error: nil}
		} else {
			fmt.Printf("Error querying session and user: %v\n", err)
			resultChan <- ResultChan[User]{Error: err}
		}
		return
	}

	resultChan <- ResultChan[User]{Data: user}
}
