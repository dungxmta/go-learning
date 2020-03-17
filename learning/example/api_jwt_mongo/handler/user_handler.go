package handler

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	driverMongo "testProject/learning/example/api_jwt_mongo/driver/mongo"
	models "testProject/learning/example/api_jwt_mongo/model"
	repo "testProject/learning/example/api_jwt_mongo/repository"
	"testProject/pkg/utils"
)

/**
NOTE: business logic here! not in "repo implement"
*/

func UserLogin(userRepo repo.UserRepo) {

	user, err := userRepo.CheckLogin("admin@gmail.com", "1")
	if err != nil {
		fmt.Println("Email or Password wrong!")
		return
	}

	fmt.Println(user)
}

func FindUser(userRepo repo.UserRepo) {
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

func AddUser(userRepo repo.UserRepo) error {
	dataUsers := []models.User{
		{
			ID:       primitive.NewObjectID().Hex(),
			Name:     "admin",
			Email:    "admin@gmail.com",
			Role:     "admin",
			Password: "1",
		},
		{
			ID:       primitive.NewObjectID().Hex(),
			Name:     "user1",
			Email:    "user1@gmail.com",
			Role:     "user",
			Password: "1",
		},
	}

	for idx, u := range dataUsers {
		// check duplicate
		// var query = make(map[string]interface{})
		queryExists := map[string]interface{}{
			"email": u.Email,
		}
		_, found := userRepo.FindOne(queryExists)
		if found == nil {
			return errors.New(fmt.Sprintf("user '%v' existed!", u.Name))
		}

		// hash password
		hashPwd, err := utils.HashedPwd(u.Password)
		if err != nil {
			return (err)
		}
		u.Password = hashPwd

		_, err = userRepo.Insert(&u)
		if err != nil {
			log.Fatal("Err when insert user", idx, " | ", err)
			return err
		}
	}
	return nil
}

// func Demo(userRepo repo.UserRepo) error {
func Demo() error {
	// userRepo := extensions.GetInstance().UserRepo
	userRepo := driverMongo.GetInstance().GetUserRI()
	fmt.Println(userRepo)

	dataUsers := []models.User{
		{
			ID:       primitive.NewObjectID().Hex(),
			Name:     "admin",
			Email:    "admin@gmail.com",
			Role:     "admin",
			Password: "1",
		},
		{
			ID:       primitive.NewObjectID().Hex(),
			Name:     "user1",
			Email:    "user1@gmail.com",
			Role:     "user",
			Password: "1",
		},
	}

	// u, _ := userRepo.FindAll(map[string]interface{}{})
	// fmt.Println(u)

	for idx, u := range dataUsers {
		// check duplicate
		// var query = make(map[string]interface{})
		queryExists := map[string]interface{}{
			"email": u.Email,
		}
		_, found := userRepo.FindOne(queryExists)
		if found == nil {
			return errors.New(fmt.Sprintf("user '%v' existed!", u.Name))
		}

		// hash password
		hashPwd, err := utils.HashedPwd(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashPwd

		_, err = userRepo.Insert(&u)
		if err != nil {
			log.Fatal("Err when insert user", idx, " | ", err)
			return err
		}
	}

	return nil
}
