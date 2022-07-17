package userrepository

import (
	"gorm.io/gorm"

	"alteacare/golang-basecode/src/drivers/gorm/models"
	"alteacare/golang-basecode/src/entities"
)

type sUserRepository struct {
	db    *gorm.DB
	model *models.User
}

type UserRepository interface {
	Create(*ParamsCreateUser) (*entities.User, error)
	FindByEmail(*string) (*entities.User, error)
	FindById(*uint) (*entities.User, error)
}

func New(db *gorm.DB) UserRepository {
	return &sUserRepository{db: db, model: &models.User{}}
}
