package controller

import (
	"log"

	"api-inventory-go/src/modules/good/model"
	"api-inventory-go/src/modules/good/repository"
)

func CreateGood(u *model.Good, repo repository.GoodRepository) error {
	if err := repo.Create(u); err != nil {
		return err
	}

	log.Printf(`New Good with id [%s] saved!`, u.ID)
	return nil
}

func UpdateGood(u *model.Good, repo repository.GoodRepository) error {
	if err := repo.Update(u.ID, u); err != nil {
		return err
	}

	log.Printf(`Good with id [%s] updated!`, u.ID)
	return nil
}

func DeleteGood(id string, repo repository.GoodRepository) error {
	if err := repo.Delete(id); err != nil {
		return err
	}

	log.Printf(`Good with id [%s] deleted!`, id)
	return nil
}

func GetGood(id string, repo repository.GoodRepository) (*model.Good, error) {
	good, err := repo.FindById(id)
	if err != nil {
		return nil, err
	}

	log.Printf(`User with id [%s] viewed!`, id)
	return good, nil
}

func GetGoods(repo repository.GoodRepository) (model.Goods, error) {
	goods, err := repo.FindAll()
	if err != nil {
		return nil, err
	}

	log.Printf(`All Goods viewed!`)
	return goods, nil
}