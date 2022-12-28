package testing

import (
	"database/sql"
	"fmt"

	"api-inventory-go/src/controllers/good/controller"
	"api-inventory-go/src/modules/good/model"
	"api-inventory-go/src/modules/good/repository"
)

func CreateGoodDbTest(db *sql.DB) error {
	fmt.Println("===============================================================")
	fmt.Println("Create Good Saved to Database Test")

	newGoodObject := model.NewGood()
	// newGoodObject.Name = "Kopi Kapal Api-Special Mix"
	// newGoodObject.Name = ""
	// newGoodObject.Name = "Rokok Surya 12"
	newGoodObject.Name = "Tango Coklat"
	newGoodObject.StockMin = 3
	newGoodObject.Stock = 13

	goodRepositoryPostgres := repository.NewGoodRepositoryPostgres(db)

	err := controller.CreateGood(newGoodObject, goodRepositoryPostgres)
	if err != nil {
		return err
	}

	fmt.Println("===============================================================")

	return nil
}

func UpdateGoodDbTest(db *sql.DB) error {
	fmt.Println("===============================================================")
	fmt.Println("Update Good Saved to Database Test")

	newGoodObject := model.UpdatedGood()
	newGoodObject.ID = "good-HmQ6HtAQ8eEg1QkoJ3hdx"
	newGoodObject.Name = "Energen Coklat"
	newGoodObject.StockMin = 7
	newGoodObject.Stock = 60

	goodRepositoryPostgres := repository.NewGoodRepositoryPostgres(db)

	err := controller.UpdateGood(newGoodObject, goodRepositoryPostgres)
	if err != nil {
		return err
	}

	fmt.Println("===============================================================")

	return nil
}

func DeleteGoodDbTest(db *sql.DB) error {
	fmt.Println("===================================================================")
	fmt.Println("Delete Good in Database Test")

	id := "good-3sD1ddeInxiVtjOjyMATk"
	goodRepositoryPostgres := repository.NewGoodRepositoryPostgres(db)

	err := controller.DeleteGood(id, goodRepositoryPostgres)
	if err != nil {
		return err
	}

	fmt.Println("===================================================================")
	return nil
}

func GetGoodDbTest(db *sql.DB) error {
	fmt.Println("===================================================================")
	fmt.Println("Get Good from Database Test")

	// id := ""
	id := "good-5Xbxdxq77GcfROShcgfTy"
	goodRepositoryPostgres := repository.NewGoodRepositoryPostgres(db)

	good, err := controller.GetGood(id, goodRepositoryPostgres)
	if err != nil {
		return err
	}

	fmt.Println("Result:", good)
	fmt.Println("===================================================================")
	return nil
}

func GetGoodsDbTest(db *sql.DB) error {
	fmt.Println("===================================================================")
	fmt.Println("Get Goods from Database Test")

	goodRepositoryPostgres := repository.NewGoodRepositoryPostgres(db)

	goods, err := controller.GetGoods(goodRepositoryPostgres)
	if err != nil {
		return err
	}

	fmt.Println("Results:")
	for _, v := range goods {
		fmt.Println(v)
	}
	fmt.Println("===================================================================")
	return nil
}