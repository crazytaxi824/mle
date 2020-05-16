package tree

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

// 右旋
func (n *node) rightRightRotate() {
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

// 右旋 -> 左旋
func (n *node) rightLeftRotate() {
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

	// 重新计算 left child 的深度
	newLeftChild.updateDepth()
}

// 左旋
func (n *node) leftLeftRotate() {
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

// 左旋 -> 右旋
func (n *node) leftRightRotate() {
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

	// 重新计算 right child 的深度
	newRightChild.updateDepth()
}

// R0 旋转
// node.balance > 1 && node.leftChild.balance == 0
func (n *node) r0r1Rotation() {
	grandFather := n.parent
	newParent := n.leftChild
	newRightChild := n
	oldRightChild := newParent.rightChild

	// n.depth--

	if grandFather == nil { // is the root
		n.tree.root = newParent
		n.tree.root.parent = nil
	} else {
		reBoundNodes(grandFather, newParent, n.isLeftChild())
	}

	reBoundNodes(newParent, newRightChild, false)
	reBoundNodes(newRightChild, oldRightChild, true)
}

// R1 旋转
// node.balance > 1 && node.leftChild.balance == 1

// R-1 旋转
// node.balance > 1 && node.leftChild.balance == -1

// L0 旋转

// L1 旋转

// L-1 旋转
