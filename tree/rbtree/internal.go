package rbtree

// true - add to leftChild , false add to rightChild
func (n *node) addNewChild(value interface{}, order int, isLeftChild bool) *node {
	newNode := &node{
		parent: n,
		value:  value,
		order:  order,
		tree:   n.tree,
		color:  RED,
	}
	if isLeftChild {
		n.leftChild = newNode
	} else {
		n.rightChild = newNode
	}

	return newNode
}

// 判断自己是left child 还是 right child
func (n *node) isLeftChild() bool {
	// 内部使用，如果是 root 节点会 panic
	return n.order < n.parent.order
}

// sibling could be nil
func (n *node) sibling() *node {
	// 内部使用，如果是 root 节点会 panic
	if n.isLeftChild() {
		return n.parent.rightChild
	}
	return n.parent.leftChild
}

// check if the node is the root node
func (n *node) isRootNode() bool {
	return n.parent == nil
}

// re-color ? or rotation ?
func (n *node) checkWhatToDo() byte {
	if n.color == RED && n.parent.color == RED { // RR conflict
		if n.parent.sibling() == nil || n.parent.sibling().color == BLACK { // sibling is nil or BLACK
			return n.checkWhatKindRotation()
		}
		return reColor // sibling is RED
	}
	return doNothing
}

// RR RL LL LR rotation
func (n *node) checkWhatKindRotation() byte {
	switch {
	case n.parent.isLeftChild() && n.isLeftChild():
		return llRotation
	case !n.parent.isLeftChild() && !n.isLeftChild():
		return rrRotation
	case !n.parent.isLeftChild() && n.isLeftChild():
		return rlRotation
	default:
		return lrRotation
	}
}

// return predecessor Or Successor, if return nil means the node has no child
func (n *node) predecessorOrSuccessor() *node {
	if n.Predecessor() != nil {
		return n.Predecessor()
	}
	return n.Successor()
}

// both child are BLACK color, nil child is considered as BLACK color
func (n *node) bothChildrenAreBlack() bool {
	return (n.leftChild == nil || n.leftChild.color == BLACK) &&
		(n.rightChild == nil || n.rightChild.color == BLACK)
}

// opposite side child of the sibling
func (n *node) farSideOfTheNephew() *node {
	if n.isLeftChild() {
		return n.sibling().rightChild
	}
	return n.sibling().leftChild
}

// near side child of the sibling
func (n *node) nearSideOfTheNephew() *node {
	if n.isLeftChild() {
		return n.sibling().leftChild
	}
	return n.sibling().rightChild
}
