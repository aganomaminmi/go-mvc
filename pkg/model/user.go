package model

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/aganomaminmi/go-mvc/pkg/database"
)

type User struct {
	ID        int            `gorm:"primary_key;column:id;autoIncrement;"`
	FirstName sql.NullString `gorm:"column:first_name;type:varchar;size:255;"`
	LastName  sql.NullString `gorm:"column:last_name;type:varchar;size:255;"`
	Email     string         `gorm:"column:email;not_null;type:varchar;size:255;"`
	Age       sql.NullInt16  `gorm:"column:age;type:int;"`
	Sex       sql.NullString `gorm:"column:sex;type:varchar;size:25;"`
}

func (u *User) Get(i string) error {
	id, err := strconv.Atoi(i)
	if err != nil {
		return fmt.Errorf("ID format error %d", err)
	}

	u.ID = id
	dbErr := database.Db.First(u).Error
	if dbErr != nil {
		return fmt.Errorf("ID format error %d", err)
	}

	return nil
}

func CreateUser() {
	nwUsr := &User{
		FirstName: sql.NullString{String: "Kohta", Valid: true},
		LastName:  sql.NullString{String: "Takanami", Valid: true},
		Email:     "hoge@fuga.com",
		Age:       sql.NullInt16{Int16: 22, Valid: true},
		Sex:       sql.NullString{String: "male", Valid: true},
	}

	database.Db.Create(&nwUsr)
}
