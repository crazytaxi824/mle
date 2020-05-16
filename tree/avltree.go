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
	tree                  *AVLTree    // 树
}

type AVLTree struct {
	root   *node
	length int
}

// duplicate value is not allowed here
func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

// add node
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
	return avl.checkBalances(parent, true)
}

// could be nil if the order is not exist
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

// delete node
func (avl *AVLTree) DeleteFromOrder(order int) error {
	delNode := avl.Find(order)
	if delNode == nil {
		return errors.New(NotExistNodeErr)
	}

	var parent *node

	switch {
	case delNode.leftChild != nil && delNode.rightChild != nil: // has both child
		// find replace
		replaceNode := delNode.LargestLeftTree()

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

		if parent == nil { // root
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
	return avl.checkBalances(parent, false)
}

func (avl *AVLTree) Delete(n *node) error {
	return avl.DeleteFromOrder(n.order)
}

// 树的容量
func (avl *AVLTree) Size() int {
	return avl.length
}

// 树得深度
func (avl *AVLTree) Depth() int {
	if avl.root == nil {
		return 0
	}
	return avl.root.depth
}

// 树的root节点
func (avl *AVLTree) Root() *node {
	return avl.root
}

// 最小的元素
func (avl *AVLTree) Smallest() *node {
	var smallest *node
	for loop := avl.root; loop != nil; loop = loop.leftChild {
		smallest = loop
	}
	return smallest
}

// 最大元素
func (avl *AVLTree) Biggest() *node {
	var biggest *node
	for loop := avl.root; loop != nil; loop = loop.rightChild {
		biggest = loop
	}
	return biggest
}
