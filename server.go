package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	Config "linkaja/config"
	"linkaja/models"
	"linkaja/routes"
	"linkaja/seeders"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var err error

func main() {
	godotenv.Load()
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&models.Account{}, &models.Customer{})
	seeders.InsertDefaultSeeder()
	if err != nil {
		fmt.Println("Status:", err)
	}
	e := echo.New()
	routes.Routes(e)
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("routes.json", data, 0644)
	e.Logger.Fatal(e.Start(":1323"))
}
