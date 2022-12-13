package payment

import (
	"gorm.io/gorm"
	"project_api/product"
	"time"
)

type Payment struct {
	ID        int             `json:"id"`
	ProductId int             `json:"-"`
	PricePaid float64         `json:"pricePaid"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	Product   product.Product `json:"product" gorm:"foreignkey:ProductId"`
}

func (p *Payment) AfterSave(db *gorm.DB) (err error) {
	return db.Limit(1).Find(&p.Product, p.ProductId).Error
}

func (p *Payment) AfterFind(db *gorm.DB) (err error) {
	return db.Limit(1).Find(&p.Product, p.ProductId).Error
}
