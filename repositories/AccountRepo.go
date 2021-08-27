package repositories

import (
	Config "linkaja/config"
	"linkaja/models"

	_ "github.com/go-sql-driver/mysql"
)

func Detail(account *models.Account, accountNumber string) (err error) {

	if err = Config.DB.Debug().Preload("Customer").Where("account_number = ?", accountNumber).First(account).Error; err != nil {
		return err
	}
	return nil
}

func Update(account *models.Account) (err error) {
	if err = Config.DB.Model(&account).Updates(account).Error; err != nil {
		return err
	}
	return nil
}
