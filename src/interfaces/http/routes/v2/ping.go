package v2routes

import (
	v2controller "alteacare/golang-basecode/src/interfaces/http/controllers/v2"
)

func (i *V2Routes) MountPing() {
	g := i.Echo.Group("/ping")
	pingController := v2controller.New(&v2controller.V2{
		DB:         i.DB,
		Cloudwatch: i.Cloudwatch,
	})

	g.GET("", pingController.Ping)
}
