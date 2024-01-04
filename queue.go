package main

import (
	"fmt"
	"sync"
)

type task struct {
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

type queue struct {
	mu    sync.Mutex
	tasks []task
}

var taskQueue = queue{}

func (q *queue) addToQueue(taskReq *addTaskRequest) {
	fmt.Println("adding task to queue")
	task := task{Name: taskReq.Name, Payload: taskReq.Payload}
	q.mu.Lock()
	q.tasks = append(q.tasks, task)
	q.mu.Unlock()
}

func (q *queue) processQueue() {
	fmt.Println("processing queue")
	go func() {
		for {
			q.mu.Lock()
			if len(q.tasks) == 0 {
				q.mu.Unlock()
				continue
			}
			task := q.tasks[0]
			q.tasks = q.tasks[1:]
			q.mu.Unlock()
			handleTask(&task)
		}
	}()
}
