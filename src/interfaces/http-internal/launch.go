package http_internal

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"

	"alteacare/golang-basecode/src/helpers"
)

func (i *HttpInternal) Launch() {
	e := echo.New()
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = helpers.ErrorHandler
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(helpers.MiddlewareLogger(i.Cloudwatch.Session))

	e.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf("Internal API Base Code for %s", os.Getenv("ENVIRONMENT")))
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("HTTP_INTERNAL_PORT"))))
}
