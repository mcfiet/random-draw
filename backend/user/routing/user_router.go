package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcfiet/goDo/user/controller"
)

func UserRouter(controller *controller.UserController) http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", controller.FindById)
	r.Get("/", controller.FindAll)
	r.Post("/", controller.Save)
	r.Put("/", controller.Update)
	return r
}
