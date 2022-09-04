package models

type Cart struct {
	ID            int                 `json:"id"`
	QTY           int                 `json:"qty"`
	SubTotal      int                 `json:"subtotal"`
	ProductID     int                 `json:"product_id"`
	Product       ProductCart         `json:"product"`
	TransactionID int                 `json:"transaction_id"`
	Transaction   TransactionResponse `json:"transaction"`
	Status        string              `json:"status"`
}
type CartResponse struct {
	ID            int         `json:"id"`
	Total         int         `json:"total"`
	TransactionID int         `json:"transaction_id"`
	ProductID     int         `json:"product_id"`
	Product       ProductCart `json:"product"`
}

func (CartResponse) TableName() string {
	return "carts"
}