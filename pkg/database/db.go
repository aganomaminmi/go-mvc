package database

import (

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
  Db = DbNew()
}

func DbNew() *gorm.DB {
  USER     := "user"
  PASS     := "password"
  PROTOCOL := "tcp(127.0.0.1:3306)"
  DBNAME   := "go_mvc_db"
  QUERY    := "charset=utf8mb4&parseTime=True&loc=Local"

  dsn := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?"+QUERY
  db, err := gorm.Open(
    mysql.Open(dsn),
    &gorm.Config{})

  if err != nil {
    panic(err.Error())
  }
  return db
}
