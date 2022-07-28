package model

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aganomaminmi/go-mvc/pkg/database"
)

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
	if err := database.DB.First(u).Error; err != nil {
		return fmt.Errorf("Not found %d", err)
	}

	return nil
}

func (u *User) Save() error {
	if u.Email == "" {
		return fmt.Errorf("error: %s code=%d", "Invalid email", http.StatusBadRequest)
	}
	if err := database.DB.Where("email = ?", u.Email).First(&User{}).Error; err == nil {
		return fmt.Errorf("error: %s code=%d", "User already exist.", http.StatusInternalServerError)
	}

	if err := database.DB.Create(&u).Error; err != nil {
		return fmt.Errorf("Unknown error occurred")
	}

	return nil

}
