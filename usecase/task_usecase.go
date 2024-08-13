package usecase

import (
	"context"
	"go-ecommerce/model"
)

type TaskUsecase interface {
	Create(c context.Context, task *model.Task) error
	FetchByUserID(c context.Context, userID string) ([]model.Task, error)
}
