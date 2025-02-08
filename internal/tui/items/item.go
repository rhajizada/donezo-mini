package items

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/rhajizada/donezo-mini/internal/service"
)

// Item represents item in the list
type Item struct {
	Itm  service.Item
	Tags []string
}

func NewList(items *[]service.Item) []list.Item {
	l := make([]list.Item, len(*items))
	defaultTags := []string{"alpha", "beta", "gamma"}
	for i, item := range *items {
		l[i] = Item{Itm: item, Tags: defaultTags}
	}
	return l
}

func NewItem(item *service.Item) list.Item {
	defaultTags := []string{"alpha", "beta", "gamma"}
	return Item{
		Itm:  *item,
		Tags: defaultTags,
	}
}

func (i Item) Title() string       { return i.Itm.Title }
func (i Item) Description() string { return i.Itm.Description }
func (i Item) FilterValue() string { return i.Itm.Title }
