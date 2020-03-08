package main

import (
	"fmt"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"testProject/learning/example/api_jwt_mongo/config"
	driverMongo "testProject/learning/example/api_jwt_mongo/driver/mongo"
	// driver "testProject/learning/example/api_jwt_mongo/driver"
	// models "testProject/learning/example/api_jwt_mongo/model"
	// repo "testProject/learning/example/api_jwt_mongo/repository"
	apiHandler "testProject/learning/example/api_jwt_mongo/handler"
	repoImpl "testProject/learning/example/api_jwt_mongo/repository/repoimpl"
)

func main() {

	mongoUri := config.GetStr("MONGO_URI")
	mongoDBName := config.GetStr("MONGO_DBNAME")

	fmt.Println(mongoUri)

	mongo := driverMongo.ConnectMongoDB(mongoUri)

	userRepo := repoImpl.NewUserRepo(mongo.Client.Database(mongoDBName))

	// apiHandler.AddUser(userRepo)
	// apiHandler.FindUser(userRepo)
	apiHandler.UserLogin(userRepo)
}
