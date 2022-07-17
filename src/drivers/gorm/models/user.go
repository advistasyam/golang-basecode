package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;"`
	Name      string `gorm:"type:varchar(255);"`
	Email     string `gorm:"type:varchar(255);index;"`
	Password  string `gorm:"type:varchar(255);"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (User) TableName() string {
	return "users"
}

func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
