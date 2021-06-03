package rbtree

func (t *tree) addNode(n *node) {
	if t.root == nil { // 添加第一个 node
		setRoot(n)
		return
	}

	var tmpParentNode *node // 临时储存 parent node
	var pos whichChild      // 判断是 leftChild 还是 rightChild

	// 步进方式对比节点
	for loop := t.root; loop != nil; {
		tmpParentNode = loop // 储存 parent

		if n.index < loop.index { // 向左走
			pos = isLeftChild
			loop = loop.leftChild
		} else if n.index > loop.index { // 向右走
			pos = isRightChild
			loop = loop.rightChild
		}
	}

	// 连接两个节点
	bindingNodes(tmpParentNode, n, pos)
}

func checkAndReBalance(n *node) {
	loop := n
	for loop != nil {
		loop = checkRedConflictAndRebalance(loop)
	}
}

// solve red-red conflict
func checkRedConflictAndRebalance(n *node) (next *node) {
	// NOTE 这里没考虑出现 nil 的情况。

	// 如果没有 red red conflict 则, 停止检查。
	if !checkRedConflict(n) {
		return nil // return nil to stop loop check
	}

	// red red conflict 解决方法
	// find parent's sibling
	parent := n.parent
	parentSibling := parent.sibling()
	grandparent := parent.parent

	// parent's sibling != nil && parent's sibling color == red,
	// recolor parent and parent's sibling, and
	// if parent's parent is not root, recolor it and recheck parent's parent.
	if parentSibling != nil && parentSibling.color == COLOR_RED {
		// recolor parent and parent's sibling,
		parent.reColor()
		parentSibling.reColor()

		// if parent's parent is NOT ROOT, recolor it and recheck parent's parent.
		if grandparent != n.tree.root {
			grandparent.reColor()
			return grandparent
		}

		// if parent's parent IS ROOT, DO NOT RECOLOR, and stop loop check.
		return nil // return nil to stop loop check
	}

	// parent's sibling is black(nil) - rotation & recolor
	// rotation 包含 n(self), n.parent, n.parent.parent 这几个节点。
	// NOTE 不同类型的 rotation 中，recolor 的 node 是不同的。
	// 判断 rotation 类型
	if n == n.parent.leftChild { // left child
		// NOTE recolor & rotation 顺序不能错，否则父子关系会出错。
		// NOTE the NODE which performs rotation is n.parent.parent.
		if n.parent == n.parent.parent.leftChild { // left child
			// recolor parent & parent's parent
			n.parent.reColor()
			n.parent.parent.reColor()
			// LL rotation
			n.parent.parent.leftLeftRotation()
		} else { // right child
			// recolor self & parent's parent
			n.reColor()
			n.parent.parent.reColor()
			// RL rotation
			n.parent.parent.rightLeftRotation()
		}
	} else { // right child
		if n.parent == n.parent.parent.leftChild { // left child
			// recolor self & parent's parent
			n.reColor()
			n.parent.parent.reColor()
			// LR rotation
			n.parent.parent.leftRightRotation()
		} else { // right child
			// recolor parent & parent's parent
			n.parent.reColor()
			n.parent.parent.reColor()
			// RR rotation
			n.parent.parent.rightRightRotation()
		}
	}

	// NOTE 完成 recolor & rotation 之后不用再 loop check.
	return nil
}

// red red conflict - node and it's parent are both red.
// true = red red conflict; false = no conflict.
func checkRedConflict(n *node) bool {
	if n.parent == nil { // root node
		return false
	}

	if n.color == COLOR_RED && n.parent.color == COLOR_RED {
		return true
	}
	return false
}
