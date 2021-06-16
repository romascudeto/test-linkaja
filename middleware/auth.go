package middleware

import (
	"echo-framework/helper"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			data := map[string]interface{}{
				"message": "not authorized",
				"status":  "token not found",
			}
			respErr := helper.ResponseError(data)
			return c.JSON(http.StatusUnauthorized, respErr)
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("secret"), nil
		})
		if err != nil {
			data := map[string]interface{}{
				"message": "not authorized",
				"status":  err.Error(),
			}
			respErr := helper.ResponseError(data)
			return c.JSON(http.StatusUnauthorized, respErr)
		}
		claims := token.Claims.(jwt.MapClaims)
		var expiredAt interface{}
		for key, val := range claims {
			if key == "expiredAt" {
				expiredAt = val
			}
		}
		if time.Now().Unix() > int64(expiredAt.(float64)) {
			data := map[string]interface{}{
				"message": "not authorized",
				"status":  "token expired",
			}
			respErr := helper.ResponseError(data)
			return c.JSON(http.StatusUnauthorized, respErr)
		}
		if token != nil && err == nil {
		} else {
			data := map[string]interface{}{
				"message": "not authorized",
				"status":  err.Error(),
			}
			respErr := helper.ResponseError(data)
			return c.JSON(http.StatusUnauthorized, respErr)
		}
		return next(c)
	}
}
