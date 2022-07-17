package userrepository

import (
	"alteacare/golang-basecode/src/drivers/gorm/models"
	"alteacare/golang-basecode/src/entities"
)

func (i *sUserRepository) FindByEmail(email *string) (*entities.User, error) {
	user := models.User{}
	dbresult := i.db.Model(i.model).Where("email", email).First(&user)

	if dbresult.Error != nil {
		return nil, dbresult.Error
	}

	return entities.ToUserEntity(&user), nil
}
