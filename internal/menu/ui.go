package menu

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func (m *Menu) Display() {
	fmt.Println("📝 To-Do List 📝")
	fmt.Println("Используйте стрелки для перемещения по меню.\nЧтобы выбрать пункт меню, нажмите ENTER")

	for i, item := range m.items {
		if m.choice == i {
			fmt.Printf("➡️  %s\n", item.Content)
		} else {
			fmt.Printf("   %s\n", item.Content)
		}
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
