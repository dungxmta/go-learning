package driver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

type MongoConnector struct {
	Client *mongo.Client
}

// https://medium.com/golang-issue/how-singleton-pattern-works-with-golang-2fdd61cd5a7f
// var singletonMongo = &MongoDB{}
var singletonMongo *MongoConnector
var once sync.Once

func GetInstance() *MongoConnector {
	// if singletonMongo == nil {
	once.Do(func() {
		singletonMongo = &MongoConnector{}
	})
	return singletonMongo
}

func ConnectMongoDB(uri string) *MongoConnector {
	if singletonMongo == nil {
		GetInstance()
	}
	// new client
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// set timeout when connect
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	// ping test
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("connect db ok!")

	singletonMongo.Client = client
	return singletonMongo
}
