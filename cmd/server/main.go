package main

import (
	"github.com/DmitryKuzmenec/CoolTeacher/handlers"
	"github.com/DmitryKuzmenec/CoolTeacher/logger"
	"github.com/DmitryKuzmenec/CoolTeacher/services"

	"github.com/DmitryKuzmenec/CoolTeacher/repositories"

	"github.com/DmitryKuzmenec/CoolTeacher/client/database"
	"github.com/DmitryKuzmenec/CoolTeacher/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config, err := config.Init()
	if err != nil {
		panic(err)
	}
	log, err := logger.Init(config)
	if err != nil {
		panic(err)
	}
	db, err := database.Dial(config)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService, log)
	//admin := handlers.NewUsers()
	e := echo.New()

	// frontend
	e.Static("/", "frontend/build")

	users := e.Group("/users")
	users.POST("/create", userHandler.Create)
	users.GET("/list", userHandler.List)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "frontend/cool-teacher/build",
		Index: "index.html",
	}))
	e.Logger.Fatal(e.Start(":8080"))
}
