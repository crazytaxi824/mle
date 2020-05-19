package stack

const (
	ErrEmptyStack = "empty stack"
	ErrOutOfRange = "index is out of range"
)

type Option struct {
	AllowDuplicate bool
	StackType      uint8
}

const (
	NormalStack = iota
	ASCStack
	DESCStack
)