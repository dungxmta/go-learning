package repoimpl

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	driverMongo "testProject/learning/example/api_jwt_mongo/driver/mongo"
	models "testProject/learning/example/api_jwt_mongo/model"
	testUtils "testProject/learning/example/api_jwt_mongo/tests"
	"testing"
)

func TestUserRepoImpl_Insert(t *testing.T) {
	// setup db
	cs := testUtils.ConnString(t)
	connector := driverMongo.ConnectMongoDB(cs.Original)
	db := connector.Client.Database(testUtils.GetDBName(cs))

	assert.Equal(t, "mongodb://localhost:27018/go_learning_test", cs.Original)

	defer func() {
		// _ = db.Drop(context.Background())
		_ = db.Collection(colName).Drop(context.Background())
		_ = connector.Client.Disconnect(context.Background())
	}()

	userRepo := NewUserRepo(db)

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
		id, err := userRepo.Insert(&u)
		if err != nil {
			t.Error("Err when insert user", idx, " | ", err)
			assert.Equal(t, "", id)
			continue
		}
		assert.Equal(t, u.ID, id)
	}
}
