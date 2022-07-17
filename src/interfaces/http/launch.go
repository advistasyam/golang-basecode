package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"

	"alteacare/golang-basecode/src/helpers"
	v1routes "alteacare/golang-basecode/src/interfaces/http/routes/v1"
	v2routes "alteacare/golang-basecode/src/interfaces/http/routes/v2"
)

func (i *Http) Launch() {
	e := echo.New()
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = helpers.ErrorHandler
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(helpers.MiddlewareLogger(i.Cloudwatch.Session))

	e.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf("API Base Code for %s", os.Getenv("ENVIRONMENT")))
	})

	v1 := v1routes.New(&v1routes.V1Routes{
		Echo:       e.Group("/v1"),
		DB:         i.DB,
		Cloudwatch: i.Cloudwatch,
	})
	v2 := v2routes.New(&v2routes.V2Routes{
		Echo:       e.Group("/v1"),
		DB:         i.DB,
		Cloudwatch: i.Cloudwatch,
	})
	v1.MountPing()
	v1.MountUser()
	v2.MountPing()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))))
}
