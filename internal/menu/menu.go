package menu

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"to-do-list/internal/task"

	"github.com/eiannone/keyboard"
)

var MenuText map[int]string = make(map[int]string)
var Choice int
var availableID int

func Init() {
	MenuText[0] = "1. Создать новую задачу"
	MenuText[1] = "2. Отметить задачу"
	MenuText[2] = "3. Удалить задачу"
}

func DisplayMenu() {
	fmt.Println("📝 To-Do List📝")
	fmt.Println("Используйте стрелки для перемещения по меню.\nЧтобы выбрать пункт меню, нажмите ENTER")
	for i := 0; i < len(MenuText); i++ {
		if Choice != i {
			fmt.Printf("%v\n", MenuText[i])
		} else {
			fmt.Printf("%v %v\n", "➡️ ", MenuText[i])
		}
	}
}

func Select() {
	switch Choice {
	case 0:
		var taskName string
		fmt.Print("Введите описание новой задачи:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		taskName = scanner.Text()

		task.List[availableID] = task.NewTask(taskName, false)
		availableID++
	case 1:
		var selectTask int
		fmt.Println("Введите ID задачи:")
		fmt.Scanln(&selectTask)
		task.List[selectTask-1].ChangeState()
	case 2:
		var idTask int
		fmt.Println("Введите ID задачи:")
		fmt.Scanln(&idTask)
		delete(task.List, idTask-1)
		fmt.Println("Задача удалена!")
	}
}

func ClearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[H\033[2J")
	}
}

func Start() {
	for {
		ClearConsole()
		task.DisplayTasks()
		DisplayMenu()

		_, key, _ := keyboard.GetKey()

		switch key {
		case keyboard.KeyArrowUp:
			if Choice > 0 {
				Choice--
			}
		case keyboard.KeyArrowDown:
			if Choice < len(MenuText)-1 {
				Choice++
			}
		case keyboard.KeyEnter:
			Select()
		}
	}
}
