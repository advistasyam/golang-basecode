package userrepository

import (
	"alteacare/golang-basecode/src/drivers/gorm/models"
	"alteacare/golang-basecode/src/entities"
)

func (i *sUserRepository) FindById(id *uint) (*entities.User, error) {
	user := models.User{}
	dbresult := i.db.Model(i.model).Where("id", id).First(&user)

	if dbresult.Error != nil {
		return nil, dbresult.Error
	}

	return entities.ToUserEntity(&user), nil
}
