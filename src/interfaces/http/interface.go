package http

import (
	"gorm.io/gorm"

	"alteacare/golang-basecode/src/drivers/cloudwatch"
)

type Http struct {
	DB         *gorm.DB
	Cloudwatch *cloudwatch.Cloudwatch
}

type iHttp interface {
	Launch()
}

func New(http *Http) iHttp {
	return http
}
