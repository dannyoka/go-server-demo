package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017/mern-shopping"

func InitDB() (*mongo.Client, error){
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)	
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts); if err != nil{
		panic(err)
	}
	fmt.Println("successfully connected to mongodb")

	return client, nil
}