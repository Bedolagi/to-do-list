package menu

import (
	"fmt"
	"os"
	"os/exec"
)

func (menu *Menu) Display() {
	fmt.Println("📝 To-Do List 📝")
	fmt.Println("Используйте стрелки для перемещения по меню.\nЧтобы выбрать пункт меню, нажмите ENTER")

	for i, item := range menu.items {
		if menu.choice == i {
			fmt.Printf("➡️  %s\n", item.Content)
		} else {
			fmt.Printf("   %s\n", item.Content)
		}
	}
}

func ClearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
