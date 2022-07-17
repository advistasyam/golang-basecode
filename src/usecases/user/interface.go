package userusecase

import (
	"alteacare/golang-basecode/src/entities"
	user "alteacare/golang-basecode/src/repositories/user"
)

type sUserUsecase struct {
	userRepository user.UserRepository
}

type UserUsecase interface {
	CreateUser(*ParamsCreateUser) (*entities.User, error)
	DetailUser(*uint) (*entities.User, error)
	Login(*ParamsLogin) (*ResultLogin, error)
}

func New(userRepository user.UserRepository) UserUsecase {
	return &sUserUsecase{userRepository: userRepository}
}
