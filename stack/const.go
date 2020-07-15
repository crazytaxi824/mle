package stack

import (
	"errors"
)

var (
	ErrEmptyStack = errors.New("empty stack")
	ErrOutOfRange = errors.New("index is out of range")
)

type Option struct {
	AllowDuplicate bool  // 允许栈中有重复元素
	StackType      uint8 // 栈的类型
}

// 栈的类型
const (
	NormalStack = iota // 普通栈
	ASCStack           // 单调递增栈
	DESCStack          // 单调递减栈
)
