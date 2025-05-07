package task

import (
	"fmt"
	"strconv"
)

var List map[int]*Task = make(map[int]*Task)

type Task struct {
	name        string
	isComplited bool
}

func NewTask(n string, state bool) *Task {
	return &Task{name: n, isComplited: state}
}

func DisplayTasks() {
	for key, v := range List {
		if v.isComplited {
			fmt.Println(strconv.Itoa(key+1) + " " + v.name + "✅")
		} else {
			fmt.Println(strconv.Itoa(key+1) + " " + v.name + "❌")
		}
	}
}

func (t *Task) ChangeState() {
	t.isComplited = !t.isComplited
}
