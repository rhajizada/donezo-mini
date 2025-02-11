package tags

import (
	"github.com/charmbracelet/bubbles/key"
)

// Keymap embeds default list keymap and adds other Binding
type Keymap struct {
	Choose      key.Binding
	ListBoards  key.Binding
	DeleteTag   key.Binding
	RefreshList key.Binding
}

func NewKeymap() Keymap {
	return Keymap{
		Choose: key.NewBinding(
			key.WithKeys("return"),
			key.WithHelp("return", "choose tag"),
		),
		ListBoards: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "list boards"),
		),
		DeleteTag: key.NewBinding(key.WithKeys("d"),
			key.WithHelp("d", "delete tag"),
		),
		RefreshList: key.NewBinding(key.WithKeys("R"),
			key.WithHelp("R", "refresh list"),
		),
	}
}

func (km Keymap) ShortHelp() []key.Binding {
	bindings := []key.Binding{}
	bindings = append(bindings, km.Choose)
	bindings = append(bindings, km.ListBoards)
	return bindings
}

func (km Keymap) FullHelp() []key.Binding {
	bindings := []key.Binding{}
	bindings = append(bindings, km.Choose)
	bindings = append(bindings, km.DeleteTag)
	bindings = append(bindings, km.RefreshList)
	return bindings
}
