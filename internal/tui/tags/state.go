package tags

type InputState uint8

const (
	DefaultState InputState = iota
	CreateTagState
	RenameTagState
)
