package usecase

import (
	"context"
	"go-ecommerce/model"
	"go-ecommerce/usecase"
	"time"
)

type taskUsecase struct {
	taskRepository model.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository model.TaskRepository, timeout time.Duration) usecase.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tu *taskUsecase) Create(c context.Context, task *model.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)

	defer cancel()

	return tu.taskRepository.Create(ctx, task)
}

func (tu *taskUsecase) FetchByUserID(c context.Context, userID string) ([]model.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)

	defer cancel()

	return tu.taskRepository.FetchByUserId(ctx, userID)
}
