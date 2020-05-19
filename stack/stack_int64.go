// 单调递增栈
package stack

import (
	"errors"
)

// 单调递增栈
type int64ASCStack struct {
	items []int64

	// if its true then allow duplicated elements in the stack
	allowDupl bool
}

func NewInt64ASCStack(allowDupl bool) *int64ASCStack {
	return &int64ASCStack{allowDupl: allowDupl}
}

// 返回 stack 长度
func (s *int64ASCStack) Len() int {
	return len(s.items)
}

// 返回所有 element 到 slice
func (s *int64ASCStack) Elements() []int64 {
	return s.items
}

// 单调栈，会踢出大于自己的元素
// n 表示有多少元素被踢出
// return n means how many items has been removed from stack
func (s *int64ASCStack) Push(i int64) (n int, res []int64) {
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

	res = s.items[index:]
	s.items = append(s.items[:index:index], i)
	return len(res), res
}

// 后进先出，返回栈顶元素
func (s *int64ASCStack) Pop() (int64, error) {
	if len(s.items) == 0 {
		return 0, errors.New(ErrEmptyStack)
	}

	res := s.items[len(s.items)-1]
	s.items = s.items[: len(s.items)-1 : len(s.items)-1]
	return res, nil
}

// 返回栈顶元素，但是不删除该元素
func (s *int64ASCStack) Peek() (int64, error) {
	if len(s.items) == 0 {
		return 0, errors.New(ErrEmptyStack)
	}

	return s.items[len(s.items)-1], nil
}

// 重置 stack
func (s *int64ASCStack) Reset() {
	s.items = make([]int64, 0)
}

// 判断 stack 是否为空
func (s *int64ASCStack) IsEmpty() bool {
	return len(s.items) == 0
}

// 返回对象距离栈顶最近的位置(可能有多个相同的元素)
// 位置 1 代表栈顶，即最后一个元素。
// 位置 -1 代表没有找到该元素。
func (s *int64ASCStack) Search(n int64) int {
	lenS := len(s.items)
	for i := lenS - 1; i >= 0; i-- {
		if s.items[i] == n {
			return lenS - i
		}
	}

	return -1
}

// 单调递减栈
type int64DESCStack struct {
	items []int64

	// if its true then allow duplicated elements in the stack
	allowDupl bool
}

func NewInt64DESCStack(allowDupl bool) *int64DESCStack {
	return &int64DESCStack{allowDupl: allowDupl}
}

// 返回 stack 长度
func (s *int64DESCStack) Len() int {
	return len(s.items)
}

// 返回所有 element 到 slice
func (s *int64DESCStack) Elements() []int64 {
	return s.items
}

// 单调栈，会踢出小于自己的元素,
// n 表示有多少元素被踢出
// return n means how many items has been removed from stack
func (s *int64DESCStack) Push(i int64) (n int, res []int64) {
	index := -1
	for k := range s.items {
		if s.allowDupl {
			if s.items[k] < i {
				index = k
				break
			}
		} else {
			if s.items[k] <= i {
				index = k
				break
			}
		}
	}

	if index < 0 {
		s.items = append(s.items, i)
		return 0, nil
	}

	res = s.items[index:]
	s.items = append(s.items[:index:index], i)
	return len(res), res
}

// 后进先出，返回栈顶元素
func (s *int64DESCStack) Pop() (int64, error) {
	if len(s.items) == 0 {
		return 0, errors.New(ErrEmptyStack)
	}

	res := s.items[len(s.items)-1]
	s.items = s.items[: len(s.items)-1 : len(s.items)-1]
	return res, nil
}

// 返回栈顶元素，但是不删除该元素
func (s *int64DESCStack) Peek() (int64, error) {
	if len(s.items) == 0 {
		return 0, errors.New(ErrEmptyStack)
	}

	return s.items[len(s.items)-1], nil
}

// 重置 stack
func (s *int64DESCStack) Reset() {
	s.items = make([]int64, 0)
}

// 判断 stack 是否为空
func (s *int64DESCStack) IsEmpty() bool {
	return len(s.items) == 0
}

// 返回对象距离栈顶最近的位置(可能有多个相同的元素)
// 位置 1 代表栈顶，即最后一个元素。
// 位置 -1 代表没有找到该元素。
func (s *int64DESCStack) Search(n int64) int {
	lenS := len(s.items)
	for i := lenS - 1; i >= 0; i-- {
		if s.items[i] == n {
			return lenS - i
		}
	}

	return -1
}
