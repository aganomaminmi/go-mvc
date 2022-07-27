package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	Db = DbNew()
}

func DbNew() *gorm.DB {
	envErr := godotenv.Load("../.env")
	if envErr != nil {
		log.Fatal(envErr)
		os.Exit(1)
	}

	USER := os.Getenv("DATABASE_USER")
	PASS := os.Getenv("DATABASE_PASS")
	PROTOCOL := "tcp(127.0.0.1:3306)"
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
