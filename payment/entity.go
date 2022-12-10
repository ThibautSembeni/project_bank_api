package payment

import (
	"project_api/product"
	"time"
)

type Payment struct {
	ID int `json:"id"`
	ProductId int `json:"productId"`
	PricePaid int `json:"pricePaid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ProductID int
	Product product.Product
}

