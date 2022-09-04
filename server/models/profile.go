package models

import "time"

type Profile struct {
	ID         int         `json:"id"`
	Image      string      `json:"image" gorm:"type :varchar(255)"`
	Address    string      `json:"address" gorm:"type: varchar(255)"`
	City       string      `json:"city" gorm:"type :varchar(255)"`
	PostalCode string      `json:"postal_code" gorm:"type: int"`
	UserId     int         `json:"user_id" ` //make id when register
	User       UserProfile `json:"user"`     //relasi user
	CreatedAt  time.Time   `json:"-"`
	UpdateAt   time.Time   `json:"-"`
}

type ProfileResponse struct {
	ID         int    `json:"id"`
	Image      string `json:"image"`
	Address    string `json:"address"`
	City       string `json:"city" `
	PostalCode string `json:"postal_code"`
	UserId     int    `json:"user_id"`
}
func (ProfileResponse) TableName() string{
	return "Profile"
}