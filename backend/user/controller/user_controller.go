package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mcfiet/goDo/user/model"
	"github.com/mcfiet/goDo/user/service"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service}
}

func (controller *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Fehler beim Parsen der ID", http.StatusBadRequest)
	}

	user, err := controller.UserService.FindById(id)
	if err != nil {
		http.Error(w, "Fehler beim Holen der Daten", http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Fehler beim Codieren der User", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

func (controller *UserController) FindByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := controller.UserService.FindByUsername(username)
	if err != nil {
		http.Error(w, "Fehler beim Holen der Daten", http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Fehler beim Codieren der User", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (controller *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := controller.UserService.FindAll()
	if err != nil {
		http.Error(w, "Fehler beim Holen der Daten", http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Fehler beim Codieren der User", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (controller *UserController) Save(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	if err := controller.UserService.CheckIfUserExists(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := controller.UserService.Save(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (controller *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	err := controller.UserService.Update(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
}
