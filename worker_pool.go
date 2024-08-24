package main

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	Process()
}

type EmailProcess struct {
	Email   string
	Subject string
	Message string
}

func (e *EmailProcess) Process() {
	fmt.Printf("Sending email to %s\n", e.Email)
	// Simulate a time consuming process
	time.Sleep(2 * time.Second)
}

type ImageProcess struct {
	ImagePath string
}

func (i *ImageProcess) Process() {
	fmt.Printf("Sending email to %s\n", i.ImagePath)
	// Simulate a time consuming process
	time.Sleep(5 * time.Second)
}

type workerPool struct {
	wg          sync.WaitGroup
	ch          chan Task
	concurrency int
	Tasks       []Task
}

func (w *workerPool) workerJob() {
	for task := range w.ch {
		task.Process()
		w.wg.Done()
	}
}

func (w *workerPool) Start() {
	w.ch = make(chan Task, w.concurrency)
	for i := 0; i < w.concurrency; i++ {
		go w.workerJob()
	}

	w.wg.Add(len(w.Tasks))
	for _, task := range w.Tasks {
		w.ch <- task
	}
	close(w.ch)

	w.wg.Wait()
}
