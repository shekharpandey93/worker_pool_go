package main

import "fmt"

func main() {
	// Create new tasks

	tasks := []Task{
		&EmailProcess{Email: "email1@sp.com", Subject: "test", Message: "test1"},
		&ImageProcess{ImagePath: "/images/sample1.jpg"},
		&EmailProcess{Email: "email2@sp.com", Subject: "test1", Message: "test2"},
		&ImageProcess{ImagePath: "/images/sample2.jpg"},
		&EmailProcess{Email: "email3@sp.com", Subject: "test2", Message: "test3"},
		&ImageProcess{ImagePath: "/images/sample3.jpg"},
		&EmailProcess{Email: "email4@sp.com", Subject: "test3", Message: "test4"},
		&ImageProcess{ImagePath: "/images/sample4.jpg"},
		&EmailProcess{Email: "email5@sp.com", Subject: "test4", Message: "test5"},
		&ImageProcess{ImagePath: "/images/sample5.jpg"},
	}

	wg := workerPool{
		Tasks:       tasks,
		concurrency: 4,
	}

	wg.Start()
	fmt.Println("All tasks have been processed!")
}
