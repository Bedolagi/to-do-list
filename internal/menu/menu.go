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
	items       []MenuItem
	choice      int
	taskManager *task.TaskManager
}

func NewMenu(taskManager *task.TaskManager) *Menu {
	menu := &Menu{
		taskManager: taskManager,
	}

	menu.items = []MenuItem{
		{
			Content: "1. Создать новую задачу",
			Handler: menu.createTask,
		},
		{
			Content: "2. Изменение описания задачи",
			Handler: menu.editTask,
		},
		{
			Content: "3. Отметить задачу",
			Handler: menu.toggleTask,
		},
		{
			Content: "4. Удалить задачу",
			Handler: menu.deleteTask,
		},
		{
			Content: "5. Удалить все задачи",
			Handler: menu.deleteAllTask,
		},
	}

	return menu
}

func (menu *Menu) Start() error {
	for {
		ClearConsole()
		menu.taskManager.Display()
		menu.Display()

		_, key, _ := keyboard.GetKey()

		switch key {
		case keyboard.KeyArrowUp:
			menu.Navigate(-1)
		case keyboard.KeyArrowDown:
			menu.Navigate(1)
		case keyboard.KeyEnter:
			menu.HandleSelection()
		}
	}
}

func (menu *Menu) createTask(taskManager *task.TaskManager) {
	fmt.Print("Введите описание новой задачи: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskName := scanner.Text()

	taskManager.Add(taskName)
}

func (menu *Menu) editTask(taskManager *task.TaskManager) {
	fmt.Print("Введите ID задачи: ")
	var id int
	fmt.Scanln(&id)

	fmt.Print("Введите описание новой задачи: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskDescription := scanner.Text()

	taskManager.EditTask(id, taskDescription)
}

func (menu *Menu) toggleTask(taskManager *task.TaskManager) {
	fmt.Print("Введите ID задачи: ")
	var id int
	fmt.Scanln(&id)

	taskManager.ToggleState(id)
}

func (menu *Menu) deleteTask(taskManager *task.TaskManager) {
	fmt.Print("Введите ID задачи для удаления: ")
	var id int
	fmt.Scanln(&id)

	taskManager.Delete(id)
}

func (menu *Menu) deleteAllTask(taskManager *task.TaskManager) {
	for key := range taskManager.GetTasks() {
		taskManager.Delete(key)
	}
}
