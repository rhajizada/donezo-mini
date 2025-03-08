package tags

type ErrorMsg struct {
	Error error
}

type ListTagsMsg struct {
	Tags []Item
}

type DeleteTagMsg struct {
	Tag   string
	Error error
}
