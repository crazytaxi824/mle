// 单调递增栈
package stack

import (
	"errors"
)

type int64Stack struct {
	items []int64

	// if its true then allow duplicated elements in the stack
	allowDupl bool
}

func NewInt64Stack(allowDupl bool) *int64Stack {
	return &int64Stack{allowDupl: allowDupl}
}

func (s *int64Stack) Len() int {
	return len(s.items)
}

func (s *int64Stack) Elements() []int64 {
	return s.items
}

// 单调栈，会踢出大于自己的元素
func (s *int64Stack) Push(i int64) (n int, res []int64) {
	index := -1
	for k := range s.items {
		if s.allowDupl {
			if s.items[k] > i {
				index = k
				break
			}
		} else {
			if s.items[k] >= i {
				index = k
				break
			}
		}
	}

	if index < 0 {
		s.items = append(s.items, i)
		return 0, nil
	}

	n = len(s.items) - index
	res = s.items[index:]
	s.items = append(s.items[:index:index], i)
	return n, res
}

// 后进先出，返回栈顶元素
func (s *int64Stack) Pop() (int64, error) {
	if len(s.items) == 0 {
		return 0, errors.New(ErrEmptyStack)
	}

	res := s.items[len(s.items)-1]
	s.items = s.items[: len(s.items)-1 : len(s.items)-1]
	return res, nil
}

// 返回栈顶元素，但是不删除该元素
func (s *int64Stack) Peek() (int64, error) {
	if len(s.items) == 0 {
		return 0, errors.New(ErrEmptyStack)
	}

	return s.items[0], nil
}

func (s *int64Stack) Reset() {
	s.items = make([]int64, 0)
}

func (s *int64Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// 返回对象距离栈顶最近的位置(可能有多个相同的元素)
// 位置 1 代表栈顶，即最后一个元素。
// 位置 -1 代表没有找到该元素。
func (s *int64Stack) Search(n int64) int {
	lenS := len(s.items)
	for i := lenS - 1; i >= 0; i-- {
		if s.items[i] == n {
			return lenS - i
		}
	}

	return -1
}
