package itemsbytag

type InputState uint8

const (
	DefaultState InputState = iota
	RenameItemNameState
	RenameItemDescState
	UpdateTagsState
)

type InputContext struct {
	State InputState
	Title string
	Desc  string
}

func NewInputContext() *InputContext {
	return &InputContext{
		State: DefaultState,
		Title: "",
		Desc:  "",
	}
}
