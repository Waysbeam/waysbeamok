package productdto

type CreateProduct struct {
	Title    string `json:"title" form:"title" validate:"required"`
	Price    int    `json:"price" form:"price" validate:"required"`
	Image    string `json:"image" form:"image" `
	Desc 	 string `json:"desc" form:"desc" validate:"required"`
	QTY      int    `json:"qty" form:"qty" validate:"required"`
}
type UpdateProduct struct {
	Title    string `json:"title" form:"title" `
	Price    int    `json:"price" form:"price" `
	Image    string `json:"image" form:"image" `
	Desc	 string `json:"desc" form:"desc" `
	QTY      int    `json:"qty" form:"qty" `
}
type ProductResponse struct {
	Title    string `json:"title" form:"title"`
	Price    int    `json:"price" form:"price"`
	Image    string `json:"image" form:"image"`
	Desc 	 string `json:"desc" form:"desc"`
	QTY      int    `json:"qty" form:"qty"`
}
