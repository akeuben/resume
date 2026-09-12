package ui

import (
	"resume/generate/data"
	"resume/generate/ui/components"

	tea "charm.land/bubbletea/v2"
)

type SubMenuView[T data.DisplayableListItem] struct {
	editor components.ListEditorModel[T]
	resume *data.Resume
}

func MakeSubMenuView[T data.DisplayableListItem](list *[]T, resume *data.Resume) SubMenuView[T] {
	editor := components.CreateListEditor[T](list, 0)

	return SubMenuView[T]{editor, resume}
}

func (view SubMenuView[T]) View(resume *data.Resume) string {
	s := "Main Menu"

	s += view.editor.View()

	return s
}

func (view SubMenuView[T]) Update(msg tea.Msg) (View, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q":
			return view, SwitchView(MakeMainView(view.resume))
		}
	}

	v, cmd := view.editor.Update(msg)
	view.editor = v
	return view, cmd
}
