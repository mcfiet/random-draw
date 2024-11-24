package db

import (
	"log"

	drawModel "github.com/mcfiet/goDo/draw/model"
	userModel "github.com/mcfiet/goDo/user/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Berlin"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&drawModel.DrawResult{})
	db.AutoMigrate(&userModel.User{})

	return db
}

func GetDB() *gorm.DB {
	return database
}
