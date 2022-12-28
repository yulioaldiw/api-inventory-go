package controller

import (
	"log"

	"api-inventory-go/src/modules/user/model"
	"api-inventory-go/src/modules/user/repository"
)

func CreateUser(u *model.User, repo repository.UserRepository) error {
	if err := repo.Create(u); err != nil {
		return err
	}

	log.Printf(`New User with id [%s] saved!`, u.ID)
	return nil
}

func UpdateUser(u *model.User, repo repository.UserRepository) error {
	if err := repo.Update(u.ID, u); err != nil {
		return err
	}

	log.Printf(`User with id [%s] updated!`, u.ID)
	return nil
}

func DeleteUser(id string, repo repository.UserRepository) error {
	if err := repo.Delete(id); err != nil {
		return err
	}

	log.Printf(`User with id [%s] deleted!`, id)
	return nil
}

func GetUser(id string, repo repository.UserRepository) (*model.User, error) {
	user, err := repo.FindById(id)
	if err != nil {
		return nil, err
	}

	log.Printf(`User with id [%s] viewed!`, id)
	return user, nil
}

func GetUsers(repo repository.UserRepository) (model.Users, error) {
	users, err := repo.FindAll()
	if err != nil {
		return nil, err
	}

	log.Printf(`All Users viewed!`)
	return users, nil
}
