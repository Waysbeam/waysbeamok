package cart

type CreateCart struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	SubTotal  *int   `json:"subtotal"`
	Status    string `json:"status"`
	ProductID int    `json:"product_id"`
}
type UpdateQTYRequest struct {
	QTY      int `json:"qty"`
	SubTotal int `json:"subtotal"`
}
type UpdateCartRequest struct {
	TransactionID int `json:"transaction_id"`
}
type CartResponse struct {
	ID       int `json:"id"`
	SubTotal int `json:"subtotal"`
}