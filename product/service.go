package product

type Service interface {
	Store(input InputProduct) (Product, error)
	Update(id int, input InputProduct) (Product, error)
	Delete(id int) error
	Find(id string) ([]Product, error)
	ListAll() ([]Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input InputProduct) (Product, error) {
	product := Product{}
	product.Name = input.Name
	product.Price = input.Price

	newProduct, err := s.repository.Insert(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) Update(id int, input InputProduct) (Product, error) {
	product := Product{}
	product.ID = id
	product.Name = input.Name
	product.Price = input.Price

	product, err := s.repository.Update(product)
	if err != nil {
		return product, err
	}

	return product, nil
}
func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Find(id string) ([]Product, error) {
	product, err := s.repository.Find(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) ListAll() ([]Product, error) {
	products, err := s.repository.ListAll()
	if err != nil {
		return products, err
	}

	return products, nil
}