package userusecase

import (
	"alteacare/golang-basecode/src/entities"
)

func (i *sUserUsecase) DetailUser(id *uint) (*entities.User, error) {
	data, _ := i.userRepository.FindById(id)

	if data == nil {
		return nil, ErrUserNotFound
	}

	return data, nil
}
