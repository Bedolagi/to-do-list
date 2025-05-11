package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"to-do-list/internal/storage"
)

type TaskManager struct {
	tasks  map[int]*Task
	nextID int
}

func NewTaskManager() *TaskManager {
	taskManager := &TaskManager{
		tasks:  make(map[int]*Task),
		nextID: 0,
	}
	taskManager.LoadTasks()
	return taskManager
}

func (taskManager *TaskManager) Add(description string) {
	id := taskManager.nextID
	taskManager.tasks[id] = NewTask(description)
	taskManager.nextID++
}

func (taskManager *TaskManager) EditTask(id int, description string) {
	taskManager.tasks[id].Description = description
}

func (taskManager *TaskManager) GetTasks() map[int]*Task {
	return taskManager.tasks
}

func (taskManager *TaskManager) ToggleState(id int) {
	task := taskManager.tasks[id]
	task.Completed = !task.Completed
}

func (taskManager *TaskManager) Delete(id int) {
	delete(taskManager.tasks, id)
}

func (taskManager *TaskManager) Display() {
	if len(taskManager.tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	} else {
		fmt.Println("Список задач:")
	}

	for id, task := range taskManager.tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("ID: %d State: %s\nDescription: %s\n\n", id, status, task.Description)
	}
}

func (taskManager *TaskManager) SaveTasks() {
	configPath := storage.InitConfig()

	data, _ := json.Marshal(taskManager.tasks)

	configFile := filepath.Join(configPath, "tasks.json")
	os.WriteFile(configFile, data, 0644)
}

func (taskManager *TaskManager) LoadTasks() {
	configPath := storage.InitConfig()

	configFile := filepath.Join(configPath, "tasks.json")
	data, err := os.ReadFile(configFile)
	if !os.IsNotExist(err) {
		json.Unmarshal(data, &taskManager.tasks)
	}
}
