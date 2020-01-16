package service

import(
	"context"
	"github.com/sukstar76/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserServiceInterface interface{
	GetID(*context.Context, string)(*model.User,error)
	Create(*context.Context, *model.User) (*model.User,error)
	Update(*context.Context, *model.User)(*model.User,error)

}

var DBCollection *mongo.Collection


func GetID(ctx *context.Context,id string) (*model.User,error){
	var user model.User
	if err:= DBCollection.FindOne(*ctx,bson.D{{"id",id}}).Decode(&user); err !=nil{
		return nil,err
	}
	return &user,nil
}

func Create(ctx *context.Context,u *model.User) (*model.User,error){
	_,err:= DBCollection.InsertOne(*ctx,u)
	if err!= nil{
		return nil,err
	}
	return u,nil

}

func Update(ctx *context.Context,u *model.User) (*model.User,error){
	result,err:= DBCollection.ReplaceOne(*ctx, bson.D{{"id",u.ID}},u)
	if err!=nil || result ==nil{
		return nil,err
	}
	return u, nil

}

