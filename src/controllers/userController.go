package controllers

import (
	"net/http"

	"github.com/Massil-br/GoGameServer1/src/models"
	"github.com/labstack/echo/v4"
)

func GetUserId(c echo.Context) error {
	user := c.Get("user").(*models.User)

	return c.JSON(http.StatusOK, echo.Map{"message": "this is your user ID", "userId": user.ID})

}
