/* avltree 要点
1. avltree 任意节点两边的高度(depth)差不能超过1(可以是1)。如果某个 node 一侧没有 child，
那么这一侧的高度为0，另一侧的高度最多为 1。如果高度差超过 1 需要进行 rotation 来平衡。
2. 自身高度 node.depth = max(leftChild.depth, rightChild.depth) + 1。
*/

package avltree

import (
	"log"
	"runtime"
)

type Tree interface {
	// return root node, if the tree
	// is empty, it will return nil.
	Root() Node

	// return node by index, if index
	// do not exist, it will return nil.
	Search(index int64) Node

	// add a new node by adding index and value.
	// if index already exists, it will return err.
	Insert(int64, interface{}) error

	// remove node by index, if index
	// do not exist, it will return err.
	Delete(index int64) error

	// return tree max depth, if depth == 0
	// means this is an empty tree.
	Depth() int

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

	// Clear the whole tree. delete every node in the tree.
	// NOTE
	// if just set tree.root = nil, GC may not release memory.
	// if you want to release memory, you have to Clear() after
	// using the tree.
	Clear()
}

type tree struct {
	root   *node
	length int
}

func NewTree() Tree {
	return &tree{}
}

// internal use only
func newTree() *tree {
	return &tree{}
}

func (t *tree) Root() Node {
	r := t.root
	if r == nil {
		return nil
	}
	return r
}

func (t *tree) Search(index int64) Node {
	node := t.search(index)
	if node == nil {
		return nil
	}

	return node
}

func (t *tree) search(index int64) *node {
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

func (t *tree) Insert(index int64, value interface{}) error {
	// 先检查 index 是否存在，避免分配内存
	if t.search(index) != nil {
		return ErrIndexExist
	}

	// 生成一个 node
	n := &node{
		value: value,
		index: index,
		depth: 1, // 新节点加入，depth 永远是 1.
		tree:  t,
	}
	// DEBUG testing node GC
	runtime.SetFinalizer(n, func(p *node) {
		log.Println(p.index, " is GC~~~~~~~~~~~~~~~~~~~~~")
	})

	// 判断位置添加节点
	t.addNode(n)

	// 长度 +1
	t.length++

	// re-balance 节点
	checkAndReBalance(n.parent)

	return nil
}

// 将新的 node 加入 tree
func (t *tree) addNode(n *node) {
	if t.root == nil { // 第一个节点, root
		setRoot(n)
		return
	}

	var tmpParentNode *node // 临时储存 parent node
	var pos whichChild      // 判断是 leftChild 还是 rightChild

	// 步进方式对比节点
	for compareNode := t.root; compareNode != nil; {
		tmpParentNode = compareNode // 储存 parent

		if n.index < compareNode.index { // 向左走
			pos = isLeftChild
			compareNode = compareNode.leftChild
		} else if n.index > compareNode.index { // 向右走
			pos = isRightChild
			compareNode = compareNode.rightChild
		}
	}

	// 连接两个节点
	bindingNodes(tmpParentNode, n, pos)
}

// NOTE BST 中，永远不要删除 internal node，
// 找到替代的 leaf node 进行替换，然后删除 leaf node。
func (t *tree) Delete(index int64) error {
	// 找到 node
	node := t.search(index)
	if node == nil {
		return ErrNodeNotExist
	}

	// remove node
	check := t.removeNode(node)

	// 长度 -1
	t.length--

	// re-balance
	checkAndReBalance(check)

	return nil
}

// 返回一个 node 用于 check balance.
func (t *tree) removeNode(n *node) *node {
	leftDepth, rightDepth := n.childrensDepth()

	// 需要删除的 node 是 leaf node 的情况。
	if leftDepth == 0 && rightDepth == 0 { // leaf node
		parent := n.parent
		// 如果没有 parent，leftChild 和 rightChild 说明
		// n 是 root 节点，而且是唯一节点。
		if parent == nil {
			// 直接删除
			n.delete()
			t.root = nil
			return nil
		}

		// parent != nil 的情况
		// 直接删除自己
		if n == parent.leftChild { // left child
			parent.leftChild = nil
		} else { // right child
			parent.rightChild = nil
		}

		// 直接删除
		n.delete()

		// need to recheck parent depth and re-balance
		return parent
	}

	// 需要删除的 node 不是 leaf node 的情况。
	// 根据 left, right child 的 depth 来选取 Predecessor || Successor
	// leftDepth > rightDepth 则选择 Predecessor
	// leftDepth < rightDepth 则选择 Successor
	// NOTE 这种选择方式主要是为了避免多次 rotation。
	var replacer *node
	if leftDepth-rightDepth < 0 {
		// leftDepth < rightDepth 则选择 Successor
		replacer = n.successor()
	} else {
		// leftDepth >= rightDepth 则选择 Predecessor
		replacer = n.predecessor()
	}

	// 这里是 replacer != nil 的情况。
	// leftDepth and rightDepth 其中有一个不是0，replacer 不可能是 nil。

	// NOTE predecessor/successor 自己可以是 leftChild 也可以是 rightChild.
	// predecessor 可能会有 leftChild; successor 可能会有 rightChild.
	// 如果 replacer 有 child 需要和 replacer.parent 连接。
	replacerParent := replacer.parent
	if replacer == replacerParent.leftChild { // left child
		if replacer.leftChild != nil {
			// 自己是 left child，同时有 left child，说明 replacer 是 predecessor
			bindingNodes(replacerParent, replacer.leftChild, isLeftChild)
		} else {
			// 自己是 left child，同时没有 left child，replacer 可能是 predecessor, successor.
			bindingNodes(replacerParent, replacer.rightChild, isLeftChild)
		}
	} else { // right child
		if replacer.leftChild != nil {
			bindingNodes(replacerParent, replacer.leftChild, isRightChild)
		} else {
			bindingNodes(replacerParent, replacer.rightChild, isRightChild)
		}
	}

	// 替换 n & replacer 的 index 和 value，不替换 depth, parent, child 等信息。
	n.replaceNode(replacer)

	// 删除 replacer
	replacer.delete()

	// need to recheck replacerParent depth and rebalance
	return replacerParent
}

func (t *tree) Depth() int {
	return t.root.depth
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

// NOTE delete every node in the tree.
func (t *tree) Clear() {
	loop := []*node{t.root}
	for loop != nil {
		var tmp []*node
		for _, n := range loop {
			if n.leftChild != nil {
				tmp = append(tmp, n.leftChild)
			}
			if n.rightChild != nil {
				tmp = append(tmp, n.rightChild)
			}
			n.delete()
		}
		loop = tmp
	}

	t.root = nil
	t.length = 0
}
