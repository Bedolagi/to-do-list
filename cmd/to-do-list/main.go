package main

import (
	"to-do-list/internal/menu"

	"github.com/eiannone/keyboard"
)

func main() {
	keyboard.Open()
	defer keyboard.Close()
	menu.Init()

	menu.Start()
}
