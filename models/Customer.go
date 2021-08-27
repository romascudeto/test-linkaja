package models

type Customer struct {
	ID             int    `json:"-" gorm:"primaryKey"`
	CustomerNumber int    `json:"customer_number"`
	Name           string `json:"name"`
}

func (b *Customer) TableName() string {
	return "customer"
}
