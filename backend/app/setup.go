package app

import (
	"github.com/mcfiet/goDo/auth/handlers"
	"github.com/mcfiet/goDo/db"
	drawController "github.com/mcfiet/goDo/draw/controller"
	drawRepository "github.com/mcfiet/goDo/draw/repository"
	drawService "github.com/mcfiet/goDo/draw/service"
	userController "github.com/mcfiet/goDo/user/controller"
	userRepository "github.com/mcfiet/goDo/user/repository"
	userService "github.com/mcfiet/goDo/user/service"
	"gorm.io/gorm"
)

type App struct {
	DB             *gorm.DB
	DrawController *drawController.DrawController
	UserController *userController.UserController
	AuthHandler    *handlers.AuthHandler
}

func InitApp() *App {
	db := db.Init()

	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	userController := userController.NewUserController(userService)

	drawRepository := drawRepository.NewDrawRepository(db)
	drawService := drawService.NewDrawService(drawRepository)
	drawController := drawController.NewDrawController(drawService, userService)

	authHandler := handlers.NewAuthHandler(userService)
	return &App{
		DB:             db,
		DrawController: drawController,
		UserController: userController,
		AuthHandler:    authHandler,
	}
}
