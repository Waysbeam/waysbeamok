package models

import "time"

type Product struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" gorm:"type : varchar(255)"`
	Price     int       `json:"price" gorm:"type :int"`
	Desc      string    `json:"desc" gorm:"type :varchar(255)"`
	QTY       int       `json:"qty" gorm:"type:int"`
	Image     string    `json:"image" gorm:"type : varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdateAt  time.Time `json:"-"`
}

type ProductCart struct {
	ID    int    `json:"id"`
	Title string `json:"Title"`
	Desc  string `json:"desc"`
	Price int    `json:"Price"`
	Image string `json:"Image"`
}

func (ProductCart) TableName() string {
	return "products"
}
