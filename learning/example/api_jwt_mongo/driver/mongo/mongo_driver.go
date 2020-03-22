package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"testProject/learning/example/api_jwt_mongo/driver"
	"time"
)

type connector struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// https://medium.com/golang-issue/how-singleton-pattern-works-with-golang-2fdd61cd5a7f
var singleton driver.Storage
var once sync.Once

// export get & set
func GetInstance() driver.Storage {
	once.Do(func() {
		singleton = &connector{}
	})
	return singleton
}

func SetInstance(ins driver.Storage) {
	singleton = ins
}

func (ins *connector) Init(uri string) (driver.Storage, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// ping test
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	ins.Client = client
	return ins, nil
}

func (ins *connector) SetDB(name string) {
	ins.DB = ins.Client.Database(name)
}

func NewClient(uri string) (driver.Storage, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// set timeout when connect
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// ping test
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &connector{
		Client: client,
		DB:     nil,
	}, nil
}
