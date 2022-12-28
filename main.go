package main

import (
	"api-inventory-go/config/database"
	"api-inventory-go/testing"
	"log"
)

func main() {
	db, err := database.GetPostgresDB()

	err = testing.DatabaseConnectionTest(db)

	if err != nil {
		log.Println(err)
	}
}
