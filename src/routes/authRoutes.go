package routes

import (
	"github.com/Massil-br/GoGameServer1/src/controllers"
	"github.com/labstack/echo/v4"
)

func InitPostAuthRoutes(e *echo.Echo) {
	e.POST("/api/login",controllers.Login)
	e.POST(("/api/register"), controllers.CreateUser)
}
