package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      string
	Email    string
	Password string
	Address  string
}

type APIResponseUser struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
