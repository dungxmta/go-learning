package main

import (
	"fmt"
	"testProject/learning/example/db_postgres_with_repository_pattern/driver"
	models "testProject/learning/example/db_postgres_with_repository_pattern/model"
	repo "testProject/learning/example/db_postgres_with_repository_pattern/repository"
	repo_impl "testProject/learning/example/db_postgres_with_repository_pattern/repository/repo_implement"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "admin"
	password = "123456"
	dbname   = "database"
)

func main() {

	db := driver.Connect(host, port, user, password, dbname)

	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connect db ok!")

	userRepo := repo_impl.NewUserRepo(db.SQL)

	// addUsers(userRepo)
	getUsers(userRepo)
}

func getUsers(userRepo repo.UserRepo) {
	users, _ := userRepo.Select()

	for _, u := range users {
		fmt.Println(u)
	}
}

func addUsers(userRepo repo.UserRepo) {
	dataUsers := []models.User{
		models.User{
			ID:     0,
			Name:   "user1",
			Gender: "M",
			Email:  "user1@gmail.com",
		},
		models.User{
			ID:     1,
			Name:   "user2",
			Gender: "F",
			Email:  "user2@gmail.com",
		},
	}

	for idx, u := range dataUsers {
		err := userRepo.Insert(u)
		if err != nil {
			fmt.Println("Err when insert user", idx, " | ", err)
		}
	}
}
