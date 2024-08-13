package repository

import (
	"context"
	"go-ecommerce/model"
	"go-ecommerce/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(database mongo.Database, collection string) model.TaskRepository {
	return &taskRepository{
		database:   database,
		collection: collection,
	}
}

func (tr *taskRepository) Create(c context.Context, task *model.Task) error {
	collection := tr.database.Collection(tr.collection)

	_, err := collection.InsertOne(c, task)

	return err
}

func (t *taskRepository) FetchByUserId(c context.Context, userId string) ([]model.Task, error) {
	collection := t.database.Collection(t.collection)

	var tasks []model.Task

	idHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return tasks, err
	}

	cursor, err := collection.Find(c, bson.M{"userID": idHex})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &tasks)
	if tasks == nil {
		return []model.Task{}, err
	}

	return tasks, err
}
