package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterInput) (Users, error)
	Login(input LoginInput) (Users, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterInput) (Users, error) {
	user := Users{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	Pswd, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(Pswd)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil	
}

func (s *service) Login(input LoginInput) (Users, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}