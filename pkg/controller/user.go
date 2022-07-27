package controller

import (
	"encoding/json"
	"net/http"

	"github.com/aganomaminmi/go-mvc/pkg/model"
	"github.com/aganomaminmi/go-mvc/pkg/view"
)

func GetUser(w http.ResponseWriter, i string) {
	w.Header().Set("Content-Type", "application/json")
	usr := model.User{}
	err := usr.Get(i)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	usrVw, vwErr := view.UserView{}.Create(usr)
	if vwErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(vwErr.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(usrVw)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bdy := model.UserNew{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&bdy)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	svErr := bdy.Save()

	if svErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(svErr.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
