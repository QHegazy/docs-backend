package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/net/context"
)

type Connect struct {
	DbRead   *pgxpool.Pool
	DbInsert *pgxpool.Pool
	DbUpdate *pgxpool.Pool
	DbDelete *pgxpool.Pool
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *Connect
)

func New() *Connect {
	if dbInstance != nil {
		return dbInstance
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)

	dbInsert, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database for inserts: %v\n", err)
	}

	dbUpdate, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database for updates: %v\n", err)
	}

	dbDelete, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database for deletes: %v\n", err)
	}

	dbRead, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database for reads: %v\n", err)
	}

	dbInstance = &Connect{
		DbRead:   dbRead,
		DbInsert: dbInsert,
		DbUpdate: dbUpdate,
		DbDelete: dbDelete,
	}

	return dbInstance
}

func (s *Connect) Close() error {
	log.Printf("Disconnected from database: %s", database)
	s.DbDelete.Close()
	s.DbInsert.Close()
	s.DbRead.Close()
	s.DbUpdate.Close()
	return nil
}
