package hashing

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

var hashCost = bcrypt.DefaultCost

func HashFromPassword(pass string) ([]byte, error) {
	bPass := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(bPass, hashCost)
	if err != nil {
		log.Println("fail during hash password")
	}

	return hash, err
}
