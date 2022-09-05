package mysql_connector

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection(uri string, waitSecondsCaseError int, retry int) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: uri}), &gorm.Config{})
	if err != nil {
		log.Printf("ERROR: %v", err)
		if retry > 0 {
			time.Sleep(time.Duration(waitSecondsCaseError) * time.Second)
			log.Println("Trying to connect the database again...")
			retry -= 1
			return GetConnection(uri, waitSecondsCaseError, retry)
		}
		return nil, err
	}

	return db, nil
}
