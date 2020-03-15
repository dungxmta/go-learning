package tests

import (
	"github.com/stretchr/testify/assert"
	// "go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "go.mongodb.org/mongo-driver/mongo/writeconcern"
	"context"
	"testing"
)

// var (
// 	connsCheckedOut int
// )

/**
func TestUserRepoImpl_Insert(t *testing.T) {
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
	...
}
*/

func TestDemo(t *testing.T) {
	cs := ConnString(t)
	// poolMonitor := &event.PoolMonitor{
	// 	Event: func(evt *event.PoolEvent) {
	// 		switch evt.Type {
	// 		case event.GetSucceeded:
	// 			connsCheckedOut++
	// 		case event.ConnectionReturned:
	// 			connsCheckedOut--
	// 		}
	// 	},
	// }
	clientOpts := options.Client().ApplyURI(cs.Original).SetReadPreference(readpref.Primary())
	// SetWriteConcern(writeconcern.New(writeconcern.WMajority())).SetPoolMonitor(poolMonitor)
	client, err := mongo.Connect(context.Background(), clientOpts)
	assert.Nil(t, err, "Connect error: %v", err)
	db := client.Database("gridfs")
	defer func() {
		// sessions := client.NumberSessionsInProgress()
		// conns := connsCheckedOut

		_ = db.Drop(context.Background())
		_ = client.Disconnect(context.Background())
		// assert.Equal(t, 0, sessions, "%v sessions checked out", sessions)
		// assert.Equal(t, 0, conns, "%v connections checked out", conns)
	}()
}
