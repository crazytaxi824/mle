package avltree

// 这里有4个node，n(unbalanced node), n.parent, n.rightChild, n.right.LeftChild
// 操作如下：
// n.parent(需要判断left/right) -> n.rightChild
// n.rightChild(leftChild) -> n
// n(rightChild) -> n.right.LeftChild(可为nil)
// 返回一个 node 用于 check balance.
func (n *node) rightRightRotation() *node {
	parent := n.parent
	rightChild := n.rightChild
	rightLeftChild := n.rightChild.leftChild

	if parent == nil { // 自己是 root 的情况
		setRoot(rightChild)
	} else {
		// n.parent(需要判断left/right) -> n.rightChild
		// 判断自己(n)是(parent) 的 leftChild 还是 rightChild
		if n.index < parent.index { // leftChild
			bindingNodes(parent, rightChild, isLeftChild)
		} else { // rightChild
			bindingNodes(parent, rightChild, isRightChild)
		}
	}

	// n.rightChild(leftChild) -> n
	bindingNodes(rightChild, n, isLeftChild)

	// n(rightChild) -> n.right.LeftChild(可为nil)
	bindingNodes(n, rightLeftChild, isRightChild)

	// need to recheck n balance, after change it's relationship
	// 节点 n 的 depth 有多种可能，需要重新计算 depth，然后 check balance。
	return n
}

// 这里有4个node，n(unbalanced node), n.parent, n.leftChild, n.left.RightChild
// 操作如下：
// n.parent(需要判断left/right) -> n.leftChild
// n.leftChild(rightChild) -> n
// n(leftChild) -> n.left.rightChild(可为nil)
// 返回一个 node 用于 check balance.
func (n *node) leftLeftRotation() *node {
	parent := n.parent
	leftChild := n.leftChild
	leftRightChild := n.leftChild.rightChild

	if parent == nil { // 自己是 root 的情况
		setRoot(leftChild)
	} else {
		// n.parent(需要判断left/right) -> n.leftChild
		// 判断自己(n)是(parent) 的 leftChild 还是 rightChild
		if n.index < parent.index { // leftChild
			bindingNodes(parent, leftChild, isLeftChild)
		} else { // rightChild
			bindingNodes(parent, leftChild, isRightChild)
		}
	}

	// n.leftChild(rightChild) -> n
	bindingNodes(leftChild, n, isRightChild)

	// n(leftChild) -> n.left.rightChild(可为nil)
	bindingNodes(n, leftRightChild, isLeftChild)

	// need to recheck n balance, after change it's relationship
	// 节点 n 的 depth 有多种可能，需要重新计算 depth，然后 check balance。
	return n
}

// 6 个 node 需要操作：n(unbalanced node), n.parent, n.leftChild,
// n.left.rightChild, n.left.right.rightChild, n.left.right.leftChild
// 操作如下：
// n.parent(需要判断left/right) -> n.left.RightChild
// n.left.RightChild(rightChild) -> n
// n.left.RightChild(leftChild) -> n.leftChild
// n(leftChild) -> n.left.right.rightChild(可为nil)
// n.leftChild(rightChild) -> n.left.right.leftChild(可为nil)
// 返回一个 node 用于 check balance.
func (n *node) leftRightRotation() *node {
	parent := n.parent
	leftChild := n.leftChild
	leftRightChild := n.leftChild.rightChild // main
	leftRightRightChild := n.leftChild.rightChild.rightChild
	leftRightLeftChild := n.leftChild.rightChild.leftChild

	if parent == nil { // 自己是 root 的情况
		setRoot(leftRightChild)
	} else {
		// n.parent(需要判断left/right) -> n.left.RightChild
		// 判断自己(n)是(parent) 的 leftChild 还是 rightChild
		if n.index < parent.index { // leftChild
			bindingNodes(parent, leftRightChild, isLeftChild)
		} else { // rightChild
			bindingNodes(parent, leftRightChild, isRightChild)
		}
	}

	// n.left.RightChild(rightChild) -> n
	bindingNodes(leftRightChild, n, isRightChild)

	// n.left.RightChild(leftChild) -> n.leftChild
	bindingNodes(leftRightChild, leftChild, isLeftChild)

	// n(leftChild) -> n.left.right.rightChild(可为nil)
	bindingNodes(n, leftRightRightChild, isLeftChild)

	// n.leftChild(rightChild) -> n.left.right.leftChild(可为nil)
	bindingNodes(leftChild, leftRightLeftChild, isRightChild)

	// need to recheck leftChild balance, after change it's relationship
	// 节点 leftChild 的 depth 有多种可能，需要重新计算 depth，然后 check balance。
	return leftChild
}

// 6 个 node 需要操作：n(unbalanced node), n.parent, n.rightChild,
// n.right.leftChild, n.right.left.rightChild, n.right.left.leftChild
// 操作如下：
// n.parent(需要判断left/right) -> n.right.leftChild
// n.right.leftChild(rightChild) -> n.rightChild
// n.right.leftChild(leftChild) -> n
// n(rightChild) -> n.right.left.leftChild(可为nil)
// n.rightChild(leftChild) -> n.right.left.rightChild(可为nil)
// 返回一个 node 用于 check balance.
func (n *node) rightLeftRotation() *node {
	parent := n.parent
	rightChild := n.rightChild
	rightLeftChild := n.rightChild.leftChild // main
	rightLeftRightChild := n.rightChild.leftChild.rightChild
	rightLeftLeftChild := n.rightChild.leftChild.leftChild

	if parent == nil { // 自己是 root 的情况
		setRoot(rightLeftChild)
	} else {
		// n.parent(需要判断left/right) -> n.right.leftChild
		// 判断自己(n)是(parent) 的 leftChild 还是 rightChild
		if n.index < parent.index { // leftChild
			bindingNodes(parent, rightLeftChild, isLeftChild)
		} else { // rightChild
			bindingNodes(parent, rightLeftChild, isRightChild)
		}
	}

	// n.right.leftChild(leftChild) -> n
	bindingNodes(rightLeftChild, n, isLeftChild)

	// n.right.leftChild(rightChild) -> n.rightChild
	bindingNodes(rightLeftChild, rightChild, isRightChild)

	// n(rightChild) -> n.right.left.leftChild(可为nil)
	bindingNodes(n, rightLeftLeftChild, isRightChild)

	// n.rightChild(leftChild) -> n.right.left.rightChild(可为nil)
	bindingNodes(rightChild, rightLeftRightChild, isLeftChild)

	// need to recheck rightChild balance, after change it's relationship
	// 节点 rightChild 的 depth 有多种可能，需要重新计算 depth，然后 check balance。
	return rightChild
}

// 绑定两个节点关系
func bindingNodes(parent, child *node, isLeftChild whichChild) {
	if isLeftChild {
		parent.leftChild = child
	} else {
		parent.rightChild = child
	}

	// 如果 child == nil, 则不进行这一步
	if child != nil {
		child.parent = parent
	}
}

// 设置 root 节点到 n，同时将 n 的 parent 设置为 nil。
func setRoot(n *node) {
	n.tree.root = n
	n.parent = nil
}
