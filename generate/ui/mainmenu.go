package ui

import (
	"resume/generate/data"
	"resume/generate/ui/components"

	tea "charm.land/bubbletea/v2"
)

type MainView struct {
	menu components.MenuModel
}

func MakeMainView() MainView {
	menu := components.CreateMenu([]components.MenuOption{
		components.CreateMenuOption("Skills", "Add, modify, reorder, or remove skills for this resume instance", nil),
		components.CreateMenuOption("Education", "Add modify, or remove education items", nil),
		components.CreateMenuOption("Experience", "Add, modify, or remove experience items", nil),
		components.CreateMenuOption("Projects", "Choose projects to display on resume", nil),
	}, 0)

	return MainView{menu}
}

func (view MainView) View(resume *data.Resume) string {
	s := "Main Menu"

	s += view.menu.View()

	return s
}

func (view MainView) Update(msg tea.Msg) (View, tea.Cmd) {
	v, cmd := view.menu.Update(msg)
	view.menu = v
	return view, cmd
}
