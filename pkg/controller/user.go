package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/aganomaminmi/go-mvc/pkg/model"
	"github.com/aganomaminmi/go-mvc/pkg/view"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	pq, err := NewPagingQuery(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	p := model.Page{
		PageSize:    pq.PageSize,
		CurrentPage: pq.Page,
	}
	usrs, err := model.GetAllUser(&p)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	j, err := view.NewUsersViewJSON(usrs, p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetUser(w http.ResponseWriter, i string) {
	w.Header().Set("Content-Type", "application/json")
	usr := model.User{}
	if err := usr.Get(i); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	j, err := view.NewUserView(usr).ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

type UserNew struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
}

func (u UserNew) MapUser() model.User {
	usr := model.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Age:       sql.NullInt16{Int16: int16(u.Age), Valid: true},
		Sex:       sql.NullString{String: u.Sex, Valid: true},
	}

	if u.Age == 0 {
		usr.Age.Valid = false
	}
	if u.Sex == "" {
		usr.Age.Valid = false
	}

	return usr
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bdy := UserNew{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&bdy); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	usr := bdy.MapUser()

	if err := usr.Save(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	j, err := view.NewUserView(usr).ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
