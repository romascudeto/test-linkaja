package services

import (
	"errors"
	Config "linkaja/config"
	"linkaja/models"
	"linkaja/repositories"
)

func DetailAccount(accountNumber string) (dataRetMap map[string]interface{}, err error) {
	var dataRet models.Account
	err = repositories.Detail(&dataRet, accountNumber)
	dataRetMap = map[string]interface{}{
		"account_number": dataRet.AccountNumber,
		"customer_name":  dataRet.Customer.Name,
		"balance":        dataRet.Balance,
	}
	return dataRetMap, err
}

func TransferBalance(dataTransfer models.Transfer) (err error) {
	tx := Config.DB.Begin()

	err = SubstractBalanceFrom(dataTransfer)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = AddBalanceTo(dataTransfer)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

func SubstractBalanceFrom(dataTransfer models.Transfer) (err error) {
	var dataFrom models.Account
	err = repositories.Detail(&dataFrom, dataTransfer.FromAccountNumber)
	if err != nil {
		return err
	}
	dataFrom.Balance = dataFrom.Balance - dataTransfer.Amount
	if dataFrom.Balance < 0 {
		return errors.New("Out of balance !")
	}
	err = repositories.Update(&dataFrom)
	if err != nil {
		return err
	}
	return nil
}

func AddBalanceTo(dataTransfer models.Transfer) (err error) {
	var dataTo models.Account
	err = repositories.Detail(&dataTo, dataTransfer.ToAccountNumber)
	if err != nil {
		return err
	}
	dataTo.Balance = dataTo.Balance + dataTransfer.Amount
	repositories.Update(&dataTo)
	if err != nil {
		return err
	}
	return nil
}
