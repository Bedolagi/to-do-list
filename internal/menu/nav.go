package menu

func (m *Menu) Navigate(direction int) {
	switch direction {
	case -1:
		if m.choice > 0 {
			m.choice--
		}
	case 1:
		if m.choice < len(m.items)-1 {
			m.choice++
		}
	}
}

func (m *Menu) HandleSelection() {
	m.items[m.choice].Handler(m.taskMgr)
}
