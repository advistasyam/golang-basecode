package userusecase

import (
	"alteacare/golang-basecode/src/entities"
	"alteacare/golang-basecode/src/helpers"
	userrepository "alteacare/golang-basecode/src/repositories/user"
)

type (
	ParamsCreateUser struct {
		Name     string
		Email    string
		Password string
	}
)

func (i *sUserUsecase) CreateUser(p *ParamsCreateUser) (*entities.User, error) {
	checkEmail, _ := i.userRepository.FindByEmail(&p.Email)

	if checkEmail != nil {
		return nil, ErrEmailAlreadyUsed
	}

	hashedPassword, _ := helpers.HashPassword(p.Password)
	data, err := i.userRepository.Create(&userrepository.ParamsCreateUser{
		Name:     p.Name,
		Email:    p.Email,
		Password: hashedPassword,
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}
