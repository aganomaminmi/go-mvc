package view

import (
	"encoding/json"
	"fmt"

	"github.com/aganomaminmi/go-mvc/pkg/model"
)

type UserView struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
}

func (UserView) Create(u model.User) ([]byte, error) {
	usrVw := UserView{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Age:       int(u.Age.Int16),
		Sex:       u.Sex.String,
	}
	vw, err := json.Marshal(usrVw)
	if err != nil {
		return vw, fmt.Errorf("Unknown error occurred")
	}
	return vw, nil
}
