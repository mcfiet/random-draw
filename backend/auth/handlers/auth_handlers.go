package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	authModel "github.com/mcfiet/goDo/auth/model"
	"github.com/mcfiet/goDo/user/service"
	"github.com/mcfiet/goDo/utils"
)

type AuthHandler struct {
	UserService *service.UserService
}

func NewAuthHandler(service *service.UserService) *AuthHandler {
	return &AuthHandler{service}
}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var authInput authModel.AuthInput

	if err := json.NewDecoder(r.Body).Decode(&authInput); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := handler.UserService.FindByUsername(authInput.Username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
	}

	if utils.VerifyPassword(authInput.Password, user.Password) {
		fmt.Println("Logged In")
		token, err := utils.GenerateToken(user.ID, user.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Authorization", "Bearer "+token)
	} else {
		fmt.Println("Wrong Credentials")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}
