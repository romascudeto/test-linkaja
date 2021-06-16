package controllers

import (
	"echo-framework/article/services"
	"echo-framework/helper"
	"echo-framework/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

//GetArticle ... Get all article
func GetListArticle(c echo.Context) error {

	data, err := services.ListArticle()
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, data)
		return c.JSON(http.StatusOK, resp)
	}
}

//GetArticleByID ... Get all article by id
func GetArticleByID(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	data, err := services.DetailArticle(int32(idInt))
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusNotFound, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, data)
		return c.JSON(http.StatusOK, resp)
	}
}

//CreateArticle ... Create article
func CreateArticle(c echo.Context) error {
	var dataArticle models.Article
	c.Bind(&dataArticle)
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
	claims := token.Claims.(jwt.MapClaims)
	var id int64
	for key, val := range claims {
		if key == "id" {
			idFloat64 := val.(float64)
			id = int64(idFloat64)
		}
	}
	dataArticle.AuthorID = id
	createdData, err := services.CreateArticle(dataArticle)
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, createdData)
		return c.JSON(http.StatusOK, resp)
	}
}

//UpdateArticle ... Update article
func UpdateArticle(c echo.Context) error {
	var dataArticle models.Article
	c.Bind(&dataArticle)
	updatedData, err := services.UpdateArticle(dataArticle)
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, updatedData)
		return c.JSON(http.StatusOK, resp)
	}
}

//DeleteArticle ... Delete article
func DeleteArticle(c echo.Context) error {
	var dataArticle models.Article
	c.Bind(&dataArticle)
	deletedData, err := services.DeleteArticle(dataArticle)
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, deletedData)
		return c.JSON(http.StatusOK, resp)
	}
}
