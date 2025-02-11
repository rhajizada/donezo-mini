package itemsbytag

import (
	"context"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rhajizada/donezo-mini/internal/service"
	"github.com/rhajizada/donezo-mini/internal/tui/list"
	"github.com/rhajizada/donezo-mini/internal/tui/tags"
)

type MenuModel struct {
	ctx     context.Context
	Parent  *tags.MenuModel
	List    list.Model
	Input   textinput.Model
	Keys    *Keymap
	Context *InputContext
	Service *service.Service
}

func (m MenuModel) Init() tea.Cmd {
	return m.ListItems()
}

func NewModel(ctx context.Context, service *service.Service, parent *tags.MenuModel) MenuModel {
	list := list.New(
		[]list.Item{},
		NewDelegate(),
		0,
		0,
	)
	parentItem := parent.List.SelectedItem().(tags.Item)
	input := textinput.New()
	keymap := NewKeymap()
	inputContext := NewInputContext()
	list.Title = parentItem.Tag
	list.AdditionalShortHelpKeys = keymap.ShortHelp
	list.AdditionalFullHelpKeys = keymap.FullHelp

	return MenuModel{
		ctx:     ctx,
		Parent:  parent,
		List:    list,
		Input:   input,
		Keys:    keymap,
		Context: inputContext,
		Service: service,
	}
}
