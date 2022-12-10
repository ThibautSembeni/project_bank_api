package product

import "time"

type Product struct {
	ID int `json:"id"`
	Name string `json:"name" binding:"required"`
	Price string `json:"price" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}