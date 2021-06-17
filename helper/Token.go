package helper

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func ParsingTokenJWT(c echo.Context) map[string]interface{} {

	tokenString := c.Request().Header.Get("Authorization")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	mappingToken := make(map[string]interface{})

	for key, val := range claims {
		mappingToken[key] = val
	}
	return mappingToken
}
