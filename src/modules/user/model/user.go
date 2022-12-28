package model

import (
	"time"

	"api-inventory-go/helper"
)

type User struct {
	ID        string
	Fullname  string
	Nickname  string
	Username  string
	NoHP      string
	Email     string
	Password  string
	Image     string
	Office    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users []User

func NewUser() *User {
	return &User{
		ID:        helper.GenerateUserID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdatedUser() *User {
	return &User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}