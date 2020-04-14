package main

// Task struct
type Task struct {
	Name     string
	Finished bool
}

// NewTask create new task
func NewTask(name string) *Task {
	return &Task{Name: name, Finished: false}
}
