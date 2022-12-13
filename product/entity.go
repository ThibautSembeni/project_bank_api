package product

import (
	"time"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name" gorm:"unique"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}