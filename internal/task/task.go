package task

type Task struct {
	Name      string
	Completed bool
}

func NewTask(name string) *Task {
	return &Task{Name: name}
}
