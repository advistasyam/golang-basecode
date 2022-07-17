package userv1controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"

	userRepository "alteacare/golang-basecode/src/repositories/user"
	userUsecase "alteacare/golang-basecode/src/usecases/user"
)

func (i *V1User) MeUser(c echo.Context) (err error) {
	u := new(meValidator)
	mapstructure.Decode(c.Get("user"), &u)

	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(u); err != nil {
		return err
	}

	ur := userRepository.New(i.DB)
	uu := userUsecase.New(ur)

	data, err := uu.DetailUser(&u.ID)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Status:  true,
		Message: "OK",
		Data:    data,
	})
}

type (
	// Request Validator
	meValidator struct {
		ID uint `mapstructure:"user_id" validate:"required"`
	}
)
