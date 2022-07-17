package userusecase

import (
	"os"

	"alteacare/golang-basecode/src/helpers"
)

type (
	ParamsLogin struct {
		Email    string
		Password string
	}
	GeneratedToken struct {
		Token     string `json:"token"`
		ExpiredAt int64  `json:"expired_at"`
	}
	ResultLogin struct {
		AccessToken  GeneratedToken
		RefreshToken GeneratedToken
	}
)

func (i *sUserUsecase) Login(p *ParamsLogin) (*ResultLogin, error) {
	user, _ := i.userRepository.FindByEmail(&p.Email)

	if user == nil {
		return nil, ErrUserNotFound
	}

	isValidPassword := helpers.CheckPasswordHash(p.Password, user.Password)
	if !isValidPassword {
		return nil, ErrInvalidPassword
	}

	accessToken, expiredAtAccessToken, errAccessToken := helpers.GenerateJWT(&helpers.ParamsGenerateJWT{
		ExpiredInMinute: 2,
		UserId:          user.ID,
		SecretKey:       os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
	})

	refreshToken, expiredAtRefreshToken, errRefreshToken := helpers.GenerateJWT(&helpers.ParamsGenerateJWT{
		ExpiredInMinute: 2,
		UserId:          user.ID,
		SecretKey:       os.Getenv("REFRESH_TOKEN_SECRET_KEY"),
	})

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	if errRefreshToken != nil {
		return nil, errRefreshToken
	}

	return &ResultLogin{
		AccessToken: GeneratedToken{
			Token:     accessToken,
			ExpiredAt: expiredAtAccessToken,
		},
		RefreshToken: GeneratedToken{
			Token:     refreshToken,
			ExpiredAt: expiredAtRefreshToken,
		},
	}, nil
}
