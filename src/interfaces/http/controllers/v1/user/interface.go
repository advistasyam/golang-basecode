package userv1controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"alteacare/golang-basecode/src/drivers/cloudwatch"
)

type V1User struct {
	DB         *gorm.DB
	Cloudwatch *cloudwatch.Cloudwatch
}

type SuccessResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type iV1User interface {
	CreateUser(c echo.Context) error
	DetailUser(c echo.Context) error
	LoginUser(c echo.Context) error
	MeUser(c echo.Context) error
}

func New(v1User *V1User) iV1User {
	return v1User
}
