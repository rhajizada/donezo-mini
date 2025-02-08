package items

import (
	"context"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rhajizada/donezo-mini/internal/service"
	"github.com/rhajizada/donezo-mini/internal/tui/list"
)

type ItemMenuModel struct {
	ctx     context.Context
	Parent  *service.Board
	List    list.Model
	Input   textinput.Model
	Keys    *Keymap
	Context *InputContext
	Service *service.Service
}

func (m ItemMenuModel) Init() tea.Cmd {
	return m.ListItems()
}

func NewModel(ctx context.Context, service *service.Service, board *service.Board) ItemMenuModel {
	list := list.New(
		[]list.Item{},
		NewDelegate(),
		0,
		0,
	)
	input := textinput.New()
	keymap := NewKeymap()
	inputContext := NewInputContext()
	list.Title = board.Name
	list.AdditionalShortHelpKeys = keymap.ShortHelp
	list.AdditionalFullHelpKeys = keymap.FullHelp

	return ItemMenuModel{
		ctx:     ctx,
		Parent:  board,
		List:    list,
		Input:   input,
		Keys:    keymap,
		Context: inputContext,
		Service: service,
	}
}
