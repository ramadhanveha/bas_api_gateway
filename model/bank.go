package model

type Bank struct {
	Bank_code string `"gorm: "primarykey"`
	Name      string
	Address   string
}

func (a *Bank) TableName() string {
	return "bank"
}
