package repoimpl

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	models "testProject/learning/example/api_jwt_mongo/model"
	repo "testProject/learning/example/api_jwt_mongo/repository"
)

const colName = "users"

type UserRepoImpl struct {
	Db *mongo.Database
}

func NewUserRepo(db *mongo.Database) repo.UserRepo {
	return &UserRepoImpl{
		Db: db,
	}
}

func (mongoWrap *UserRepoImpl) FindAll(args ...interface{}) ([]models.User, error) {
	users := []models.User{}

	return users, nil
}

func (mongoWrap *UserRepoImpl) FindOne(queryData map[string]interface{}) (models.User, error) {
	query := bson.M{}
	for k, v := range queryData {
		query[k] = v
	}
	result := mongoWrap.Db.Collection(colName).
		FindOne(context.Background(), query)
	// FindOne(context.Background(), bson.M{"email": u.Email})

	user := models.User{}
	err := result.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (mongoWrap *UserRepoImpl) Insert(u *models.User) (string, error) {
	// check duplicate
	// var query = make(map[string]interface{})
	queryExists := map[string]interface{}{
		"email": u.Email,
	}
	_, found := mongoWrap.FindOne(queryExists)
	if found == nil {
		return "", errors.New("user existed")
	}
	// TODO: encrypt pwd
	// encode data
	bbytes, _ := bson.Marshal(u)

	result, err := mongoWrap.Db.Collection(colName).InsertOne(context.Background(), bbytes)
	if err != nil {
		return "", err
	}

	// TODO: return _id
	fmt.Println("Inserted user ", u.ID, u.Email, result)
	return "", nil
}

// check email + password
func (mongoWrap *UserRepoImpl) CheckLogin(email, password string) (models.User, error) {
	queryExists := map[string]interface{}{
		"email":    email,
		"password": password,
	}
	user, err := mongoWrap.FindOne(queryExists)
	if err != nil {
		return user, err
	}
	// TODO: compare pwd

	return user, nil
}
