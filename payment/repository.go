package payment

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	Insert(payment Payment) (Payment, error)
	Update(payment Payment) (Payment, error)
	Delete(id int) error
	Find(id string) ([]Payment, error)
	ListAll() ([]Payment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(payment Payment) (Payment, error) {
	err := r.db.Create(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *repository) Update(payment Payment) (Payment, error) {
	currentPayment := payment
	errFind := r.db.First(&payment, payment.ID).Error
	if errors.Is(errFind, gorm.ErrRecordNotFound) {
		return payment, errFind
	}
	currentPayment.ProductId = payment.ProductId
	err := r.db.Model(&currentPayment).Updates(currentPayment).Error
	if err != nil {
		return currentPayment, err
	}

	return currentPayment, nil
}

func (r *repository) Delete(id int) error {
	errFind := r.db.First(&Payment{}, id).Error
	if errors.Is(errFind, gorm.ErrRecordNotFound) {
		return errFind
	}
	err := r.db.Delete(&Payment{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Find(id string) ([]Payment, error) {
	var payment []Payment
	err := r.db.First(&payment, id).Error
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (r *repository) ListAll() ([]Payment, error) {
	var payments []Payment
	err := r.db.Find(&payments).Error
	if err != nil {
		return nil, err
	}

	return payments, nil
}
