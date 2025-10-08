package controllers

import (
	"log"
	"net/http"

	"github.com/Massil-br/GoGameServer1/src/config"
	"github.com/Massil-br/GoGameServer1/src/models"
	moneymodels "github.com/Massil-br/GoGameServer1/src/models/MoneyModels"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetMoney(c echo.Context) error {
	user := c.Get("user").(*models.User)

	var money models.Money

	err := config.DB.Where("user_id = ?", user.ID).First(&money).Error
	if err != nil {
		log.Println("[WARN]User's money not found , userId : ", user.ID, " error:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "you don't have money data attributed in database or unknown error"})
	}

	log.Println("[OK] get money successfull , Sender : ", c.RealIP())
	return c.JSON(http.StatusOK, echo.Map{
		"message": "getMoney successfull",
		"money":   money.Amount,
	})

}

func CreateMoney(c echo.Context) error {
	user := c.Get("user").(*models.User)

	var money models.Money

	err := config.DB.Where("user_id = ?", user.ID).First(&money).Error
	if err == nil {
		log.Println("[WARN] trying to create a money acount, but already exists, Sender : ", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "money account already exists"})
	}

	if err != gorm.ErrRecordNotFound {
		log.Println("[ERROR] unknown error while trying to find if money account exists, Error :  ", err, " Sender : ", c.RealIP())
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": " unknown error while checking if money account exists"})
	}

	money = models.Money{
		UserId: user.ID,
		Amount: 0,
	}

	err = config.DB.Create(&money).Error
	if err != nil {
		log.Println("[ERROR] couldn't create money for user : ", user.ID, " Sender : ", c.RealIP())
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "couldn't create money data for your user"})
	}

	log.Println("[OK] money data created for user : ", user.ID, " Sender: ", c.RealIP())
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "money data created",
		"money":   money.Amount,
	})
}

func SetMoney(c echo.Context) error {
	user := c.Get("user").(*models.User)

	var req moneymodels.MoneyRequest

	err := c.Bind(&req)
	if err != nil {
		log.Println("[WARN] couldn't bind money amount :", err, " Sender : ", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if req.UserId != nil {
		err := config.DB.Where("id = ?", req.UserId).First(&user).Error
		if err != nil {
			log.Println("[WARN] Can't find user with id : ", req.UserId, " Sender : ", c.RealIP())
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Can't find user"})
		}
	}

	if req.Amount < 0 {
		log.Println("[WARN] req.amount can't be lower than 0, Sender: ", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Amount can't be 0 or lower"})
	}

	var money models.Money

	err = config.DB.Where("user_id = ?", user.ID).First(&money).Error
	if err != nil {
		log.Println("[WARN]User's money not found , userId : ", user.ID, " error:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "you don't have money data attributed in database or unknown error"})
	}

	money.Amount = req.Amount

	err = config.DB.Save(&money).Error
	if err != nil {
		log.Println("[ERROR] Couldn't save money in database  error:  ", err, " Sender : ", c.RealIP())
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to update money"})
	}

	log.Println("[OK] Set ", req.Amount, " money to ", user.ID)
	return c.JSON(http.StatusOK, echo.Map{
		"message": " money of target updated successfully",
		"amount":  req.Amount,
	})

}

func AddMoney(c echo.Context) error {
	user := c.Get("user").(*models.User)

	var req moneymodels.MoneyRequest

	err := c.Bind(&req)
	if err != nil {
		log.Println("[WARN] couldn't bind money amount :", err, " Sender : ", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if req.UserId != nil {
		err := config.DB.Where("id = ?", req.UserId).First(&user).Error
		if err != nil {
			log.Println("[WARN] Can't find user with id : ", req.UserId, " Sender : ", c.RealIP())
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Can't find user"})
		}
	}

	if req.Amount <= 0 {
		log.Println("[WARN] req.amount can't be 0 or lower, Sender: ", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Amount can't be 0 or lower"})
	}

	var money models.Money

	err = config.DB.Where("user_id = ?", user.ID).First(&money).Error
	if err != nil {
		log.Println("[WARN]User's money not found , userId : ", user.ID, " error:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "you don't have money data attributed in database or unknown error"})
	}

	money.Amount += req.Amount
	err = config.DB.Save(&money).Error
	if err != nil {
		log.Println("[ERROR] Couldn't save money in database  error:  ", err, " Sender : ", c.RealIP())
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to update money"})
	}

	log.Println("[OK] Added ", req.Amount, " money to ", user.ID)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "added money for target successfully",
		"amount":  req.Amount,
	})

}

func RemoveMoney(c echo.Context) error {
	user := c.Get("user").(*models.User)

	var req moneymodels.MoneyRequest

	err := c.Bind(&req)
	if err != nil {
		log.Println("[WARN] couldn't bind money amount :", err, " Sender : ", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if req.UserId != nil {
		err := config.DB.Where("id = ?", req.UserId).First(&user).Error
		if err != nil {
			log.Println("[WARN] Can't find user with id : ", req.UserId, " Sender : ", c.RealIP())
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Can't find user"})
		}
	}

	if req.Amount <= 0 {
		log.Println("[WARN] req.amount can't be 0 or lower, Sender: ", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Amount can't be 0 or lower"})
	}

	var money models.Money

	err = config.DB.Where("user_id = ?", user.ID).First(&money).Error
	if err != nil {
		log.Println("[WARN]User's money not found , userId : ", user.ID, " error:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "you don't have money data attributed in database or unknown error"})
	}

	money.Amount -= req.Amount
	err = config.DB.Save(&money).Error
	if err != nil {
		log.Println("[ERROR] Couldn't save money in database  error:  ", err, " Sender : ", c.RealIP())
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to update money"})
	}

	log.Println("[OK] Added ", req.Amount, " money to ", user.ID)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "removed money of target successfully",
		"amount":  req.Amount,
	})

}
