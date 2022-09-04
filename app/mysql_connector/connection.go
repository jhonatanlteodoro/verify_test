package mysql_connector

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection(waitSecondsCaseError int, retry int) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: "root:Password@tcp(127.0.0.1:33066)/mylocalsql?charset=utf8&parseTime=True&loc=Local"}), &gorm.Config{})
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
