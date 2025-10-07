package main

import (
	"github.com/Massil-br/GoGameServer1/src/config"
	"github.com/Massil-br/GoGameServer1/src/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	config.Init()
	file := config.InitLogger()
	defer file.Close()
	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":9090"))
}
