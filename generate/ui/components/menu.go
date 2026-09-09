package components

import (
	tea "charm.land/bubbletea/v2"
)

type MenuModel struct {
	options  []MenuOption
	selected int
}

type MenuOption struct {
	text        string
	description string
	cmd         tea.Cmd
}

func CreateMenu(options []MenuOption, selected int) MenuModel {
	return MenuModel{options, selected}
}

func CreateMenuOption(text string, description string, cmd tea.Cmd) MenuOption {
	return MenuOption{text, description, cmd}
}

func (m MenuOption) View(selected bool) string {
	s := ""
	if selected {
		s += "> "
	} else {
		s += "  "
	}
	s += m.text
	s += "\n"
	s += "  "
	s += m.description
	s += "\n\n"
	return s
}

func (m MenuModel) View() string {
	s := "\n\n"
	for index, element := range m.options {
		s += element.View(index == m.selected)
	}
	return s
}

func (m MenuModel) Update(msg tea.Msg) (MenuModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Key().Code {
		case tea.KeyDown:
			m.selected += 1
			if m.selected >= len(m.options) {
				m.selected = len(m.options) - 1
			}
		case tea.KeyUp:
			m.selected -= 1
			if m.selected < 0 {
				m.selected = 0
			}
		}
	}
	return m, nil
}
