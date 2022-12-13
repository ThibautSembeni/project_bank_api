package product

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	Insert(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(id int) error
	Find(id string) ([]Product, error)
	ListAll() ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	errFind := r.db.First(&product, product.ID).Error
	if errors.Is(errFind, gorm.ErrRecordNotFound) {
		return product, errFind
	}
	err := r.db.Model(&product).Updates(product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Delete(id int) error {
	errFind := r.db.First(&Product{}, id).Error
	if errors.Is(errFind, gorm.ErrRecordNotFound) {
		return errFind
	}
	err := r.db.Delete(&Product{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Find(id string) ([]Product, error) {
	var product []Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *repository) ListAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}