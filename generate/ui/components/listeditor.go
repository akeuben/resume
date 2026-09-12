package components

import (
	"resume/generate/data"

	tea "charm.land/bubbletea/v2"
)

type ListEditorModel[T data.DisplayableListItem] struct {
	List     *[]T
	Selected int
}

func CreateListEditor[T data.DisplayableListItem](list *[]T, selected int) ListEditorModel[T] {
	return ListEditorModel[T]{list, selected}
}

func (m ListEditorModel[T]) View() string {
	s := "\n\n"
	for index, element := range *m.List {
		s += ViewItem(element.Display(), index == m.Selected)
	}
	return s
}

func ViewItem(text string, selected bool) string {
	s := ""
	if selected {
		s += "> "
	} else {
		s += "  "
	}
	s += text
	s += "\n\n"
	return s
}

func (m ListEditorModel[T]) Update(msg tea.Msg) (ListEditorModel[T], tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Key().Code {
		case tea.KeyDown:
			m.Selected += 1
			if m.Selected >= len(*m.List) {
				m.Selected = len(*m.List) - 1
			}
		case tea.KeyUp:
			m.Selected -= 1
			if m.Selected < 0 {
				m.Selected = 0
			}
		}
	}
	return m, nil
}
