package controller

import (
	"encoding/json"
	"net/http"

	"golang_mvc_sample/pkg/model"
	"golang_mvc_sample/pkg/view"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewUsers := make([]view.User, len(users))
	for i, user := range users {
		viewUsers[i] = view.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(viewUsers)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var viewUser view.User
	err := json.NewDecoder(r.Body).Decode(&viewUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := model.User{
		Name:  viewUser.Name,
		Email: viewUser.Email,
	}
	err = model.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(viewUser)
}
