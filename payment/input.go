package payment

type InputPayment struct {
	Id 	int `json:"id"`
	ProductId int `json:"productId" binding:"required"`
	PricePaid float64 `json:"pricePaid" binding:"required"`
}
