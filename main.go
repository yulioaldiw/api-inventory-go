package main

import (
	"api-inventory-go/config/database"
	"api-inventory-go/testing"
	"log"
)

func main() {
	db, err := database.GetPostgresDB()

	err = testing.DatabaseConnectionTest(db)
	// err = testing.CreateUserDbTest(db)
	// err = testing.UpdateUserDbTest(db)
	err = testing.DeleteUserDbTest(db)
	// err = testing.GetUserDbTest(db)
	// err = testing.GetUsersDbTest(db)

	if err != nil {
		log.Println(err)
	}
}
