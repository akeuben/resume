package ui

import (
	"resume/generate/data"

	tea "charm.land/bubbletea/v2"
)

type View interface {
	View(resume *data.Resume) string
	Update(tea.Msg) (View, tea.Cmd)
}

type model struct {
	resume *data.Resume
	view   View
}

func InitialModel(resume *data.Resume) model {
	return model{
		resume: resume,
		view:   MakeMainView(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	view, cmd := m.view.Update(msg)
	m.view = view

	return m, cmd
}

func (m model) View() tea.View {
	return tea.NewView(m.view.View(m.resume))
}

type NewViewMsg struct {
	View View
}

func SwitchView(view View) tea.Cmd {
	return func() tea.Msg {
		return NewViewMsg{view}
	}
}
