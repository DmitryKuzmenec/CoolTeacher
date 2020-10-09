package main

import (
	"log"

	"github.com/DmitryKuzmenec/CoolTeacher/config"
	"github.com/DmitryKuzmenec/CoolTeacher/handlers"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
)

func main() {
	config, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(config)
	admin := handlers.NewAdmin()
	e := echo.New()
	e.GET("/", admin.Home)
	e.Logger.Fatal(e.Start(":8080"))
}
