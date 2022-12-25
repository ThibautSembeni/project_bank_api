package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"project_api/utils/token"
)

type Repository interface {
	Login(user User) error
	Register(user User) error
	LoginCheck(user User) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(user User) error {
	errFind := r.db.First(&user).Error
	if errors.Is(errFind, gorm.ErrRecordNotFound) {
		return errFind
	}
	return nil
}

func (r *repository) Register(user User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r *repository) LoginCheck(user User) (string, error) {
	var err error
	u := User{}

	err = r.db.Model(User{}).Where("username = ?", user.Username).Take(&u).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(user.Password, u.Password)
	if err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	tokenJWT, _ := token.GenerateJWT(user.Username)
	return "Bearer " + tokenJWT, nil

}
