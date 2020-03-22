package driver

type Storage interface {
	Init(string) error
	SetDB(string)

	Find()
}

type MsgQueue interface {
	Init(string, string, int) (MsgQueue, error)

	LPush(string, ...interface{}) (int64, error)
	RPop(string) (string, error)
	HGet(string, string) (string, error)
}
