package model

import (
	"time"

	"api-inventory-go/helper"
)

type Good struct {
	ID string
	Name string
	StockMin int8
	Stock int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Goods []Good

func NewGood() *Good {
	return &Good{
		ID: helper.GenerateGoodID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdatedGood() *Good {
	return &Good{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}