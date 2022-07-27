package view

import (
	"encoding/json"

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

func (UserView) Create(u model.User) []byte {
	usrVi := UserView{
		ID:        u.ID,
		FirstName: u.FirstName.String,
		LastName:  u.LastName.String,
		Email:     u.Email,
		Age:       int(u.Age.Int16),
		Sex:       u.Sex.String,
	}
	vi, err := json.Marshal(usrVi)
	if err != nil {
		return []byte("Unknown error occurred.")
	}
	return vi
}
