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

// return nil means it is the root node, or it needs to stop
func (n *node) updateDepth() *node {
	var lDep, rDep int
	if n.leftChild != nil {
		lDep = n.leftChild.depth
	}

	if n.rightChild != nil {
		rDep = n.rightChild.depth
	}

	if n.depth == max(lDep, rDep)+1 {
		return nil
	}

	n.depth = max(lDep, rDep) + 1
	return n.parent
}

// 检查树中的节点是否平衡
func (avl *AVLTree) checkBalances(_node *node, isAddNode bool) error {
	loop := _node
	for loop != nil { // TODO 可以优化不用一直检测到root
		// balance factor
		err := loop.balanceFactor(isAddNode)
		if err != nil {
			return err
		}

		loop = loop.updateDepth()
	}
	return nil
}
