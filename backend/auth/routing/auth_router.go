package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcfiet/goDo/auth/handlers"
	"github.com/mcfiet/goDo/user/controller"
)

func AuthRouter(handlers *handlers.AuthHandler, controller *controller.UserController) http.Handler {
	r := chi.NewRouter()
	r.Post("/login", handlers.Login)
	r.Post("/register", controller.Save)
	return r
}
