package avltree

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
