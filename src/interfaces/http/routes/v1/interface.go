package v1routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"alteacare/golang-basecode/src/drivers/cloudwatch"
)

type V1Routes struct {
	Echo       *echo.Group
	DB         *gorm.DB
	Cloudwatch *cloudwatch.Cloudwatch
}

type iV1Routes interface {
	MountPing()
	MountUser()
}

func New(v1Routes *V1Routes) iV1Routes {
	return v1Routes
}
