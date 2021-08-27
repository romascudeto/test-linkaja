package seeders

import (
	Config "linkaja/config"
	"linkaja/models"
	"log"
)

var customers = []models.Customer{
	{
		CustomerNumber: 1001,
		Name:           "Bob Martin",
	},
	{
		CustomerNumber: 1002,
		Name:           "Linus Torvalds",
	},
}

var accounts = []models.Account{
	{
		AccountNumber:  555001,
		CustomerNumber: 1001,
		Balance:        10000,
	},
	{
		AccountNumber:  555002,
		CustomerNumber: 1002,
		Balance:        15000,
	},
}

func InsertDefaultSeeder() {

	for i, _ := range customers {
		var checkCustomer models.Customer
		Config.DB.Debug().Where("customer_number = ?", customers[i].CustomerNumber).First(&checkCustomer)
		if checkCustomer.CustomerNumber == 0 {
			err := Config.DB.Debug().Model(&models.Customer{}).Create(&customers[i]).Error
			if err != nil {
				log.Fatalf("cannot seed customer table: %v", err)
			}
		}
	}

	for i, _ := range accounts {
		var checkAccount models.Account
		Config.DB.Debug().Where("account_number = ?", accounts[i].AccountNumber).First(&checkAccount)
		if checkAccount.AccountNumber == 0 {
			err := Config.DB.Debug().Model(&models.Account{}).Create(&accounts[i]).Error
			if err != nil {
				log.Fatalf("cannot seed account table: %v", err)
			}
		}
	}
}
