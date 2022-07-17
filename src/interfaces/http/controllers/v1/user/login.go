package userv1controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	userRepository "alteacare/golang-basecode/src/repositories/user"
	userUsecase "alteacare/golang-basecode/src/usecases/user"
)

func (i *V1User) LoginUser(c echo.Context) (err error) {
	u := new(loginRequest)

	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(u); err != nil {
		return err
	}

	ur := userRepository.New(i.DB)
	uu := userUsecase.New(ur)

	data, err := uu.Login(&userUsecase.ParamsLogin{
		Email:    u.Email,
		Password: u.Password,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Status:  true,
		Message: "OK",
		Data: &loginResponse{
			AccessToken:  data.AccessToken,
			RefreshToken: data.RefreshToken,
		},
	})
}

type (
	loginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	loginResponse struct {
		AccessToken  interface{} `json:"access_token"`
		RefreshToken interface{} `json:"refresh_token"`
	}
)
