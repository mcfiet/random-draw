package db

import (
	"fmt"
	"log"
	"os"

	drawModel "github.com/mcfiet/goDo/draw/model"
	userModel "github.com/mcfiet/goDo/user/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func Init() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatal("DB connection string is not set")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin", host, user, password, dbname, port)
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
