package orm

import "time"

type Order struct {
	ID uint `gorm:"primaryKey" json:"id"`
	UserId uint `gorm:"not null" json:"user_id,omitempty"`
	ProductID uint `json:"product_id"`
	Product Product `json:"product,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}