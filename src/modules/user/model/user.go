package model

import (
	"time"
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
