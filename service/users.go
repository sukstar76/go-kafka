package service

import (
	"context"
	"github.com/sukstar76/go-kafka/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceInterface interface {
	GetID(*context.Context, string) (*model.User, error)
	Create(*context.Context, *model.User) (*model.User, error)
	Update(*context.Context, *model.User) (*model.User, error)
}

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(client *mongo.Client) *UserService {
	return &UserService{
		collection: client.Database("test1").Collection("user"),
	}
}

func (us *UserService) GetID(ctx *context.Context, id string) (*model.User, error) {
	var user model.User
	if err := us.collection.FindOne(*ctx, bson.D{{"id", id}}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserService) Create(ctx *context.Context, u *model.User) (*model.User, error) {
	_, err := us.collection.InsertOne(*ctx, u)
	if err != nil {
		return nil, err
	}
	return u, nil

}

func (us *UserService) Update(ctx *context.Context, u *model.User) (*model.User, error) {
	result, err := us.collection.ReplaceOne(*ctx, bson.D{{"id", u.ID}}, u)
	if err != nil || result == nil {
		return nil, err
	}
	return u, nil

}
