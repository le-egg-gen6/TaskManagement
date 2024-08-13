package usecase

import (
	"context"
	"go-ecommerce/dto"
)

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*dto.Profile, error)
}
