package product

type InputProduct struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Price 		float64 `json:"price" binding:"required"`
}