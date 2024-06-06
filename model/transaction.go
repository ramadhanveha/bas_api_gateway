package model

import "time"

type Transaction struct {
	Id               int `gorm:"primarykey"`
	Account_id       string
	Bank_id          string
	Amount           int
	Transaction_date *time.Time
}

func (a *Transaction) TableName() string {
	return "transaction"
}
