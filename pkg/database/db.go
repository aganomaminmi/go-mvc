package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	DB = DBNew()
}

func DBNew() *gorm.DB {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	USER := os.Getenv("DATABASE_USER")
	PASS := os.Getenv("DATABASE_PASS")
	PROTOCOL := os.Getenv("DATABASE_PROTOCOL")
	DBNAME := os.Getenv("DATABASE_NAME")
	QUERY := "charset=utf8mb4&parseTime=True&loc=Local"

	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + QUERY
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
