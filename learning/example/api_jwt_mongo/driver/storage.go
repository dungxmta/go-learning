package driver

import "context"

type Storage interface {
	Init(string) (Storage, error)
	SetDB(string)

	Find(colName string, ctx context.Context, results interface{}, filter interface{}, opts ...*interface{}) error
	FindOne(colName string, ctx context.Context, result interface{}, filter interface{}, opts ...*interface{}) error
}

type MsgQueue interface {
	Init(string, string, int) (MsgQueue, error)

	LPush(string, ...interface{}) (int64, error)
	RPop(string) (string, error)
	HGet(string, string) (string, error)
}
