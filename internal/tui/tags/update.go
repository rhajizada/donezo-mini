package tags

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rhajizada/donezo-mini/internal/tui/styles"
	"golang.design/x/clipboard"
)

// ListTags fetches the list of tags from the client.
func (m *MenuModel) ListTags() tea.Cmd {
	return func() tea.Msg {
		data, err := m.Client.ListTags(m.ctx)
		tags := make([]Item, len(data))
		if err != nil {
			return ErrorMsg{err}
		}
		for i, v := range data {
			count, _ := m.Client.CountItemsByTag(m.ctx, v)
			tags[i] = NewItem(v, count)
		}
		return ListTagsMsg{
			tags,
		}
	}
}

// Copy copies tag to system clipboard
func (m *MenuModel) Copy() tea.Cmd {
	currentName := m.List.SelectedItem().(Item).Tag
	clipboard.Write(clipboard.FmtText, []byte(currentName))
	return m.List.NewStatusMessage(
		styles.StatusMessage.Render(
			fmt.Sprintf("copied \"%s\" to system clipboard", currentName),
		),
	)
}

// DeleteTag deletes current selected tag
func (m *MenuModel) DeleteTag() tea.Cmd {
	return func() tea.Msg {
		selected := m.List.SelectedItem().(Item)
		err := m.Client.DeleteTag(m.ctx, selected.Tag)
		return DeleteTagMsg{Error: err, Tag: selected.Tag}
	}
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

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

	case ListTagsMsg:
		m.List.SetItems(NewList(msg.Tags))

	case DeleteTagMsg:
		cmd := m.HandleDeleteTag(msg)
		cmds = append(cmds, cmd)
		cmd = m.ListTags()
		cmds = append(cmds, cmd)

	}

	listModel, listCmd := m.List.Update(msg)
	m.List = listModel
	cmds = append(cmds, listCmd)

	return m, tea.Batch(cmds...)
}
