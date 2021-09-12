package storage

const (
	Badger int = iota
)

type Opts struct {
	// badger
	AutoMonitor bool
	ReadOnly    bool
}

type KvStorage interface {
	Init(dir string, opts *Opts) error

	Exists(k string) (bool, error)
	SetNull(k string) error
	Set(k string, v interface{}) error

	Monitor()
	Close() error
}

func NewKvStorage(sType int) KvStorage {
	switch sType {
	case Badger:
		return &BadgerDB{}
	default:
		return nil
	}
}
