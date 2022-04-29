package main

import (
	"fmt"
	"time"
)

// Version information, comes from the build flags (see Makefile)
var (
	version = `¯\_(ツ)_/¯`
)

type Task struct {
	Name      string        // Name of the Task
	StartTime time.Time     // Start time of the Task
	Duration  time.Duration // Duration of the Task
}

type TasksList []Task // List of tasks

func main() {
	fmt.Printf("timer %s\n", version)
}
