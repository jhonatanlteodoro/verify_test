package models

import (
	"github.com/jhonatanlteodoro/verify_test/app/hashing"
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

func (u *User) SetHashPassword() error {
	hPass, err := hashing.HashFromPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hPass)
	return nil
}
