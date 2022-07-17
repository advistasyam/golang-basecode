package v1routes

import (
	v1controller "alteacare/golang-basecode/src/interfaces/http/controllers/v1"
)

func (i *V1Routes) MountPing() {
	g := i.Echo.Group("/ping")
	pingController := v1controller.New(&v1controller.V1{
		DB:         i.DB,
		Cloudwatch: i.Cloudwatch,
	})

	g.GET("", pingController.Ping)
}
