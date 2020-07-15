// 普通栈，单调递增栈，单调递减栈
package stack

type int16Stack struct {
	items []int16

	// if its true then allow duplicated elements in the stack
	allowDupl bool

	// Normal/ASC/DESC 普通/单调递增/单调递减
	stackType uint8
}

func NewInt16Stack(opt *Option) *int16Stack {
	if opt == nil {
		opt = &Option{
			AllowDuplicate: true,
			StackType:      NormalStack,
		}
	}

	return &int16Stack{allowDupl: opt.AllowDuplicate, stackType: opt.StackType}
}

// 单调栈，会踢出大于自己的元素
// n 表示有多少元素被踢出
// return n means how many items has been removed from stack
func (s *int16Stack) Push(i int16) (n int, res []int16) {
	switch s.stackType {
	case ASCStack:
		return s.pushToASCStack(i)

	case DESCStack:
		return s.pushToDESCStack(i)
	}

	return s.pushToNormalStack(i)
}

func (s *int16Stack) pushToASCStack(i int16) (n int, res []int16) {
	index := -1

	if s.allowDupl {
		for k := range s.items {
			if s.items[k] > i {
				index = k
				break
			}
		}
	} else {
		for k := range s.items {
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

func (s *int16Stack) pushToDESCStack(i int16) (n int, res []int16) {
	index := -1

	if s.allowDupl {
		for k := range s.items {
			if s.items[k] < i {
				index = k
				break
			}
		}
	} else {
		for k := range s.items {
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

func (s *int16Stack) pushToNormalStack(i int16) (n int, res []int16) {
	if !s.allowDupl {
		for k := range s.items {
			if s.items[k] == i {
				return 1, []int16{i}
			}
		}
	}

	s.items = append(s.items, i)
	return 0, nil
}

// 返回 stack 长度
func (s *int16Stack) Len() int {
	return len(s.items)
}

// this is for test only
func (s *int16Stack) elements() []int16 {
	return s.items
}

// 后进先出，返回栈顶元素
func (s *int16Stack) Pop() (int16, error) {
	if len(s.items) == 0 {
		return 0, ErrEmptyStack
	}

	res := s.items[len(s.items)-1]
	s.items = s.items[: len(s.items)-1 : len(s.items)-1]
	return res, nil
}

// 返回栈顶元素，但是不删除该元素
func (s *int16Stack) Peek() (int16, error) {
	if len(s.items) == 0 {
		return 0, ErrEmptyStack
	}

	return s.items[len(s.items)-1], nil
}

// 重置 stack
func (s *int16Stack) Reset() {
	s.items = make([]int16, 0)
}

// 判断 stack 是否为空
func (s *int16Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// 返回对象距离栈顶最近的位置(可能有多个相同的元素)
// 位置 1 代表栈顶，即最后一个元素。
// 位置 -1 代表没有找到该元素。
func (s *int16Stack) Search(n int16) int {
	switch {
	case s.allowDupl == false && s.stackType == ASCStack:
		return s.searchASCDichotomy(n)
	case s.allowDupl == false && s.stackType == DESCStack:
		return s.searchDESCDichotomy(n)
	}
	return s.searchNormalStack(n)
}

// 二分法查找
func (s *int16Stack) searchASCDichotomy(n int16) int {
	lenS := len(s.items)

	// [start, end)
	for start, end := 0, lenS; start < end; {
		index := (start + end) / 2
		switch {
		case s.items[index] == n:
			// 找到之后遍历相同元素，allow duplicated element
			for index < lenS && s.items[index] == n {
				index++
			}
			return lenS - index + 1
		case s.items[index] > n:
			end = index
		case s.items[index] < n:
			start = index + 1
		}
	}
	return -1
}

func (s *int16Stack) searchDESCDichotomy(n int16) int {
	lenS := len(s.items)

	// [start, end)
	for start, end := 0, lenS; start < end; {
		index := (start + end) / 2
		switch {
		case s.items[index] == n:
			// 找到之后遍历相同元素，allow duplicated element
			for index < lenS && s.items[index] == n {
				index++
			}
			return lenS - index + 1
		case s.items[index] < n:
			end = index
		case s.items[index] > n:
			start = index + 1
		}
	}
	return -1
}

func (s *int16Stack) searchNormalStack(n int16) int {
	lenS := len(s.items)
	for i := lenS - 1; i >= 0; i-- {
		if s.items[i] == n {
			return lenS - i
		}
	}

	return -1
}

// for range element,
// if fn return false, stop range.
func (s *int16Stack) Range(fn func(element int16) bool) {
	for k := range s.items {
		if !fn(s.items[k]) {
			return
		}
	}
}

// return elements by index, start from 0
func (s *int16Stack) Index(index int) (int16, error) {
	if index >= len(s.items) {
		return 0, ErrOutOfRange
	}
	return s.items[index], nil
}
