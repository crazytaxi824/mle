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
		return RECOLOR // sibling is RED
	}
	return DONothing
}

// RR RL LL LR rotation
func (n *node) checkWhatKindRotation() byte {
	switch {
	case n.parent.isLeftChild() && n.isLeftChild():
		return LLRotation
	case !n.parent.isLeftChild() && !n.isLeftChild():
		return RRRotation
	case !n.parent.isLeftChild() && n.isLeftChild():
		return RLRotation
	default:
		return LRRotation
	}
}
