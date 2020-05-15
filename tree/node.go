package tree

import (
	"errors"
)

// true - add to leftChild , false add to rightChild
func (n *node) addNode(value interface{}, order int, isLeftChild bool) {
	newChild := &node{
		parent:      n,
		isLeftChild: isLeftChild,
		value:       value,
		order:       order,
		depth:       1,
		tree:        n.tree,
	}
	if isLeftChild {
		n.leftChild = newChild
	} else {
		n.rightChild = newChild
	}
}

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

func (n *node) balanceFactor() error {
	// cal balance factor
	balanceFactor := n.calBalance()
	switch {
	case balanceFactor > 1: // 左边长
		if n.leftChild.calBalance() > 0 { // 右旋
			n.rightRotate()
		} else if n.leftChild.calBalance() < 0 { // 右旋左旋
			n.rightLeftRotate()
		} else {
			return errors.New("some thing wrong")
		}
	case balanceFactor < -1: // 右边长
		if n.rightChild.calBalance() < 0 { // 左旋
			n.leftRotate()
		} else if n.rightChild.calBalance() > 0 { // 左旋右旋
			n.leftRightRotate()
		} else {
			return errors.New("some thing wrong")
		}
	}
	return nil
}

// return node indicate which node need to update its depth,
// return nil means the depth update action of the tree stops here
func (n *node) updateDepth() *node {
	var lDep, rDep int
	if n.leftChild != nil {
		lDep = n.leftChild.depth
	}

	if n.rightChild != nil {
		rDep = n.rightChild.depth
	}

	maxDepth := max(lDep, rDep) + 1

	if n.depth != maxDepth {
		n.depth = maxDepth
		return n.parent
	}
	return nil
}

func (n *node) Value() interface{} {
	return n.value
}

func (n *node) Order() int {
	return n.order
}

func (n *node) LeftChild() *node {
	return n.leftChild
}

func (n *node) RightChild() *node {
	return n.rightChild
}

func (n *node) Parent() *node {
	return n.parent
}

func (n *node) Tree() *AVLTree {
	return n.tree
}
