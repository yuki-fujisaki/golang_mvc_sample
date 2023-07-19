package controller

import (
	"encoding/json"
	"net/http"

	"golang_mvc_sample/pkg/view"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []view.User{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := view.User{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
