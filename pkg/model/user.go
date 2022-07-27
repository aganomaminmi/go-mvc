package model

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aganomaminmi/go-mvc/pkg/database"
)

type UserNew struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
}

type User struct {
	ID        int            `gorm:"primary_key;column:id;autoIncrement;"`
	FirstName string         `gorm:"column:first_name;not_null;type:varchar;size:255;"`
	LastName  string         `gorm:"column:last_name;not_null;type:varchar;size:255;"`
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
		return fmt.Errorf("Not found %d", err)
	}

	return nil
}

func (u UserNew) Save() error {
	if u.Email == "" {
		return fmt.Errorf("error: %s code=%d", "Invalid email", http.StatusBadRequest)
	}
	err := database.Db.Where("email = ?", u.Email).First(&User{}).Error
	if err == nil {
		return fmt.Errorf("error: %s code=%d", "User already exist.", http.StatusInternalServerError)
	}

	nwUsr := User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Age:       sql.NullInt16{Int16: int16(u.Age), Valid: true},
		Sex:       sql.NullString{String: u.Sex, Valid: true},
	}

	if u.Age == 0 {
		nwUsr.Age.Valid = false
	}
	if u.Sex == "" {
		nwUsr.Age.Valid = false
	}

	crtErr := database.Db.Create(&nwUsr).Error
	if crtErr != nil {
		return fmt.Errorf("Unknown error occurred")
	}

	return nil

}
