package components

import (
	"resume/generate/data"
	"slices"

	tea "charm.land/bubbletea/v2"
)

type ListEditorModel[T data.DisplayableListItem] struct {
	List     *[]T
	Selected int
	Length   int
}

func CreateListEditor[T data.DisplayableListItem](list *[]T, selected int) ListEditorModel[T] {
	return ListEditorModel[T]{list, selected, len(*list)}
}

func (m ListEditorModel[T]) View() string {
	s := "\n\n"
	for index, element := range *m.List {
		if index >= m.Length {
			break
		}
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
		switch msg.String() {
		case "down":
			m.Selected += 1
			if m.Selected >= m.Length {
				m.Selected = m.Length - 1
			}
		case "up":
			m.Selected -= 1
			if m.Selected < 0 {
				m.Selected = 0
			}
		case "shift+down":
			if m.Selected >= m.Length-1 {
				return m, nil
			}
			(*m.List)[m.Selected], (*m.List)[m.Selected+1] = (*m.List)[m.Selected+1], (*m.List)[m.Selected]
			m.Selected += 1
		case "shift+up":
			if m.Selected <= 0 {
				return m, nil
			}
			(*m.List)[m.Selected], (*m.List)[m.Selected-1] = (*m.List)[m.Selected-1], (*m.List)[m.Selected]
			m.Selected -= 1
		case "delete":
			if m.Length == 0 {
				return m, nil
			}
			_ = slices.Delete(*m.List, m.Selected, m.Selected+1)
			m.Length -= 1
			m.Selected = min(m.Length-1, m.Selected)
		}
	}
	return m, nil
}
