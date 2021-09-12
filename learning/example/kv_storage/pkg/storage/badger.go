package storage

import (
	"time"

	"github.com/dgraph-io/badger/v3"
	"go.uber.org/zap"

	"testProject/pkg/utils"
)

type BadgerDB struct {
	db     *badger.DB
	logger *zap.Logger
}

func (b *BadgerDB) Init(dir string, opts *Opts) (err error) {
	if err = utils.InitializeDir(dir); err != nil {
		return
	}
	// Create the options
	cfg := badger.DefaultOptions(dir)
	// opts.SyncWrites = false
	// opts.ValueLogLoadingMode = options.FileIO
	// opts.IndexCacheSize = 512 << 20 // 512 MB
	cfg.InMemory = false

	// Attempt to open the database
	db, err := badger.Open(cfg)
	if err != nil {
		return err
	}

	b.db = db

	// client need to manual trigger GC
	if opts.AutoMonitor {
		go b.Monitor()
	}
	return
}

func (b *BadgerDB) Exists(k string) (ok bool, err error) {
	err = b.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(k))
		// if err != nil {
		return err
		// }

		// var valNot, valCopy []byte
		// err := item.Value(func(val []byte) error {
		// 	// This func with val would only be called if item.Value encounters no error.
		//
		// 	// Accessing val here is valid.
		// 	fmt.Printf("The answer is: %s\n", val)
		//
		// 	// Copying or parsing val is valid.
		// 	valCopy = append([]byte{}, val...)
		//
		// 	// Assigning val slice to another variable is NOT OK.
		// 	valNot = val // Do not do this.
		// 	return nil
		// })
		// handle(err)
		//
		// // You must copy it to use it outside item.Value(...).
		// fmt.Printf("The answer is: %s\n", valCopy)
		//
		// // Alternatively, you could also use item.ValueCopy().
		// valCopy, err = item.ValueCopy(nil)
		// handle(err)
		// fmt.Printf("The answer is: %s\n", valCopy)
		// return nil
	})
	if err == badger.ErrKeyNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (b *BadgerDB) SetNull(k string) (err error) {
	err = b.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(k), nil)
		return err
	})
	return
}

func (b *BadgerDB) Set(k string, v interface{}) (err error) {
	// TODO
	return
}

func (b *BadgerDB) Close() (err error) {
	err = b.db.Close()
	return
}

// https://dgraph.io/docs/badger/get-started/#garbage-collection
func (b *BadgerDB) Monitor() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
	again:
		err := b.db.RunValueLogGC(0.7)
		if err == nil {
			goto again
		}
	}
}
