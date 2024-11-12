package items

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/rhajizada/donezo-mini/internal/service"
)

// Item represents item in the list
type Item struct {
	Itm service.Item
}

func NewList(items *[]service.Item) []list.Item {
	l := make([]list.Item, len(*items))
	for i, item := range *items {
		l[i] = Item{Itm: item}
	}
	return l
}

func NewItem(item *service.Item) list.Item {
	return Item{
		Itm: *item,
	}
}

func (i Item) Title() string       { return i.Itm.Title }
func (i Item) Description() string { return i.Itm.Description }
func (i Item) FilterValue() string { return i.Itm.Title }
