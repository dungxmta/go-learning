package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"testProject/learning/example/api_jwt_mongo/repository"
	"testProject/learning/example/api_jwt_mongo/repository/repoimpl"
	"time"
)

type Storage interface {
	Init(uri string)
	SetDB(name string)
	GetUserRI() repository.UserRepo
}

type MongoConnector struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// https://medium.com/golang-issue/how-singleton-pattern-works-with-golang-2fdd61cd5a7f
// var singletonMongo = &MongoDB{}
var singletonMongo Storage
var once sync.Once

// export get & set
func GetInstance() Storage {
	once.Do(func() {
		singletonMongo = &MongoConnector{}
	})
	return singletonMongo
}

func SetInstance(ins Storage) {
	singletonMongo = ins
}

// TODO: remove me
func ConnectMongoDB(uri string) *MongoConnector {
	// if singletonMongo == nil {
	// 	GetInstance()
	// }
	// // new client
	// client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // set timeout when connect
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	// err = client.Connect(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // ping test
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println("connect db ok!")
	//
	// singletonMongo.Client = client
	// return singletonMongo
	return nil
}

func (ins *MongoConnector) Init(uri string) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	singletonMongo.(*MongoConnector).Client, _ = mongo.Connect(ctx, options.Client().ApplyURI(uri))
}

func (ins *MongoConnector) SetDB(name string) {
	ins.DB = ins.Client.Database(name)
}

// get repo implement
func (ins *MongoConnector) GetUserRI() repository.UserRepo {
	return repoimpl.NewUserRepo(ins.DB)
}
