package repository

import (
	models "testProject/learning/example/api_jwt_mongo/model"
)

type UserRepo interface {
	FindAll(args ...interface{}) ([]models.User, error)
	FindOne(queryData map[string]interface{}) (models.User, error)
	Insert(u *models.User) (string, error)
	CheckLogin(email, password string) (models.User, error)
}
