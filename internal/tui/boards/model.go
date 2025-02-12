package boards

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rhajizada/donezo-mini/internal/service"
	"github.com/rhajizada/donezo-mini/internal/stack"
)

type MenuModel struct {
	ctx       context.Context
	List      list.Model
	Input     textinput.Model
	Keys      *Keymap
	State     InputState
	ItemStack *stack.Stack[service.Item]
	Client    *service.Service
}

func (m MenuModel) Init() tea.Cmd {
	return m.ListBoards()
}

func New(ctx context.Context, client *service.Service) MenuModel {
	list := list.New(
		[]list.Item{},
		list.NewDefaultDelegate(),
		0,
		0,
	)
	input := textinput.New()
	keymap := NewKeymap()
	itemStack := stack.New[service.Item](10)
	list.Title = "donezo | Boards"
	list.AdditionalShortHelpKeys = keymap.ShortHelp
	list.AdditionalFullHelpKeys = keymap.FullHelp
	return MenuModel{
		ctx:       ctx,
		List:      list,
		Input:     input,
		Keys:      &keymap,
		State:     DefaultState,
		ItemStack: itemStack,
		Client:    client,
	}
}
