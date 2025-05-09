package menu

func (menu *Menu) Navigate(direction int) {
	switch direction {
	case -1:
		if menu.choice > 0 {
			menu.choice--
		}
	case 1:
		if menu.choice < len(menu.items)-1 {
			menu.choice++
		}
	}
}

func (menu *Menu) HandleSelection() {
	menu.items[menu.choice].Handler(menu.taskManager)
	menu.taskManager.SaveTasks()
}
