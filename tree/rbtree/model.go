package rbtree

import "errors"

var (
	ErrIndexExist   = errors.New("index is already exist")
	ErrNodeNotExist = errors.New("node is not exist")
)

type whichChild bool

const (
	isLeftChild  whichChild = true
	isRightChild whichChild = false
)

type NodeColor bool

// nolint
const (
	// NOTE new node always red, except root.
	// so color red is set to default.
	COLOR_RED NodeColor = false
	COLOR_BLK NodeColor = true
)

func (n NodeColor) String() string {
	if n {
		return "BLACK"
	}
	return "RED"
}
