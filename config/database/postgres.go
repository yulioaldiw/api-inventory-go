package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetPostgresDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB_NAME")

	desc := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		user,
		password,
		dbName,
	)

	db, err := createConnection(desc)
	if err != nil {
		return nil, err
	}

	log.Printf(`Database named [%s] with host: [%s] and user: [%s] has been connected!`, dbName, host, user)
	return db, nil

}

func createConnection(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
