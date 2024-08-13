package usecase

import (
	"context"
	"go-ecommerce/dto"
	"go-ecommerce/model"
	"go-ecommerce/usecase"
	"time"
)

type profileUsecase struct {
	userRepository model.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository model.UserRepository, timeout time.Duration) usecase.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*dto.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &dto.Profile{Name: user.Name, Email: user.Email}, nil
}
