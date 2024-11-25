package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/mcfiet/goDo/app"
	authRouter "github.com/mcfiet/goDo/auth/routing"
	todoRouter "github.com/mcfiet/goDo/draw/routing"
	userRouter "github.com/mcfiet/goDo/user/routing"
)

func main() {
	_ = godotenv.Load("../.env")

	r := chi.NewRouter()
	api := chi.NewRouter()

	app := app.InitApp()
	todoRouter := todoRouter.DrawRouter(app.DrawController)
	authRouter := authRouter.AuthRouter(app.AuthHandler, app.UserController)
	userRouter := userRouter.UserRouter(app.UserController)

	api.Mount("/", authRouter)
	api.Mount("/draw", todoRouter)
	api.Mount("/users", userRouter)
	r.Mount("/api", api)
	log.Println("Server starting on :3000")

	http.ListenAndServe(":3000", r)
}
