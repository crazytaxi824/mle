package tree

import (
	"errors"
)

const (
	ExistNodeErr    = "the node is already exist"
	NotExistNodeErr = "the node is not in the tree"
)

type node struct {
	parent                *node       // 上级节点
	leftChild, rightChild *node       // 左右节点
	depth                 int         // 自己的深度，最下层默认为1
	value                 interface{} // 内容
	order                 int         // 排序号码
	tree                  *AVLTree    // 所属树
}

type AVLTree struct {
	root   *node
	length int
}

// duplicate order number is not allowed here
func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

// add node to the tree
func (avl *AVLTree) Add(order int, value interface{}) error {
	// 添加第一个节点
	if avl.root == nil {
		avl.root = &node{
			parent:     nil,
			leftChild:  nil,
			rightChild: nil,
			depth:      1,
			value:      value,
			order:      order,
			tree:       avl,
		}
		avl.length = 1
		return nil
	}

	// whose child
	parent, isLeftChild, err := avl.whoseChild(order)
	if err != nil {
		return err
	}

	// add node
	parent.addNewChild(value, order, isLeftChild)
	avl.length++

	// 计算是否需要 Re-balance
	avl.checkBalances(parent)

	return nil
}

// true - add to leftChild , false add to rightChild
func (avl *AVLTree) whoseChild(order int) (*node, bool, error) {
	var result *node
	var isLeftNode bool

	for loop := avl.root; loop != nil; {
		if order == loop.order {
			return nil, false, errors.New(ExistNodeErr)
		}

		if order < loop.order { // left
			result = loop
			loop = loop.leftChild
			isLeftNode = true
		} else { // right
			result = loop
			loop = loop.rightChild
			isLeftNode = false
		}
	}
	return result, isLeftNode, nil
}

// find node from order number, could be nil if the order is not exist
func (avl *AVLTree) Find(order int) *node {
	var result *node

	for result = avl.root; result != nil && result.order != order; {
		if order < result.order { // 左边
			result = result.leftChild
		} else if order > result.order { // 右边
			result = result.rightChild
		}
	}

	return result
}

// delete node from order number
func (avl *AVLTree) DeleteFromOrder(order int) error {
	delNode := avl.Find(order)
	if delNode == nil {
		return errors.New(NotExistNodeErr)
	}

	var parent *node

	switch {
	case delNode.leftChild != nil && delNode.rightChild != nil: // has both child
		// find replace
		replaceNode := delNode.Predecessor()

		// 删除 replaceNode
		parent = replaceNode.parent
		if replaceNode.isLeftChild() {
			parent.leftChild = nil
		} else {
			parent.rightChild = nil
		}

		// 替换value，order，不替换 depth，left，right child
		delNode.order = replaceNode.order
		delNode.value = replaceNode.value

		replaceNode = nil

	default:
		parent = delNode.parent

		if parent == nil { // delNode is root
			avl.root = nil
			delNode = nil
			return nil
		}

		if delNode.leftChild != nil {
			reBoundNodes(parent, delNode.leftChild, delNode.isLeftChild())
		} else if delNode.rightChild != nil {
			reBoundNodes(parent, delNode.rightChild, delNode.isLeftChild())
		} else {
			// 删除自己
			if delNode.isLeftChild() {
				parent.leftChild = nil
			} else {
				parent.rightChild = nil
			}
		}

		delNode = nil
	}

	avl.length--

	// 计算是否需要 Re-balance
	avl.checkBalances(parent)

	return nil
}

// delete node
func (avl *AVLTree) Delete(n *node) error {
	return avl.DeleteFromOrder(n.order)
}

// total number of nodes in the tree
func (avl *AVLTree) Size() int {
	return avl.length
}

// total depth of the tree
func (avl *AVLTree) Depth() int {
	if avl.root == nil {
		return 0
	}
	return avl.root.depth
}

// root node of the tree
func (avl *AVLTree) Root() *node {
	return avl.root
}

// smallest node in the tree
func (avl *AVLTree) Smallest() *node {
	var smallest *node
	for loop := avl.root; loop != nil; loop = loop.leftChild {
		smallest = loop
	}
	return smallest
}

// biggest node in the tree
func (avl *AVLTree) Biggest() *node {
	var biggest *node
	for loop := avl.root; loop != nil; loop = loop.rightChild {
		biggest = loop
	}
	return biggest
}

// sort the nodes in ASC order
func (avl *AVLTree) Sort() []*node {
	result := make([]*node, 0, avl.length)
	smallest := avl.Smallest()

	// s -> small right tree
	for loop := smallest; loop != nil; {
		result = append(result, loop)
		if loop.Successor() != nil { // 先找自己 smallest right
			loop = loop.Successor()
		} else { // 再找 parent
			if loop.parent != nil && loop.order < loop.parent.order { // 在左边
				loop = loop.parent
			} else if loop.parent != nil && loop.order >= loop.parent.order {
				// 如果自己是 right child 继续向上找一直到 left child
				loop = loop.findLeftParent()
			} else {
				break
			}
		}
	}

	return result
}

// for sort，一直寻找 parent，直到自己是 parent 的 left child，返回 parent，
// 如果自己是 parent 的 right child，继续向上寻找。
func (n *node) findLeftParent() *node {
	for loop := n; ; {
		if loop.parent == nil {
			return nil
		}
		if loop.order >= loop.parent.order {
			loop = loop.parent
		} else {
			return loop.parent
		}
	}
}
