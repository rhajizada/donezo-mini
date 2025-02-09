package items

import (
	"strings"

	"github.com/rhajizada/donezo-mini/internal/service"
	"github.com/rhajizada/donezo-mini/internal/tui/list"
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
func (i Item) Footer() string {
	var message string
	if len(i.Itm.Tags) > 0 {
		message += "Tags: "
		message += strings.Join(i.Itm.Tags, ", ")
	} else {
		message += "No tags"
	}
	return message
}
func (i Item) FilterValue() string { return i.Itm.Title }
