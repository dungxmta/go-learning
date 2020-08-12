package driver

import (
	"context"
	"github.com/go-redsync/redsync/v3"
	"time"
)

type Storage interface {
	Init(string) (Storage, error)
	SetDB(string)

	Find(colName string, ctx context.Context, results interface{}, filter interface{}, opts ...*interface{}) error
	FindOne(colName string, ctx context.Context, result interface{}, filter interface{}, opts ...*interface{}) error
}

type MsgQueue interface {
	Init(addr, password string, db int) (MsgQueue, error)
	InitLocker()

	LPush(string, ...interface{}) (int64, error)
	RPop(string) (string, error)
	HGet(string, string) (string, error)

	TTL(key string) (time.Duration, error)

	NewMutex(name string, options ...redsync.Option) Locker
}

type Locker interface {
	Lock() error
	Unlock() (bool, error)
	Extend() (bool, error)
	Valid() (bool, error)
}
