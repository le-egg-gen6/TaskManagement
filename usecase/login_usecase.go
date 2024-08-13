package usecase

import (
	"context"
	"go-ecommerce/model"
)

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (model.User, error)
	CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error)
}
