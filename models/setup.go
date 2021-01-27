package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DataBase *gorm.DB

func SetupDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(&User{}, &Profile{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatalln(err)
	}
	DataBase = db
}
