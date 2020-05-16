package tree

import (
	"errors"
)

// true - add to leftChild , false add to rightChild
func (n *node) addNewChild(value interface{}, order int, isLeftChild bool) {
	newChild := &node{
		parent: n,
		value:  value,
		order:  order,
		depth:  1,
		tree:   n.tree,
	}
	if isLeftChild {
		n.leftChild = newChild
	} else {
		n.rightChild = newChild
	}
}

// 判断自己是left child 还是 right child
func (n *node) isLeftChild() bool {
	// 内部使用，如果是 root 节点会 panic
	return n.order < n.parent.order
}

// 获取左右高度，计算左右高度差
func (n *node) calBalance() int {
	var lDep, rDep int
	if n.leftChild != nil {
		lDep = n.leftChild.depth
	}

	if n.rightChild != nil {
		rDep = n.rightChild.depth
	}

	return lDep - rDep
}

// 判断需要按照什么方式旋转, is Add Node 表示是 Add node OR delete node 时使用该方法
// 不同地方使用，可能会出现不同情况。
func (n *node) balanceFactor(isAddNode bool) error {
	// cal balance factor
	balanceFactor := n.calBalance()
	switch {
	case balanceFactor > 1: // 左边长
		if n.leftChild.calBalance() > 0 { // 右旋
			n.rightRightRotate()
		} else if n.leftChild.calBalance() < 0 { // 右旋左旋
			n.rightLeftRotate()
		} else {
			if isAddNode {
				return errors.New("balance factor err: the left Child is balanced")
			}
			n.rightRightRotate()
		}
	case balanceFactor < -1: // 右边长
		if n.rightChild.calBalance() < 0 { // 左旋
			n.leftLeftRotate()
		} else if n.rightChild.calBalance() > 0 { // 左旋右旋
			n.leftRightRotate()
		} else {
			if isAddNode {
				return errors.New("balance factor err: the right Child is balanced")
			}
			n.leftLeftRotate()
		}
	}
	return nil
}

// return nil means it is the root node
func (n *node) updateDepth() *node {
	var lDep, rDep int
	if n.leftChild != nil {
		lDep = n.leftChild.depth
	}

	if n.rightChild != nil {
		rDep = n.rightChild.depth
	}

	n.depth = max(lDep, rDep) + 1
	return n.parent
}

// return node's value
func (n *node) Value() interface{} {
	return n.value
}

// return node's order
func (n *node) Order() int {
	return n.order
}

// return node's depth
func (n *node) Depth() int {
	return n.depth
}

// could be nil
func (n *node) LeftChild() *node {
	return n.leftChild
}

// could be nil
func (n *node) RightChild() *node {
	return n.rightChild
}

// if node is root , it will return nil
func (n *node) Parent() *node {
	return n.parent
}

// node 所属 tree
func (n *node) Tree() *AVLTree {
	return n.tree
}

// left child -> right child -> right child -> right child...
func (n *node) LargestLeftTree() *node {
	var result *node
	for loop := n.leftChild; loop != nil; loop = loop.rightChild {
		result = loop
	}
	return result
}

// right child -> left child -> left child -> left child...
func (n *node) SmallestRightTree() *node {
	var result *node
	for loop := n.rightChild; loop != nil; loop = loop.leftChild {
		result = loop
	}
	return result
}
