package boards

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rhajizada/donezo-mini/internal/service"
)

type MenuModel struct {
	ctx    context.Context
	List   list.Model
	Input  textinput.Model
	Keys   *Keymap
	State  InputState
	Client *service.Service
}

func (m MenuModel) Init() tea.Cmd {
	return m.ListBoards()
}

func NewModel(ctx context.Context, client *service.Service) MenuModel {
	list := list.New(
		[]list.Item{},
		list.NewDefaultDelegate(),
		0,
		0,
	)
	input := textinput.New()
	keymap := NewKeymap()
	list.Title = "donezo"
	list.AdditionalShortHelpKeys = keymap.ShortHelp
	list.AdditionalFullHelpKeys = keymap.FullHelp
	return MenuModel{
		ctx:    ctx,
		List:   list,
		Input:  input,
		State:  DefaultState,
		Keys:   &keymap,
		Client: client,
	}
}
