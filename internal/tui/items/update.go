package items

import (
	"strings"

	"github.com/rhajizada/donezo-mini/internal/tui/styles"

	tea "github.com/charmbracelet/bubbletea"
)

const TagsSeparator = ", "

// ListItems fetches items ine the selected board.
func (m *ItemMenuModel) ListItems() tea.Cmd {
	return func() tea.Msg {
		items, err := m.Service.ListItems(m.ctx, m.Parent)
		if err != nil {
			return ErrorMsg{err}
		}
		return ListItemsMsg{
			items,
		}
	}
}

// CreateItem creates a new item
func (m *ItemMenuModel) CreateItem() tea.Cmd {
	return func() tea.Msg {
		item, err := m.Service.CreateItem(m.ctx, m.Parent, m.Context.Title, m.Context.Desc)
		return CreateItemMsg{
			item,
			err,
		}
	}
}

func (m *ItemMenuModel) InitCreateItem() tea.Cmd {
	m.Context.State = CreateItemNameState
	m.Input.Placeholder = "Enter item name"
	m.Input.SetValue("")
	m.Input.Focus()
	return nil
}

// RenameItem renames selected item
func (m *ItemMenuModel) RenameItem() tea.Cmd {
	return func() tea.Msg {
		selected := m.List.SelectedItem().(Item)
		selected.Itm.Title = m.Context.Title
		selected.Itm.Description = m.Context.Desc
		item, err := m.Service.UpdateItem(m.ctx, &selected.Itm)
		return RenameItemMsg{
			item,
			err,
		}
	}
}

// UpdateTags updates item tags
func (m *ItemMenuModel) UpdateTags() tea.Cmd {
	return func() tea.Msg {
		selected := m.List.SelectedItem().(Item)
		tags := strings.Split(m.Context.Title, TagsSeparator)
		selected.Itm.Tags = tags
		item, err := m.Service.UpdateItem(m.ctx, &selected.Itm)
		return UpdateTagsMsg{
			item,
			err,
		}
	}
}

// initiateRename starts the renaming process for the selected item.
func (m *ItemMenuModel) InitRenameItem() tea.Cmd {
	if len(m.List.Items()) == 0 {
		return m.List.NewStatusMessage(
			styles.StatusMessage.Render("no item selected"))
	}

	m.Context.State = RenameItemNameState
	selected := m.List.SelectedItem().(Item)
	m.Input.SetValue(selected.Itm.Title)
	m.Input.Focus()
	return nil
}

// InitUpdateTags initiaizes tag updates
func (m *ItemMenuModel) InitUpdateTags() tea.Cmd {
	m.Context.State = CreateItemNameState
	m.Context.State = UpdateTagsState
	m.Input.Placeholder = "Enter comma-separated list of tags"
	selected := m.List.SelectedItem().(Item)
	m.Input.SetValue(strings.Join(selected.Itm.Tags, TagsSeparator))
	m.Input.Focus()
	return nil
}

// DeleteBoard deletes current selected board
func (m *ItemMenuModel) DeleteItem() tea.Cmd {
	return func() tea.Msg {
		selected := m.List.SelectedItem().(Item)
		err := m.Service.DeleteItem(m.ctx, &selected.Itm)
		return DeleteItemMsg{Error: err, Item: &selected.Itm}
	}
}

func (m ItemMenuModel) ToggleComplete() tea.Cmd {
	if len(m.List.Items()) == 0 {
		return m.List.NewStatusMessage(
			styles.ErrorMessage.Render("no item selected"))
	}

	selected := m.List.SelectedItem().(Item)
	selected.Itm.Completed = !selected.Itm.Completed
	m.List.SetItem(m.List.Index(), selected)

	return func() tea.Msg {
		i, err := m.Service.UpdateItem(m.ctx, &selected.Itm)
		return ToggleItemMsg{i, err}
	}
}

func (m ItemMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	if m.Context.State != DefaultState {
		m.Input, cmds = m.HandleInputState(msg)
		return m, tea.Batch(cmds...)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cmd := m.HandleWindowSize(msg)
		cmds = append(cmds, cmd)

	case tea.KeyMsg:
		cmd := m.HandleKeyInput(msg)
		cmds = append(cmds, cmd)

	case ErrorMsg:
		cmd := m.HandleError(msg)
		cmds = append(cmds, cmd)

	case ListItemsMsg:
		m.List.SetItems(NewList(msg.Items))

	case CreateItemMsg:
		cmd := m.HandleCreateItem(msg)
		cmds = append(cmds, cmd)

	case DeleteItemMsg:
		cmd := m.HandleDeleteItem(msg)
		cmds = append(cmds, cmd)
		cmd = m.ListItems()
		cmds = append(cmds, cmd)

	case RenameItemMsg:
		cmd := m.HandleRenameItem(msg)
		cmds = append(cmds, cmd)

	case UpdateTagsMsg:
		cmd := m.HandleUpdateTags(msg)
		cmds = append(cmds, cmd)

	case ToggleItemMsg:
		cmd := m.HandleToggleItem(msg)
		cmds = append(cmds, cmd)
	}

	listModel, listCmd := m.List.Update(msg)
	m.List = listModel
	cmds = append(cmds, listCmd)

	return m, tea.Batch(cmds...)
}
