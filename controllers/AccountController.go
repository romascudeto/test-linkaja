package controllers

import (
	"linkaja/helper"
	"linkaja/models"
	"linkaja/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

//Get Account By Number
func GetAccountByNumber(c echo.Context) error {
	accountNumber := c.Param("account_number")
	data, err := services.DetailAccount(accountNumber)
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusNotFound, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, data)
		return c.JSON(http.StatusOK, resp)
	}
}

//Transfer Balance
func TransferBalance(c echo.Context) error {
	fromAccountNumber := c.Param("from_account_number")
	var dataTransfer models.Transfer
	c.Bind(&dataTransfer)
	dataTransfer.FromAccountNumber = fromAccountNumber
	err := services.TransferBalance(dataTransfer)
	if err != nil {
		respErr := helper.ResponseError(map[string]interface{}{"message": err.Error()})
		return c.JSON(http.StatusBadRequest, respErr)
	} else {
		resp := helper.ResponseSuccess(nil, dataTransfer)
		return c.JSON(http.StatusCreated, resp)
	}
}
