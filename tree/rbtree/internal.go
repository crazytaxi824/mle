package rbtree

// 判断自己是哪一边的 child
func (n *node) whichPos() whichChild {
	if n == n.parent.leftChild {
		return isLeftChild
	}
	return isRightChild
}

// change black -> red, red -> black
func (n *node) reColor() {
	n.color = !n.color
}

// swap the color of two nodes
func (n *node) swapColor(s *node) {
	n.color, s.color = s.color, n.color
}

// sibling could be nil
func (n *node) sibling() *node {
	// 判断 left right child
	// left child
	if n == n.parent.leftChild {
		return n.parent.rightChild
	}

	// right child
	return n.parent.leftChild
}

// NOTE BST 中，永远不要删除 internal node，
// 找到替代的 leaf node 进行替换，然后删除 leaf node。
// 替换 n & replacer 的 index 和 value，不替换 color, parent, child 等信息。
func (n *node) replaceNode(replacer *node) {
	n.index, replacer.index = replacer.index, n.index
	n.value = replacer.value
}

// 用于 Sort
func (n *node) findFirstRightSideParent() *node {
	// 从 n 开始向 parent 步进
	for loop := n; ; loop = loop.parent {
		// 一直找到 root 的情况, 这时是没有 first right side parent 的, 返回 nil
		if loop.parent == nil {
			return nil
		}

		if loop == loop.parent.leftChild { // left child
			return loop.parent
		}
	}
}
