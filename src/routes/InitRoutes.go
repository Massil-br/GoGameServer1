package routes

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {
	InitGetRoutes(e)
	InitPostRoutes(e)
	InitPutRoutes(e)
	InitPatchRoutes(e)
	InitDeleteRoutes(e)
}

func InitGetRoutes(e *echo.Echo) {
	InitGetMoneyRoutes(e)
	InitGetUserRoutes(e)
}

func InitPostRoutes(e *echo.Echo) {
	InitPostAuthRoutes(e)
	InitPostMoneyRoutes(e)
}

func InitPutRoutes(e *echo.Echo) {
	InitPutMoneyRoutes(e)
}

func InitPatchRoutes(e *echo.Echo) {
	InitPatchMoneyRoutes(e)
}

func InitDeleteRoutes(e *echo.Echo) {

}
