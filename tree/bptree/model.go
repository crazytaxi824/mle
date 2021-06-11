package bptree

import "errors"

var (
	ErrIndexExist    = errors.New("index is already exist")
	ErrIndexNotExist = errors.New("index is not exist")
	ErrInternalValue = errors.New("internal nodes does not have value")
)

type NodeType bool

const (
	Internal NodeType = false
	Leaf     NodeType = true
)

func (nt NodeType) String() string {
	if nt == Internal {
		return "Internal"
	}
	return "Leaf"
}
