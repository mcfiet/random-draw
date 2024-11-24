package router

import (
	"net/http"

	"github.com/mcfiet/goDo/auth/middleware"
	"github.com/mcfiet/goDo/draw/controller"

	"github.com/go-chi/chi/v5"
)

func DrawRouter(controller *controller.DrawController) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.AuthenticationMiddleware)
	r.Get("/", controller.GetDrawByGiverId)
	r.Post("/", controller.CreateDraw)
	r.Get("/{id}", controller.GetDrawByGiverId)

	return r
}
