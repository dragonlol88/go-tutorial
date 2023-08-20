package models

import "time"

type Product struct {
	ID 		  uint 		`gorm:"primaryKey" json:"id"`
	Name 	  string	`gorm:"index" json:"name"`
	Price	  float32	`gorm:"type:decimal(9,2)" json:"price"`
	Code 	  *string	`gorm:"unique;not null" json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}