package rbtree

func reBoundNodes(parent, child *node, isLeftChild bool) {
	if isLeftChild {
		parent.leftChild = child
	} else {
		parent.rightChild = child
	}

	if child != nil {
		child.parent = parent
	}
}

// LL
func (n *node) leftLeftRotate() {
	grandFather := n.parent
	newParent := n.leftChild
	newRightChild := n
	oldRightChild := newParent.rightChild

	if grandFather == nil { // is the root
		n.tree.root = newParent
		n.tree.root.parent = nil
	} else {
		reBoundNodes(grandFather, newParent, n.isLeftChild())
	}

	reBoundNodes(newParent, newRightChild, false)
	reBoundNodes(newRightChild, oldRightChild, true)
}

// LR
func (n *node) leftRightRotate() {
	grandFather := n.parent
	newParent := n.leftChild.rightChild
	newLeftChild := n.leftChild
	newRightChild := n
	oldRightChild := newParent.rightChild
	oldLeftChild := newParent.leftChild

	if grandFather == nil {
		n.tree.root = newParent
		n.tree.root.parent = nil
	} else {
		reBoundNodes(grandFather, newParent, n.isLeftChild())
	}

	reBoundNodes(newParent, newLeftChild, true)
	reBoundNodes(newParent, newRightChild, false)
	reBoundNodes(newRightChild, oldRightChild, true)
	reBoundNodes(newLeftChild, oldLeftChild, false)
}

// RR
func (n *node) rightRightRotate() {
	grandFather := n.parent
	newParent := n.rightChild
	newLeftChild := n
	oldLeftChild := newParent.leftChild

	if grandFather == nil {
		n.tree.root = newParent
		n.tree.root.parent = nil
	} else {
		reBoundNodes(grandFather, newParent, n.isLeftChild())
	}

	reBoundNodes(newParent, newLeftChild, true)
	reBoundNodes(newLeftChild, oldLeftChild, false)
}

// RL
func (n *node) rightLeftRotate() {
	grandFather := n.parent
	newParent := n.rightChild.leftChild
	newRightChild := n.rightChild
	newLeftChild := n
	oldLeftChild := newParent.leftChild
	oldRightChild := newParent.rightChild

	if grandFather == nil {
		n.tree.root = newParent
		n.tree.root.parent = nil
	} else {
		reBoundNodes(grandFather, newParent, n.isLeftChild())
	}

	reBoundNodes(newParent, newRightChild, false)
	reBoundNodes(newParent, newLeftChild, true)
	reBoundNodes(newLeftChild, oldLeftChild, false)
	reBoundNodes(newRightChild, oldRightChild, true)
}
