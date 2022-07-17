package v2controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"alteacare/golang-basecode/src/drivers/cloudwatch"
)

type V2 struct {
	DB         *gorm.DB
	Cloudwatch *cloudwatch.Cloudwatch
}

type iV2 interface {
	Ping(c echo.Context) error
}

func New(v2 *V2) iV2 {
	return v2
}
