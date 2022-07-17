package v1controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"alteacare/golang-basecode/src/drivers/cloudwatch"
)

type V1 struct {
	DB         *gorm.DB
	Cloudwatch *cloudwatch.Cloudwatch
}

type iV1 interface {
	Ping(c echo.Context) error
}

func New(v1 *V1) iV1 {
	return v1
}
