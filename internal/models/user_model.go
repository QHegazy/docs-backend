package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"fmt"
)

type User struct {
	Name     string `validate:"required,min=2,max=100"`
	OauthID  string `validate:"required"`
	ImageURL string `validate:"omitempty,url"`
	Email    string `validate:"required,email"`
}

func (u *User) Insert(pool *pgxpool.Pool) error {
	insertSQL := "INSERT INTO users (name, oauth_id, image_url, email) VALUES ($1, $2, $3, $4)"
	result, err := pool.Exec(context.Background(), insertSQL, u.Name, u.OauthID, u.ImageURL, u.Email)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	rowsAffected := result.RowsAffected()
	fmt.Printf("Inserted %d user(s) successfully.\n", rowsAffected)
	return nil
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
