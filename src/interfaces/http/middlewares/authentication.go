package middlewares

import (
	"errors"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"

	"alteacare/golang-basecode/src/helpers"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)

		if token == "" {
			return errors.New("Token not found")
		}

		claims, err := helpers.ValidateJWT(&helpers.ParamsValidateJWT{
			Token:     token,
			SecretKey: os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
		})

		if err != nil {
			return err
		}

		user := make(map[string]interface{})
		mapstructure.Decode(claims, &user)

		c.Set("user", user)

		return next(c)
	}
}
