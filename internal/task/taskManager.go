package task

import "fmt"

type TaskManager struct {
	tasks  map[int]*Task
	nextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  make(map[int]*Task),
		nextID: 0,
	}
}

func (tm *TaskManager) Add(name string) {
	id := tm.nextID
	tm.tasks[id] = NewTask(name)
	tm.nextID++
}

func (tm *TaskManager) ToggleState(id int) {
	task := tm.tasks[id]
	task.Completed = !task.Completed
}

func (tm *TaskManager) Delete(id int) {
	delete(tm.tasks, id)
}

func (tm *TaskManager) Display() {
	if len(tm.tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	} else {
		fmt.Println("Список задач:")
	}

	for id, task := range tm.tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("ID: %d State: %s\nDescription: %s\n\n", id, status, task.Name)
	}
}
