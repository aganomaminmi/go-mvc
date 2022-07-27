package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
  r := chi.NewRouter()
  r.Use(middleware.RequestID)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

  NewRoutes(r)
  db := dbNew()

	r.Get("/create", func(w http.ResponseWriter, r *http.Request) {
    nwUsr := &User{
      // ID: 1,
      FirstName: sql.NullString{ String: "Kohta", Valid: true },
      LastName: sql.NullString{ String: "Takanami", Valid: true },
      Email: "hoge@fuga.com",
      Age: sql.NullInt16{ Int16: 22, Valid: true },
      Sex: sql.NullString{ String: "male", Valid: true },
    }
    usr := []User{}

    db.Create(&nwUsr)
    db.Find(&usr)
    log.Println("created")
		w.Write([]byte("Created."))
	})

  if err := http.ListenAndServe(":9000", r); err !=nil {
    log.Fatal(err)
  }
}

func dbNew() *gorm.DB {
  // DBMS     := "mysql"
  USER     := "user"
  PASS     := "password"
  PROTOCOL := "tcp(127.0.0.1:3306)"
  DBNAME   := "go_mvc_db"
  QUERY    := "charset=utf8mb4&parseTime=True&loc=Local"

  dsn := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?"+QUERY
  db,err := gorm.Open(
    mysql.Open(dsn),
      // DriverName: DBNAME,
      // DSN: CONNECT }),
    &gorm.Config{})

  if err != nil {
    panic(err.Error())
  }
  return db
  // defer db.Close()
}

type User struct {
  ID        int            `gorm:"primary_key;column:id;autoIncrement;"`
  FirstName sql.NullString `gorm:"column:first_name;type:varchar;size:255;"`
  LastName  sql.NullString `gorm:"column:last_name;type:varchar;size:255;"`
  Email     string         `gorm:"column:email;not_null;type:varchar;size:255;"`
  Age       sql.NullInt16  `gorm:"column:age;type:int;"`
  Sex       sql.NullString `gorm:"column:sex;type:varchar;size:25;"`
}

