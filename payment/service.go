package payment

type Service interface {
	Store(input InputPayment) (Payment, error)
	Update(id int, input InputPayment) (Payment, error)
	Delete(id int) error
	Find(id string) ([]Payment, error)
	ListAll() ([]Payment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input InputPayment) (Payment, error) {
	payment := Payment{}
	payment.ProductId = input.ProductId
	payment.PricePaid = input.PricePaid

	newPayment, err := s.repository.Insert(payment)
	if err != nil {
		return newPayment, err
	}

	return newPayment, nil
}

func (s *service) Update(id int, input InputPayment) (Payment, error) {
	payment := Payment{}
	payment.ID = id
	payment.ProductId = input.ProductId
	payment.PricePaid = input.PricePaid
	payment, err := s.repository.Update(payment)
	if err != nil {
		return payment, err
	}

	return payment, nil
}
func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Find(id string) ([]Payment, error) {
	payment, err := s.repository.Find(id)
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (s *service) ListAll() ([]Payment, error) {
	payments, err := s.repository.ListAll()
	if err != nil {
		return payments, err
	}

	return payments, nil
}