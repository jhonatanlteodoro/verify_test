package models

import "gorm.io/gorm"

func RunMigrations(db *gorm.DB) {

	db.AutoMigrate(&User{})
}
