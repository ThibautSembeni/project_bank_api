package user

type Service interface {
	Login(input InputUser) error
	Register(input InputUser) error
	LoginCheck(input InputUser) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Login(input InputUser) error {
	user := User{}
	user.Username = input.Username
	user.Password = input.Password

	err := s.repository.Login(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Register(input InputUser) error {
	user := User{}
	user.Username = input.Username
	user.Password = input.Password

	err := s.repository.Register(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) LoginCheck(input InputUser) (string, error) {
	user := User{}
	user.Username = input.Username
	user.Password = input.Password

	token, err := s.repository.LoginCheck(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
