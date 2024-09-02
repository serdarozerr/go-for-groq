package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type QueryAnswer struct {
	Query  string
	Answer string
}

type User struct {
	Name         string        `json:"name" bson:"name"`
	Surname      string        `json:"surname" bson:"surname"`
	QueryAnswers []QueryAnswer `json:"query_answers" bson:"query_answers"`
}

type Response struct {
	Id mongo.InsertOneResult `json:"id" bson:"_id"`
}

func (s User) Create(collection *mongo.Collection) (*mongo.InsertOneResult, error) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	result, err := collection.InsertOne(ctx, s)
	if err != nil {
		return nil, err
	}
	return result, nil

}
