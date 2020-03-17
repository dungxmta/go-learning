package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	driverMongo "testProject/learning/example/api_jwt_mongo/driver/mongo"
	// "testProject/learning/example/api_jwt_mongo/extensions"

	apiHandler "testProject/learning/example/api_jwt_mongo/handler"
	models "testProject/learning/example/api_jwt_mongo/model"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"testProject/learning/example/api_jwt_mongo/config"
	// repoImpl "testProject/learning/example/api_jwt_mongo/repository/repoimpl"
)

func main() {

	mongoUri := config.GetStr("MONGO_URI")
	mongoDBName := config.GetStr("MONGO_DBNAME")

	fmt.Println(mongoUri)

	driverMongo.GetInstance().Init(mongoUri)
	driverMongo.GetInstance().SetDB(mongoDBName)

	// userRepo := driverMongo.GetInstance().GetUserRI()
	// userRepo := extensions.UserRepo
	// ext := extensions.NewExt()
	// ext := extensions.GetInstance()
	// // userRepo := ext.UserRepo
	// userRepo := ext.UserRepo
	// fmt.Println(userRepo)
	// err := apiHandler.Demo(userRepo)
	err := apiHandler.Demo()
	fmt.Println(err)
	/*
		mongo := driverMongo.ConnectMongoDB(mongoUri)

		userRepo := repoImpl.NewUserRepo(mongo.Client.Database(mongoDBName))
		comRepo := repoImpl.NewCommonRepo(mongo.Client.Database(mongoDBName))

		// apiHandler.AddUser(userRepo)
		// apiHandler.FindUser(userRepo)
		apiHandler.UserLogin(userRepo)

		queryData := map[string]interface{}{
			"email": "admin@gmail.com",
		}

		user, err := comRepo.FindOne("users", queryData)
		if err != nil {
			fmt.Println("User not found!", err)
			return
		}

		// not work
		f, ok := user.(*models.User)
		fmt.Println(f, ok)

		// convert map to json/bson first then decode back to struct
		jsonData, err := bson.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}
		var u models.User
		err = bson.Unmarshal(jsonData, &u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u)
		fmt.Println(u.Role)

		users, err := comRepo.FindAll("users", map[string]interface{}{})
		if err != nil {
			fmt.Println("User not found!", err)
			return
		}

		fmt.Println(users)
		for _, d := range users {
			fmt.Println(map2user(d))
		}
	*/
}

func map2user(mapData interface{}) models.User {
	bsonData, err := bson.Marshal(mapData)
	if err != nil {
		log.Fatal(err)
	}
	var u models.User
	err = bson.Unmarshal(bsonData, &u)
	if err != nil {
		log.Fatal(err)
	}
	return u
}
