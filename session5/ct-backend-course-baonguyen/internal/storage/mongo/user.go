package mongostore

import (
	"context"
	"ct-backend-course-baonguyen/internal/entity"
	"errors"
	"fmt"

	_ "net/url"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewUserCollection(uri string, dbName, collName string) *userCollection {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	return &userCollection{
		client:  client.Database(dbName).Collection(collName),
		timeout: 3 * time.Second,
	}
}

type userCollection struct {
	client  *mongo.Collection
	timeout time.Duration
}

func (c *userCollection) Save(info entity.UserInfo) error {
	ctx, cancelFn := context.WithTimeout(context.Background(), c.timeout)
	defer cancelFn()
	db := c.client.Database()
	
	condition := bson.M{
		"username": info.Username,
	}
	var userInfo entity.UserInfo
    e := db.Collection("user").FindOne(ctx, condition).Decode(&userInfo)

	if e == nil {
		return errors.New("User existed!")
	}
	
	res, err := db.Collection("user").InsertOne(context.Background(), info)
	if err != nil {
		return errors.New("Cannot create new user")
	}
	fmt.Print(res)
	return nil
}

func (c *userCollection) Get(username string) (entity.UserInfo, error) {
	// panic("TODO implement it")

	ctx, cancelFn := context.WithTimeout(context.Background(), c.timeout)
	defer cancelFn()
	//
	//panic("TODO implement it")
	//
	//return  entity.UserInfo{}, nil
	db := c.client.Database()
	
	condition := bson.M{
		"username": username,
	}
	var userInfo entity.UserInfo
    e := db.Collection("user").FindOne(ctx, condition).Decode(&userInfo)
	if e != nil {
		return entity.UserInfo{}, errors.New("User not existed!")
	}
	return userInfo, nil
}