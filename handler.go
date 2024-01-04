package main

import (
	"fmt"
	"strconv"
	"time"
)

func handleTask(t *task) {
	fmt.Println("handling task")
	switch t.Name {
	case "timeout":
		handleTimeoutTask(t)
	}
}

func handleTimeoutTask(t *task) {
	fmt.Println("handling timeout task")
	timeout, ok := t.Payload["timeout"]
	if !ok {
		fmt.Println("no timeout given for timeout task")
		return
	}
	i, err := strconv.Atoi(timeout.(string))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("sleeping for %v seconds\n", i)
	time.Sleep(time.Duration(i) * time.Second)
}
