package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

var (
	MaxWorker int
	MaxQueue  int
)

func init() {
	var err error
	MaxWorker, err = strconv.Atoi(os.Getenv("MAX_WORKERS"))
	if err != nil {
		MaxWorker = 10
	}

	MaxQueue, err = strconv.Atoi(os.Getenv("MAX_QUEUE"))
	if err != nil {
		MaxQueue = 2
	}
}

// Job represents the job to be run
type Job interface {
	Info() string
	Do() error
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

// --------------------------------------------------------
// Implement Job interface
// --------------------------------------------------------
type Sleeper struct {
	Id   string
	Data string
}

func (obj *Sleeper) Info() string {
	return "I'm just print data to console and take a nap !!! zzz ..."
}

func (obj *Sleeper) Do() error {
	log.Printf("%v -> %v", obj.Id, obj.Data)
	time.Sleep(time.Second * 5)
	return nil
}
