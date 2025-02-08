package items

import (
	"strings"

	"github.com/rhajizada/donezo-mini/internal/tui/styles"
)

func (m ItemMenuModel) View() string {
	if m.Context.State != DefaultState {
		return styles.App.Render(m.Input.View())
	}
	item, ok := m.List.SelectedItem().(Item)
	var tagView string
	if ok {
		tags := strings.Join(item.Tags, ", ")
		tagView = styles.Tags.Render(tags)
	}
	return styles.App.Render(m.List.View(), tagView)
}
