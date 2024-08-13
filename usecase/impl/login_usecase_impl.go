package usecase

import (
	"context"
	"go-ecommerce/model"
	"go-ecommerce/usecase"
	"go-ecommerce/utils/tokenutil"
	"time"
)

type loginUsecase struct {
	userRepository model.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository model.UserRepository, timeout time.Duration) usecase.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)

}

func (lu *loginUsecase) CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
