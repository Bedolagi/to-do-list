package menu

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"to-do-list/internal/task"
)

var MenuText map[int]string = make(map[int]string)
var Choice int
var availableID int

func MenuInit() {
	MenuText[0] = "1. Создать новую заметку"
	MenuText[1] = "2. Отметить заметку"
	MenuText[2] = "3. Удалить заметку"
}

func DisplayMenu() {
	fmt.Println("📝To-Do List📝")
	for i := 0; i < len(MenuText); i++ {
		if Choice != i {
			fmt.Printf("%v\n", MenuText[i])
		} else {
			fmt.Printf("%v %v\n", "➡️", MenuText[i])
		}
	}
}

func Select() {
	switch Choice {
	case 0:
		var taskName string
		fmt.Print("Введите имя новой заметки:")
		fmt.Scan(&taskName)

		task.List[availableID] = task.NewTask(taskName, false)
		availableID++
	case 1:
		var selectTask int
		fmt.Println("Введите ID задачи:")
		fmt.Scan(&selectTask)
		task.List[selectTask-1].ChangeState()
	case 2:
		var idTask int
		fmt.Println("Введите ID задачи:")
		fmt.Scan(&idTask)
		delete(task.List, idTask-1)
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
