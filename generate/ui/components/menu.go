package components

import (
	tea "charm.land/bubbletea/v2"
)

type MenuModel struct {
	options  []MenuOption
	Selected int
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
		s += element.View(index == m.Selected)
	}
	return s
}

func (m MenuModel) Update(msg tea.Msg) (MenuModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Key().Code {
		case tea.KeyDown:
			m.Selected += 1
			if m.Selected >= len(m.options) {
				m.Selected = len(m.options) - 1
			}
		case tea.KeyUp:
			m.Selected -= 1
			if m.Selected < 0 {
				m.Selected = 0
			}
		case tea.KeyEnter:
			return m, m.options[m.Selected].cmd
		case tea.KeyKpEnter:
			return m, m.options[m.Selected].cmd
		}
	}
	return m, nil
}
