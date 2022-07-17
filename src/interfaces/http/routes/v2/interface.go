package v2routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"alteacare/golang-basecode/src/drivers/cloudwatch"
)

type V2Routes struct {
	Echo       *echo.Group
	DB         *gorm.DB
	Cloudwatch *cloudwatch.Cloudwatch
}

type iV2Routes interface {
	MountPing()
}

func New(s *V2Routes) iV2Routes {
	return s
}
