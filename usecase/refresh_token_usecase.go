package usecase

import (
	"context"
	"go-ecommerce/model"
)

type RefreshTokenUsecase interface {
	GetUserByID(c context.Context, id string) (model.User, error)
	CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
}
