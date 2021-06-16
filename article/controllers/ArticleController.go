package controllers

import (
	"echo-framework/article/services"
	"echo-framework/helper"
	"echo-framework/models"
	"net/http"
	"strconv"

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
