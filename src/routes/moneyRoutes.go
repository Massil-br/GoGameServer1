package routes

import (
	"github.com/Massil-br/GoGameServer1/src/controllers"
	"github.com/Massil-br/GoGameServer1/src/middlewares"
	"github.com/labstack/echo/v4"
)

func InitPostMoneyRoutes(e *echo.Echo) {
	e.POST("/api/money", controllers.CreateMoney, middlewares.AuthMiddleware("user"))

}

func InitPatchMoneyRoutes(e *echo.Echo) {
	e.PATCH("/api/money/add", controllers.AddMoney, middlewares.AuthMiddleware("admin"))
	e.PATCH("/api/money/remove", controllers.RemoveMoney, middlewares.AuthMiddleware("admin"))
}

func InitPutMoneyRoutes(e *echo.Echo) {
	e.PUT("/api/money/set", controllers.SetMoney, middlewares.AuthMiddleware("admin"))

}

func InitGetMoneyRoutes(e *echo.Echo) {
	e.GET("/api/money", controllers.GetMoney, middlewares.AuthMiddleware("user"))
}
