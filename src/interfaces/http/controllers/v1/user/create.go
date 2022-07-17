package userv1controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	userRepository "alteacare/golang-basecode/src/repositories/user"
	userUsecase "alteacare/golang-basecode/src/usecases/user"
)

func (i *V1User) CreateUser(c echo.Context) (err error) {
	u := new(createRequest)

	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(u); err != nil {
		return err
	}

	ur := userRepository.New(i.DB)
	uu := userUsecase.New(ur)

	data, err := uu.CreateUser(&userUsecase.ParamsCreateUser{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})

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
	createRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)
