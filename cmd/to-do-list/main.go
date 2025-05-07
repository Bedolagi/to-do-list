package main

import (
	"to-do-list/internal/menu"
	"to-do-list/internal/task"

	"github.com/eiannone/keyboard"
)

func main() {
	keyboard.Open()
	defer keyboard.Close()
	menu.MenuInit()

	for {
		menu.ClearConsole()
		task.DisplayTasks()
		menu.DisplayMenu()

		_, key, _ := keyboard.GetKey()

		switch key {
		case keyboard.KeyArrowUp:
			if menu.Choice > 0 {
				menu.Choice--
			}
		case keyboard.KeyArrowDown:
			if menu.Choice < len(menu.MenuText)-1 {
				menu.Choice++
			}
		case keyboard.KeyEnter:
			menu.Select()
		}
	}

}
