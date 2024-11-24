package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/mcfiet/goDo/app"
	authRouter "github.com/mcfiet/goDo/auth/routing"
	todoRouter "github.com/mcfiet/goDo/draw/routing"
	userRouter "github.com/mcfiet/goDo/user/routing"
)

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // URL des React-Frontends
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Erlaube Anfrage-Header
		ExposedHeaders:   []string{"Authorization"},                 // Erlaube Antwort-Header
		AllowCredentials: true,
	}))

	app := app.InitApp()
	todoRouter := todoRouter.DrawRouter(app.DrawController)
	authRouter := authRouter.AuthRouter(app.AuthHandler, app.UserController)
	userRouter := userRouter.UserRouter(app.UserController)

	r.Mount("/", authRouter)
	r.Mount("/draw", todoRouter)
	r.Mount("/users", userRouter)
	log.Println("Server starting on :3000")

	http.ListenAndServe(":3000", r)
}
