package models

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// User represents a user in the system.
type User struct {
	UserID   uuid.UUID `json:"user_id"`
	Name     string    `validate:"required,min=2,max=100"`
	OauthID  string    `validate:"required"`
	ImageURL string    `validate:"omitempty,url"`
	Email    string    `validate:"required,email"`
}

// Insert inserts a new user into the database and sends the result to the channel.
func (u *User) Insert(pool *pgxpool.Pool, resultChan chan<- ResultChan[uuid.UUID]) {
	insertSQL := "INSERT INTO users (name, oauth_id, image_url, email) VALUES ($1, $2, $3, $4) RETURNING user_id"
	var userID uuid.UUID

	err := pool.QueryRow(context.Background(), insertSQL, u.Name, u.OauthID, u.ImageURL, u.Email).Scan(&userID)
	if err != nil {
		resultChan <- ResultChan[uuid.UUID]{Error: fmt.Errorf("failed to insert user: %w", err)}
		return
	}
	resultChan <- ResultChan[uuid.UUID]{Data: userID}
}

// Update updates an existing user in the database and sends the result to the channel.
func (u *User) Update(pool *pgxpool.Pool, resultChan chan<- ResultChan[int64]) {
	updateSQL := "UPDATE users SET name = $1, oauth_id = $2, image_url = $3, email = $4 WHERE oauth_id = $5"

	result, err := pool.Exec(context.Background(), updateSQL, u.Name, u.OauthID, u.ImageURL, u.Email, u.OauthID)
	if err != nil {
		resultChan <- ResultChan[int64]{Error: fmt.Errorf("failed to update user: %w", err)}
		return
	}

	rowsAffected := result.RowsAffected()
	fmt.Printf("Updated %d user(s) successfully.\n", rowsAffected)
	resultChan <- ResultChan[int64]{Data: rowsAffected}
}

// Delete deletes a user from the database by OAuth ID and sends the result to the channel.
func (u *User) Delete(pool *pgxpool.Pool, resultChan chan<- ResultChan[int64]) {
	deleteSQL := "DELETE FROM users WHERE oauth_id = $1"

	result, err := pool.Exec(context.Background(), deleteSQL, u.OauthID)
	if err != nil {
		resultChan <- ResultChan[int64]{Error: fmt.Errorf("failed to delete user: %w", err)}
		return
	}

	rowsAffected := result.RowsAffected()
	fmt.Printf("Deleted %d user(s) successfully.\n", rowsAffected)
	resultChan <- ResultChan[int64]{Data: rowsAffected}
}

// UserOauthIdQuery queries for a user's UUID by their OAuth ID and sends the result to the channel.
func (u *User) UserOauthIdQuery(pool *pgxpool.Pool, OauthID string, ch chan<- uuid.UUID) {
	querySQL := "SELECT user_id FROM users WHERE oauth_id = $1"
	var userID uuid.UUID

	err := pool.QueryRow(context.Background(), querySQL, OauthID).Scan(&userID)
	if err != nil {
		// If no rows are found, you can handle the error or return an indication via the channel
		if err == pgx.ErrNoRows {
			ch <- uuid.Nil
		} else {
			fmt.Println("Error querying user ID:", err)
			close(ch) // Close the channel on error
			return
		}
	}

	ch <- userID
}

func (u *User) QueryById(pool *pgxpool.Pool, resultChan chan<- ResultChan[User]) {
	querySQL := "SELECT name, oauth_id, image_url, email FROM users WHERE user_id = $1"
	row := pool.QueryRow(context.Background(), querySQL, u.UserID)

	var user User
	if err := row.Scan(&user.Name, &user.OauthID, &user.ImageURL, &user.Email); err != nil {
		fmt.Println("Error querying user:", err)
		return
	}
	fmt.Println(user)
}
