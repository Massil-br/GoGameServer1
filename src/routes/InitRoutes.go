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

}

func InitPostRoutes(e *echo.Echo) {
	InitPostAuthRoutes(e)
}

func InitPutRoutes(e *echo.Echo) {

}

func InitPatchRoutes(e *echo.Echo) {

}

func InitDeleteRoutes(e *echo.Echo) {

}
