package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testProject/learning/example/api_jwt_mongo/config"
	"testProject/learning/example/api_jwt_mongo/driver"
	models "testProject/learning/example/api_jwt_mongo/model"
	repo "testProject/learning/example/api_jwt_mongo/repository"
	repoImpl "testProject/learning/example/api_jwt_mongo/repository/repoimpl"
)

func main() {

	mongoUri := config.GetStr("MONGO_URI")
	mongoDBName := config.GetStr("MONGO_DBNAME")

	fmt.Println(mongoUri)

	mongo := driver.ConnectMongoDB(mongoUri)

	userRepo := repoImpl.NewUserRepo(mongo.Client.Database(mongoDBName))

	// addUser(userRepo)
	// findUser(userRepo)
	userLogin(userRepo)
}

func userLogin(userRepo repo.UserRepo) {

	user, err := userRepo.CheckLogin("admin@gmail.com", "1")
	if err != nil {
		fmt.Println("Email or Password wrong!")
		return
	}

	fmt.Println(user)
}

func findUser(userRepo repo.UserRepo) {
	queryData := map[string]interface{}{
		"email": "admin@gmail.com",
	}

	user, err := userRepo.FindOne(queryData)
	if err != nil {
		fmt.Println("User not found!", err)
		return
	}

	fmt.Println(user)
}

func addUser(userRepo repo.UserRepo) {
	dataUsers := []models.User{
		models.User{
			ID:       primitive.NewObjectID().Hex(),
			Name:     "admin",
			Email:    "admin@gmail.com",
			Role:     "admin",
			Password: "1",
		},
		models.User{
			ID:       primitive.NewObjectID().Hex(),
			Name:     "user1",
			Email:    "user1@gmail.com",
			Role:     "user",
			Password: "1",
		},
	}

	for idx, u := range dataUsers {
		_, err := userRepo.Insert(&u)
		if err != nil {
			fmt.Println("Err when insert user", idx, " | ", err)
		}
	}
}
