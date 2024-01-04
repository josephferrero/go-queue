package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func startServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/add-task", addTask)

	http.ListenAndServe(":8080", mux)
}

type addTaskRequest struct {
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

func addTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "incorrect method: %s", r.Method)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "failed to read body")
		return
	}

	task := &addTaskRequest{}
	err = json.Unmarshal(body, task)
	if err != nil {
		fmt.Fprintf(w, "failed to unmarshal json")
		return
	}

	taskQueue.addToQueue(task)
}
