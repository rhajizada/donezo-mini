package app

import (
	"context"

	"github.com/rhajizada/donezo-mini/internal/tui/boards"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rhajizada/donezo-mini/internal/service"
)

type AppModel struct {
	ctx            context.Context
	Service        *service.Service
	ViewStack      []tea.Model
	LastWindowSize *tea.WindowSizeMsg // Store the last known WindowSizeMsg
}

func NewModel(ctx context.Context, service *service.Service) AppModel {
	boardMenu := boards.NewModel(ctx, service)
	return AppModel{
		ctx:            ctx,
		Service:        service,
		ViewStack:      []tea.Model{boardMenu},
		LastWindowSize: nil, // Initialize without a WindowSizeMsg
	}
}

func (m AppModel) Init() tea.Cmd {
	// Initialize the top model in the stack
	return m.ViewStack[len(m.ViewStack)-1].Init()
}

func (m *AppModel) GetCurrentBoard() *service.Board {
	// Retrieve the currently selected board from the board menu
	if boardMenu, ok := m.ViewStack[0].(boards.MenuModel); ok {
		if selected, ok := boardMenu.List.SelectedItem().(boards.Item); ok {
			return &selected.Board
		}
	}
	return nil
}
