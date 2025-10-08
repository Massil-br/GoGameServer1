package routes

import (
	"github.com/Massil-br/GoGameServer1/src/controllers"
	"github.com/Massil-br/GoGameServer1/src/middlewares"
	"github.com/labstack/echo/v4"
)

func InitGetUserRoutes(e *echo.Echo) {
	e.GET("/api/user", controllers.GetUserId, middlewares.AuthMiddleware("user"))
}
