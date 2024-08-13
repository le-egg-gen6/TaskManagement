package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoModel struct {
	ID primitive.ObjectID `bson:"id" json:"-"`
}
