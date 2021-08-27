package models

type Account struct {
	ID             int      `json:"-" gorm:"primaryKey"`
	AccountNumber  int      `json:"account_number" gorm:"primaryKey"`
	CustomerNumber int      `json:"customer_number"`
	Balance        float32  `json:"balance"`
	Customer       Customer `gorm:"ForeignKey:CustomerNumber;AssociationForeignKey:CustomerNumber"`
}

func (b *Account) TableName() string {
	return "account"
}
