package db

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection(waitSecondsCaseError int, retry int) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		log.Printf("ERROR: %v", err)
		if retry > 0 {
			time.Sleep(time.Duration(waitSecondsCaseError) * time.Second)
			log.Println("Trying to connect the database again...")
			retry -= 1
			return GetConnection(waitSecondsCaseError, retry)
		}
		return nil, err
	}

	return db, nil
}
