package controllers

import (
	"echo-framework/author/services"
	"echo-framework/helper"
	"echo-framework/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//GetAuthor ... Get all author
func GetListAuthor(c echo.Context) error {
	data, err := services.ListAuthor()
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, data)
		return c.JSON(http.StatusOK, resp)
	}
}

//GetAuthorByID ... Get all author by id
func GetAuthorByID(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	data, err := services.DetailAuthor(int32(idInt))
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, data)
		return c.JSON(http.StatusOK, resp)
	}
}

//CreateAuthor ... Create author
func CreateAuthor(c echo.Context) error {
	var dataAuthor models.Author
	c.Bind(&dataAuthor)
	createdData, err := services.CreateAuthor(dataAuthor)
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, createdData)
		return c.JSON(http.StatusOK, resp)
	}
}

//UpdateAuthor ... Update author
func UpdateAuthor(c echo.Context) error {
	var dataAuthor models.Author
	c.Bind(&dataAuthor)
	updatedData, err := services.UpdateAuthor(dataAuthor)
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, updatedData)
		return c.JSON(http.StatusOK, resp)
	}
}

//DeleteAuthor ... Delete author
func DeleteAuthor(c echo.Context) error {
	var dataAuthor models.Author
	c.Bind(&dataAuthor)
	deletedData, err := services.DeleteAuthor(dataAuthor)
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, deletedData)
		return c.JSON(http.StatusOK, resp)
	}
}

//LoginAuthor ... Login author
func LoginAuthor(c echo.Context) error {
	var dataAuthor models.Author
	c.Bind(&dataAuthor)
	token, err := services.LoginAuthor(dataAuthor)
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": "invalid email / password"})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, map[string]interface{}{"token": token})
		return c.JSON(http.StatusOK, resp)
	}
}
