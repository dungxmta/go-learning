package repoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	repo "testProject/learning/example/api_jwt_mongo/repository"
)

type CommonRepoImpl struct {
	DB *mongo.Database
}

func NewCommonRepo(db *mongo.Database) repo.CommonRepo {
	return &CommonRepoImpl{
		DB: db,
	}
}

func (ins *CommonRepoImpl) FindAll(colName string, queryData map[string]interface{}) ([]interface{}, error) {
	var results []interface{}

	query := bson.M{}
	for k, v := range queryData {
		query[k] = v
	}

	cur, err := ins.DB.Collection(colName).Find(context.Background(), query)
	if err != nil {
		return results, err
	}

	for cur.Next(context.Background()) {
		var obj map[string]interface{}
		err := cur.Decode(&obj)
		if err != nil {
			return results, err
		}
		results = append(results, obj)
	}

	if err := cur.Err(); err != nil {
		return results, err
	}

	return results, nil
}

func (ins *CommonRepoImpl) FindOne(colName string, queryData map[string]interface{}) (interface{}, error) {
	query := bson.M{}
	for k, v := range queryData {
		query[k] = v
	}
	result := ins.DB.Collection(colName).FindOne(context.Background(), query)

	var obj map[string]interface{}
	err := result.Decode(&obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
