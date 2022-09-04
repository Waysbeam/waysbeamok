package models

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	User      User      `json:"user"`
	Status    string    `json:"status"`
	Total     int       `json:"total"`
	Carts     []Cart    `json:"carts"`
	CreatedAt time.Time `json:"-"`
	UpdateAt  time.Time `json:"-"`
}
type TransactionResponse struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}

func (TransactionResponse) TableName() string{
	return"transactions"
}
