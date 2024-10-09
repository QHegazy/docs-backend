package models

import (
	"context"

	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Name     string `validate:"required,min=2,max=100"`
	OauthID  string `validate:"required"`
	ImageURL string `validate:"omitempty,url"`
	Email    string `validate:"required,email"`
}

func (u *User) Insert(pool *pgxpool.Pool) (uuid.UUID, error) {
	insertSQL := "INSERT INTO users (name, oauth_id, image_url, email) VALUES ($1, $2, $3, $4) RETURNING user_id"
	var userID uuid.UUID // Change this type if your UUID is a different type

	err := pool.QueryRow(context.Background(), insertSQL, u.Name, u.OauthID, u.ImageURL, u.Email).Scan(&userID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert user: %w", err)
	}

	fmt.Printf("Inserted user with ID: %s\n", userID)
	return userID, nil
}

func (u *User) Update(pool *pgxpool.Pool, user User) error {
	updateSQL := "UPDATE users SET name = $1, oauth_id = $2, image_url = $3, email = $4 WHERE oauth_id = $5"

	result, err := pool.Exec(context.Background(), updateSQL, u.Name, u.OauthID, u.ImageURL, user.Email, user.OauthID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected := result.RowsAffected()

	fmt.Printf("Updated %d user(s) successfully.\n", rowsAffected)
	return nil
}

func (u *User) Delete(pool *pgxpool.Pool, userID string) error {
	deleteSQL := "DELETE FROM users WHERE oauth_id = $1"

	result, err := pool.Exec(context.Background(), deleteSQL, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected := result.RowsAffected()

	fmt.Printf("Deleted %d user(s) successfully.\n", rowsAffected)
	return nil
}

func (u *User) UserIdQuery(pool *pgxpool.Pool, OauthID string, ch chan<- uuid.UUID) {
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
