package v1routes

import (
	userv1controller "alteacare/golang-basecode/src/interfaces/http/controllers/v1/user"
	"alteacare/golang-basecode/src/interfaces/http/middlewares"
)

func (i *V1Routes) MountUser() {
	g := i.Echo.Group("/user")

	userController := userv1controller.New(&userv1controller.V1User{
		DB:         i.DB,
		Cloudwatch: i.Cloudwatch,
	})

	g.POST("", userController.CreateUser)
	g.POST("/login", userController.LoginUser)
	g.GET("/me", userController.MeUser, middlewares.Authentication)
	g.GET("/:id", userController.DetailUser)
}
