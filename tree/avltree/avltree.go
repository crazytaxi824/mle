package avltree

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
}

type tree struct {
	root   *node
	length int
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

	// 生成一个 node
	n := &node{
		value: value,
		index: index,
		depth: 1, // 新节点加入，depth 永远是 1.
		tree:  t,
	}
	// DEBUG testing node GC
	// runtime.SetFinalizer(n, func(p *node) {
	// 	log.Println(p.index, " is GC~~~~~~~~~~~~~~~~~~~~~")
	// })

	// 判断位置添加节点
	t.addNode(n)

	// 长度 +1
	t.length++

	// re-balance 节点
	loopCheckDepthAndRebalance(n.parent)

	return nil
}

// 将新的 node 加入 tree
func (t *tree) addNode(n *node) {
	// 第一个节点, root
	if t.root == nil {
		t.root = n
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
func (t *tree) Remove(index int64) error {
	// 找到 node
	node := t.getNode(index)
	if node == nil {
		return ErrNodeNotExist
	}

	// remove node
	check := t.removeNode(node)

	// 长度 -1
	t.length--

	// re-balance
	loopCheckDepthAndRebalance(check)

	return nil
}

// 返回一个 node 用于 check balance.
func (t *tree) removeNode(n *node) *node {
	var replacer *node
	// 根据 left, right child 的 depth 来选取 Predecessor || Successor
	// leftDepth > rightDepth 则选择 Predecessor
	// leftDepth < rightDepth 则选择 Successor
	// NOTE 这种选择方式主要是为了避免多次 rotation。
	leftDepth, rightDepth := n.childrensDepth()

	// 先判断自己是否是 leaf node
	if leftDepth == 0 && rightDepth == 0 { // leaf node
		parent := n.parent
		if parent != nil {
			// 删除自己
			if n.index < parent.index { // left child
				parent.leftChild = nil
			} else {
				parent.rightChild = nil
			}

			// need to recheck parent depth and re-balance
			return parent
		}

		// 如果没有 parent，leftChild 和 rightChild 说明
		// n 是 root 节点，而且是唯一节点。
		t.root = nil
		return nil
	} else if leftDepth-rightDepth < 0 {
		// leftDepth < rightDepth 则选择 Successor
		replacer = n.successor()
	} else {
		// leftDepth >= rightDepth 则选择 Predecessor
		replacer = n.predecessor()
	}

	// 这里是 replacer != nil 的情况。
	// leftDepth and rightDepth 其中有一个不是0，replace 不可能是 nil。

	// NOTE 这里顺序不能错，需要先解除关系然后再互换节点信息。
	// n 本身是 replacer.parent 的情况下，如果先互换节点信息的话，
	// 后面 index 信息是反的，会导致删除错误的 child。

	// NOTE predecessor/successor 自己可以是 leftChild 也可以是 rightChild.
	// predecessor/successor 可能会有 leftChild 或者 rightChild，
	// 如果有 child 需要和自己的 parent 连接。
	replacerParent := replacer.parent
	if replacer.index < replacerParent.index { // left child
		if replacer.leftChild != nil {
			bindingNodes(replacerParent, replacer.leftChild, isLeftChild)
		} else {
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
