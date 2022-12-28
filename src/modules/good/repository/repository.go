package repository

import "api-inventory-go/src/modules/good/model"

type GoodRepository interface {
	Create(*model.Good) error
	Update(string, *model.Good) error
	Delete(string) error
	FindById(string) (*model.Good, error)
	FindAll() (model.Goods, error)
}