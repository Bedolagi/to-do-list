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
	for key := 0; key < len(List); key++ {
		if List[key].isComplited {
			fmt.Print("ID: " + strconv.Itoa(key+1) + " State: ✅\nDescription:\n" + List[key].name)
		} else {
			fmt.Print("ID: " + strconv.Itoa(key+1) + " State: ❌\nDescription:\n" + List[key].name)
		}
	}
}

func (t *Task) ChangeState() {
	t.isComplited = !t.isComplited
}
