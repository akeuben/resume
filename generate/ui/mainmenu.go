package ui

import (
	"resume/generate/data"
	"resume/generate/ui/components"

	tea "charm.land/bubbletea/v2"
)

type MainView struct {
	menu components.MenuModel
}

func MakeMainView(resume *data.Resume) MainView {
	menu := components.CreateMenu([]components.MenuOption{
		components.CreateMenuOption("Skills", "Add, modify, reorder, or remove skills for this resume instance", SwitchView(MakeSubMenuView(&resume.Skills, resume))),
		components.CreateMenuOption("Education", "Add modify, or remove education items", SwitchView(MakeSubMenuView(&resume.Education, resume))),
		components.CreateMenuOption("Experience", "Add, modify, or remove experience items", SwitchView(MakeSubMenuView(&resume.Experience, resume))),
		components.CreateMenuOption("Projects", "Choose projects to display on resume", SwitchView(MakeSubMenuView(&resume.Projects, resume))),
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
