package models

import "time"

type Order struct {
	ID uint `gorm:"primaryKey" json:"id"`
	UserId *uint `gorm:"not null" json:"user_id"`
	ProductID uint `json:"product_id"`
	Product *Product `json:"product"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}