package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Massil-br/GoGameServer1/src/config"
	"github.com/Massil-br/GoGameServer1/src/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var roleHierarchy = map[string]int{
	"user":  1,
	"admin": 2,
}

func AuthMiddleware(minRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Authentification (JWT)
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid Authorization header")
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			secret := os.Getenv("JWT_SECRET")

			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(secret), nil
			})
			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
			}

			idStr, ok := claims["id"].(string)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid id in token")
			}
			userID, err := uuid.Parse(idStr)

			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "error binding user ID from token")
			}

			var user models.User
			if err := config.DB.First(&user, userID).Error; err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "User not found")
			}

			// Autorisation (rôle)
			userLevel, okUser := roleHierarchy[user.Role]
			minLevel, okMin := roleHierarchy[minRole]
			if !okUser || !okMin {
				return echo.ErrForbidden
			}
			if userLevel < minLevel {
				return echo.ErrForbidden
			}

			c.Set("user", &user)

			return next(c)
		}
	}
}
