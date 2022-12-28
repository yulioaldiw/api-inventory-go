package testing

import (
	"database/sql"
	"fmt"
)

func DatabaseConnectionTest(db *sql.DB) error {
	fmt.Println("===============================================================")
	fmt.Println("Database Connection Test")

	if err := db.Ping(); err != nil {
		return err
	}

	fmt.Println("Database connected!")
	fmt.Println("===============================================================")

	return nil
}
