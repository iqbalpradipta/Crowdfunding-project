package users

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterInput) (Users, error)
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