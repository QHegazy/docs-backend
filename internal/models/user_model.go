package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"fmt"
)

type User struct {
	Name     string
	OauthID  string
	ImageURL string
	Email    string
}

type Operations interface {
	InsertUser(pool *pgxpool.Pool) error
	UpdateUser(pool *pgxpool.Pool) error
	DeleteUser(pool *pgxpool.Pool) error
	QueryUser(pool *pgxpool.Pool) (User, error)
}

func (u *User) InsertUser(pool *pgxpool.Pool) error {
	insertSQL := "INSERT INTO users (name, oauth_id, image_url, email) VALUES ($1, $2, $3, $4)"
	result, err := pool.Exec(context.Background(), insertSQL, u.Name, u.OauthID, u.ImageURL, u.Email)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	rowsAffected := result.RowsAffected()
	fmt.Printf("Inserted %d user(s) successfully.\n", rowsAffected)
	return nil
}

func (u *User) UpdateUser(pool *pgxpool.Pool, user User) error {
	updateSQL := "UPDATE users SET name = $1, oauth_id = $2, image_url = $3, email = $4 WHERE oauth_id = $5"

	result, err := pool.Exec(context.Background(), updateSQL, u.Name, u.OauthID, u.ImageURL, user.Email, user.OauthID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected := result.RowsAffected()

	fmt.Printf("Updated %d user(s) successfully.\n", rowsAffected)
	return nil
}

func (u *User) DeleteUser(pool *pgxpool.Pool, userID string) error {
	deleteSQL := "DELETE FROM users WHERE oauth_id = $1"

	result, err := pool.Exec(context.Background(), deleteSQL, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected := result.RowsAffected()

	fmt.Printf("Deleted %d user(s) successfully.\n", rowsAffected)
	return nil
}
