package main

import (
	"api-inventory-go/config/database"
	"api-inventory-go/testing"
	"log"
)

func main() {
	db, err := database.GetPostgresDB()
	if err != nil {
		log.Println(err)
	}

	// err = testing.DatabaseConnectionTest(db)

	// err = testing.CreateUserDbTest(db)
	// err = testing.UpdateUserDbTest(db)
	// err = testing.DeleteUserDbTest(db)
	// err = testing.GetUserDbTest(db)
	// err = testing.GetUsersDbTest(db)

	// err = testing.CreateGoodDbTest(db)
	// err = testing.UpdateGoodDbTest(db)
	// err = testing.DeleteGoodDbTest(db)
	// err = testing.GetGoodDbTest(db)
	err = testing.GetGoodsDbTest(db)

	if err != nil {
		log.Println(err)
	}
}
