package menu

import (
	"bufio"
	"fmt"
	"os"
	"to-do-list/internal/task"

	"github.com/eiannone/keyboard"
)

type MenuItem struct {
	Content string
	Handler func(*task.TaskManager)
}

type Menu struct {
	items   []MenuItem
	choice  int
	taskMgr *task.TaskManager
}

func NewMenu(taskMgr *task.TaskManager) *Menu {
	m := &Menu{
		taskMgr: taskMgr,
	}

	m.items = []MenuItem{
		{
			Content: "1. Создать новую задачу",
			Handler: m.createTask,
		},
		{
			Content: "2. Отметить задачу",
			Handler: m.toggleTask,
		},
		{
			Content: "3. Удалить задачу",
			Handler: m.deleteTask,
		},
	}

	return m
}

func (m *Menu) Start() error {
	for {
		ClearConsole()
		m.taskMgr.Display()
		m.Display()

		_, key, _ := keyboard.GetKey()

		switch key {
		case keyboard.KeyArrowUp:
			m.Navigate(-1)
		case keyboard.KeyArrowDown:
			m.Navigate(1)
		case keyboard.KeyEnter:
			m.HandleSelection()
		}
	}
}

func (m *Menu) createTask(tm *task.TaskManager) {
	fmt.Print("Введите описание новой задачи: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskName := scanner.Text()

	tm.Add(taskName)
}

func (m *Menu) toggleTask(tm *task.TaskManager) {
	fmt.Print("Введите ID задачи: ")
	var id int
	fmt.Scanln(&id)

	tm.ToggleState(id)
}

func (m *Menu) deleteTask(tm *task.TaskManager) {
	fmt.Print("Введите ID задачи для удаления: ")
	var id int
	fmt.Scanln(&id)

	tm.Delete(id)
}
