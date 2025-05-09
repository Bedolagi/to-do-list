package main

import (
	"to-do-list/internal/menu"
	"to-do-list/internal/task"

	"github.com/eiannone/keyboard"
)

func main() {
	keyboard.Open()
	defer keyboard.Close()

	taskManager := task.NewTaskManager()

	menu := menu.NewMenu(taskManager)
	menu.Start()
}
