package model

import (
	"database/sql"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pressly/goose"
)


func Migrate(db *sql.DB) (error){
	err := goose.Up(db, "./internal/migrations")
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
		return err
	} else {
		log.Println("Migrations applied successfully!")
	}
	return nil
}


