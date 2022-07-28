package view

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aganomaminmi/go-mvc/pkg/model"
)

type UserView struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Sex       string    `json:"sex"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUserView(u model.User) UserView {
	return UserView{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Age:       int(u.Age.Int16),
		Sex:       u.Sex.String,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u UserView) ToJSON() ([]byte, error) {
	JSON, err := json.Marshal(u)
	if err != nil {
		return JSON, fmt.Errorf("Unknown error occurred")
	}
	return JSON, nil
}
