package main

import (
	"fmt"
	"log"
)

type Pool struct {
	work chan func()
	sem  chan struct{}
}

func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func (p *Pool) Schedule(task func(), id int) error {
	defer fmt.Println("--- end Schedule")

	select {
	case p.work <- task:
		fmt.Println("+ new work")
	case p.sem <- struct{}{}:
		fmt.Println("+ new sem")
		go p.worker(task, id)
	}

	return nil
}

func (p *Pool) worker(task func(), id int) {
	defer func() {
		<-p.sem
		log.Printf("worker %v done!\n", id)
	}()

	for {
		log.Printf("...start new task %v...\n", id)
		task()
		log.Printf("...done 1 task %v...\n", id)
		task = <-p.work
	}
}
