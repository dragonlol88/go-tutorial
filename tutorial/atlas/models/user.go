package models

import "time"


type User struct {
	ID 		  uint 		`gorm:"primaryKey" json:"id"`
	Name 	  string	`gorm:"index" json:"name"`
	Orders []*Order		`json:"orders"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}