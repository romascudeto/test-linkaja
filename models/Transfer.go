package models

type Transfer struct {
	FromAccountNumber string `json:"from_account_number"`
	ToAccountNumber   string `json:"to_account_number"`
	Amount            float32
}
