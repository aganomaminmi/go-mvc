package controller

import (
	"net/http"

	"github.com/aganomaminmi/go-mvc/pkg/model"
	"github.com/aganomaminmi/go-mvc/pkg/view"
)

func GetUser(w http.ResponseWriter, i string) {
	usr := model.User{}
	err := usr.Get(i)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("null"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(view.UserView{}.Create(usr))
}
