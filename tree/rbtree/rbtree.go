/* 红黑树要点
1. root 一定是 black.
2. node 和 node.parent 不能同时是 red。 red-red conflict.
3. 红黑树每个分支上的 black node 数量相等。如果一个 leaf node 如果是 black，
那么他一定会有 sibling，否则无法满足该条件。只有 red leaf node 可能没有 sibling。

红黑树特点：
1. 在红黑树中，任意节点两边的高度相差不会超过1倍。
即：最短分支的长度 h*2 >= 最长分支长度 H。如果一个 node 一侧没有 child，
那么他这一侧分支的长度为 1(自己)，那么另一侧最多只能有长度为 2 的分支(包括
自己)，所以另一侧最多只能有一个 node(child). 而且这个 node 一定是 leaf。
如果一个 node 一侧有 1 个 node(child)，那么这一侧的分支长度为2(包括自己)，
所以另一侧可以有最多长度为 4 的分支(包括自己)。
*/

package rbtree

type Tree interface {
	// return root node, if the tree
	// is empty, it will return nil.
	Root() Node

	// return node by index, if index
	// do not exist, it will return nil.
	GetNode(index int64) Node

	// add a new node by adding index and value.
	// if index already exists, it will return err.
	Add(int64, interface{}) error

	// remove node by index, if index
	// do not exist, it will return err.
	Remove(index int64) error

	// return node of smallest index.
	// If the tree is empty, it will return nil.
	Smallest() Node

	// return node of largest index.
	// If the tree is empty, it will return nil.
	Largest() Node

	// return number of nodes of this tree, if
	// Size == 0 means this is an empty tree.
	Size() int

	// return a list of ASC nodes
	Sort() []Node
}

type tree struct {
	root         *node
	length       int
	cacheDelNode *node
}

func NewTree() Tree {
	return &tree{}
}

func (t *tree) Root() Node {
	r := t.root
	if r == nil {
		return nil
	}
	return r
}

func (t *tree) GetNode(index int64) Node {
	node := t.getNode(index)
	if node == nil {
		return nil
	}

	return node
}

func (t *tree) getNode(index int64) *node {
	loop := t.root
	if loop == nil {
		return nil
	}

	for loop != nil {
		if index == loop.index {
			return loop
		} else if index < loop.index { // left side
			loop = loop.leftChild
		} else { // right side
			loop = loop.rightChild
		}
	}

	return nil
}

func (t *tree) Add(index int64, value interface{}) error {
	// 先检查 index 是否存在，避免分配内存
	if t.getNode(index) != nil {
		return ErrIndexExist
	}

	n := &node{
		index: index,
		value: value,
		// color: red, // NOTE new node always red color, except root.
		tree: t,
	}
	// DEBUG testing node GC
	// runtime.SetFinalizer(n, func(p *node) {
	// 	log.Println(p.index, " is GC~~~~~~~~~~~~~~~~~~~~~")
	// })

	// 判断位置添加节点
	t.addNode(n)

	t.length++

	// 检查颜色冲突
	checkAndReBalance(n)

	return nil
}

func (t *tree) Remove(index int64) error {
	// 找到 node
	node := t.getNode(index)
	if node == nil {
		return ErrNodeNotExist
	}

	// NOTE 标记需要 remove 的 node, 返回 double black node.
	// 这里不直接删除 node 是为了保持 nodes 之间的关系便于解决 DB node 问题.
	dbNode := t.markNodes(node)

	// 解决 double black node 情况.
	resolveDoubleBlack(dbNode)

	// 实际删除被缓存的 node.
	t.deleteCachedNode()

	t.length--

	return nil
}

func (t *tree) Smallest() Node {
	small := t.smallest()
	if small == nil {
		return nil
	}
	return small
}

func (t *tree) smallest() *node {
	var small *node
	for loop := t.root; loop != nil; loop = loop.leftChild {
		small = loop
	}
	return small
}

func (t *tree) Largest() Node {
	large := t.largest()
	if large == nil {
		return nil
	}
	return large
}

func (t *tree) largest() *node {
	var large *node
	for loop := t.root; loop != nil; loop = loop.rightChild {
		large = loop
	}
	return large
}

func (t *tree) Size() int {
	return t.length
}

// 先从 Smallest 开始，检查自己有没有 Successor。如果有，move 到 Successor。
// 如果没有 Successor 则寻找第一个 right side parent，即自己是 parent 的 leftChild。
// 如果有 first right side parent 则 move 到 first right side parent,
// 如果 first right side parent 不存在，说明已经到 root 了，结束循环。
func (t *tree) Sort() []Node {
	var result []Node
	loop := t.smallest()
	for loop != nil {
		result = append(result, loop)

		// 先判断 Successor
		if loop.successor() != nil {
			// 如果 Successor 存在则 move 到 Successor
			loop = loop.successor()
		} else {
			// 如果 Successor 不存在则寻找 first Right parent
			loop = loop.findFirstRightSideParent()
		}
	}
	return result
}
