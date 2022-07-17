package http_internal

import (
	"gorm.io/gorm"

	"alteacare/golang-basecode/src/drivers/cloudwatch"
)

type HttpInternal struct {
	DB         *gorm.DB
	Cloudwatch *cloudwatch.Cloudwatch
}

type iHttpInternal interface {
	Launch()
}

func New(httpInternal *HttpInternal) iHttpInternal {
	return httpInternal
}
