package main

import (
	"fmt"
	"time"
)

type Job interface {
	Process()
}

type Worker struct {
	WorkerId   int
	Done       chan bool
	JobRunning chan Job
}

func NewWorker(workerID int, jobChan chan Job) *Worker {
	return &Worker{
		WorkerId:   workerID,
		Done:       make(chan bool),
		JobRunning: jobChan,
	}
}

func (w *Worker) Run() {
	fmt.Println("Run worker id ", w.WorkerId)

	go func() {
		for {
			select {
			case job := <-w.JobRunning:
				fmt.Println("Job running ", w.WorkerId)
				job.Process()
			case <-w.Done:
				fmt.Println("Stop worker id ", w.WorkerId)
				return
			}
		}
	}()
}

func (w *Worker) StopWorker() {
	w.Done <- true
}

type JobQueue struct {
	Workers    []*Worker
	Done       chan bool
	JobRunning chan Job
}

func NewJobQueue(numOfWorkers int) JobQueue {
	workers := make([]*Worker, numOfWorkers, numOfWorkers)
	jobRunning := make(chan Job)

	for i := 0; i < numOfWorkers; i++ {
		workers[i] = NewWorker(i, jobRunning)
	}

	return JobQueue{
		Workers:    workers,
		Done:       make(chan bool),
		JobRunning: jobRunning,
	}
}

func (jq *JobQueue) Push(job Job) {
	jq.JobRunning <- job
}

func (jq *JobQueue) Stop() {
	jq.Done <- true
}

func (jq *JobQueue) Start() {
	// call .Run()
	go func() {
		for i := 0; i < len(jq.Workers); i++ {
			jq.Workers[i].Run()
		}
	}()

	// wait Done signal then .StopWorker()
	go func() {
		for {
			select {
			case <-jq.Done:
				for i := 0; i < len(jq.Workers); i++ {
					jq.Workers[i].StopWorker()
				}
				return
			}
		}
	}()
}

// define logic job queue to handle
type SenderJob struct {
	Email string
}

func (s SenderJob) Process() { // override from interface Job
	fmt.Println(s.Email)
}

func main() {
	emails := []string{
		"a@gmail.com",
		"b@gmail.com",
		"c@gmail.com",
		"d@gmail.com",
		"e@gmail.com",
	}

	jq := NewJobQueue(4)
	jq.Start()

	for _, email := range emails {
		sender := SenderJob{Email: email}
		jq.Push(sender)
	}

	time.AfterFunc(time.Second*2, func() {
		jq.Stop()
	})

	time.Sleep(time.Second * 6)
}
