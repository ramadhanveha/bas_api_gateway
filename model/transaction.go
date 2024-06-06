package model

import "time"

type Transaction struct {
	Id               int    `gorm:"primarykey"`
	Account_id       string `gorm:"foreignkey"`
	Bank_id          string `gorm:"foreignkey"`
	Amount           int    `gorm:"column:amount"`
	Transaction_date *time.Time
}

func (a *Transaction) TableName() string {
	return "transaction"
}
