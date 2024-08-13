package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "tasks"
)

type Task struct {
	MongoModel
	Title  string             `bson:"title" json:"title"`
	UserID primitive.ObjectID `bson:"userID" json:"-"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserId(c context.Context, userId string) ([]Task, error)
}
