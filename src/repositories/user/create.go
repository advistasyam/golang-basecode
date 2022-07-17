package userrepository

import (
	"alteacare/golang-basecode/src/drivers/gorm/models"
	"alteacare/golang-basecode/src/entities"
)

type (
	ParamsCreateUser struct {
		Name     string
		Email    string
		Password string
	}
)

func (i *sUserRepository) Create(p *ParamsCreateUser) (*entities.User, error) {
	user := models.User{
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
	}

	result := i.db.Model(i.model).Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities.ToUserEntity(&user), nil
}
