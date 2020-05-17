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
	case balanceFactor > 1 && n.leftChild.calBalance() > 0:
		n.leftLeftRotate()

	case balanceFactor > 1 && n.leftChild.calBalance() < 0:
		n.leftRightRotate()

	case balanceFactor > 1 && n.leftChild.calBalance() == 0:
		if isAddNode {
			return errors.New("balance factor err: the left Child is balanced")
		}
		n.leftLeftRotate()

	case balanceFactor < -1 && n.rightChild.calBalance() < 0:
		n.rightRightRotate()

	case balanceFactor < -1 && n.rightChild.calBalance() > 0:
		n.rightLeftRotate()

	case balanceFactor < -1 && n.rightChild.calBalance() == 0:
		if isAddNode {
			return errors.New("balance factor err: the right Child is balanced")
		}
		n.rightRightRotate()
	}
	return nil
}

// return nil means it is the root node, or it needs to stop
// 如果返回nil，说明该节点是 root，或者该节点不用再继续向上查找了
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
	for loop != nil { // 优化不用一直检测到root
		// balance factor
		err := loop.balanceFactor(isAddNode)
		if err != nil {
			return err
		}

		loop = loop.updateDepth()
	}
	return nil
}
