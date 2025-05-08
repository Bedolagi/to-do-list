package main

import (
	"to-do-list/internal/menu"
	"to-do-list/internal/task"

	"github.com/eiannone/keyboard"
)

func main() {
	keyboard.Open()
	defer keyboard.Close()

	taskMng := task.NewTaskManager()

	menu := menu.NewMenu(taskMng)
	menu.Start()
}
