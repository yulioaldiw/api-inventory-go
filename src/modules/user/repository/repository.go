package repository

import "api-inventory-go/src/modules/user/model"

type UserRepository interface {
	Create(*model.User) error
	Update(string, *model.User) error
	Delete(string) error
	FindById(string) (*model.User, error)
	FindAll() (model.Users, error)
}
