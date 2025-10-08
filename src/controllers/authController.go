package controllers

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/Massil-br/GoGameServer1/src/config"
	"github.com/Massil-br/GoGameServer1/src/models"
	authmodels "github.com/Massil-br/GoGameServer1/src/models/AuthModels"
	"github.com/Massil-br/GoGameServer1/src/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateUser(c echo.Context) error {
	var req authmodels.CreateUserRequest
	err := c.Bind(&req)
	if err != nil {
		log.Println("[WARN] could not  bind CreateUserRequest: ", err, " Sender : ", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input,"})
	}

	if req.Password != req.ConfirmPassword {
		log.Println("[WARN] password & confirmPassword don't match,  sender : ", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Password and confirmPassword don't match"})
	}

	if len(req.Password) < 8 || !regexp.MustCompile("[0-9]").MatchString(req.Password) {
		log.Println("[WARN] Incorrect password size & requirements, sender : ", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Password must be at least 8 characters & contain at least one digit"})
	}

	var existingUser models.User

	err = config.DB.Unscoped().Where("email = ?", req.Email).First(&existingUser).Error

	if err == nil {
		if existingUser.DeletedAt.Valid {
			if err := config.DB.Unscoped().Delete(&existingUser).Error; err != nil {
				log.Println("[ERROR], Failed to permanently delete user : ", err)
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to permanently delete user"})
			}
		} else {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Email already in use"})
		}
	} else if err != gorm.ErrRecordNotFound {
		log.Println("[ERROR], unknown error when searching for existing user ", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Println("[ERROR] Could not hash password : ", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not hash password"})
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "user",
	}

	err = config.DB.Create(&user).Error
	if err != nil {
		log.Println("[ERROR] Could not create user : ", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not create user"})
	}

	log.Println("[OK]User Created ", user)

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User created successfully",
		"user":    user,
	})

}

func Login(c echo.Context) error {
	var req authmodels.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		log.Println("[WARN] could not  bind LoginRequest,  ", err, " Sender : ", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	var user models.User

	err = config.DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		log.Println("[WARN] Invalid creadentials, Sender : ", c.RealIP())
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		log.Println("[WARN] Invalid crendentials, Sender : ", c.RealIP())
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("[ERROR] failed to generate token  : ", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	log.Println("[OK] Login successfull , Sender : ", c.RealIP())

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successfull",
		"token":   tokenString,
	})

}
