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

	n.depth--

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

	newParent.depth++
	n.depth--
	n.leftChild.depth--

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

// 左旋
func (n *node) leftLeftRotate() {
	grandFather := n.parent
	newParent := n.rightChild
	newLeftChild := n
	oldLeftChild := newParent.leftChild

	n.depth--

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

	newParent.depth++
	n.depth--
	n.rightChild.depth--

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
