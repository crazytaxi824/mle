package avltree

// 循环向上(parent)检查 depth 和 balance 信息。
func checkAndReBalance(n *node) {
	loop := n
	for loop != nil {
		// 循环检查直到 root
		loop = checkNodeDepthAndRebalance(loop)
	}
}

// 重新计算 depth，同时判断是否 balance，如果 unbalance，需要用哪一种 rotation
func checkNodeDepthAndRebalance(n *node) (next *node) {
	// NOTE 这里没考虑出现 nil 的情况。

	// 获取 left, right Child depth
	leftDep, rightDep := n.childrensDepth()

	// 重新计算自己的 depth
	n.depth = max(leftDep, rightDep) + 1

	// NOTE check balance
	// 如果 leftDep - rightDep != {-1, 0, 1}, 则 node unbalance.
	if leftDep-rightDep < -1 { // R-unbalance
		// 判断 rotation
		childLeftDep, childRightDep := n.rightChild.childrensDepth()
		diff := childLeftDep - childRightDep

		// RR rotation,
		// NOTE 删除节点时会出现两边一样长(diff==0)的情况，使用 RR rotation。
		if diff <= 0 {
			// return next node that need to adjust depth and re-check balance
			return n.rightRightRotation()
		}

		// RL rotation, diff > 0 的情况
		// NOTE 任何情况下，只要进行 LR/RL rotation，n 的 depth 一定会 -2。
		// 其他的节点的 depth 在 add / remove 的不同情况下会发生变化。
		// 如果这里不调整 n.depth 的话，需要返回两个节点(n, rightChild) 用于重新检查。
		n.depth -= 2
		// return next node that need to adjust depth and re-check balance
		return n.rightLeftRotation()
	} else if leftDep-rightDep > 1 { // L-unbalance
		// 判断 rotation
		childLeftDep, childRightDep := n.leftChild.childrensDepth()
		diff := childLeftDep - childRightDep

		// LL rotation
		// NOTE 删除节点时会出现两边一样长(diff==0)的情况，使用 LL rotation。
		if diff >= 0 {
			// return next node that need to adjust depth and re-check balance
			return n.leftLeftRotation()
		}

		// LR rotation, diff < 0 的情况
		// NOTE 任何情况下，只要进行 LR/RL rotation，n 的 depth 一定会 -2。
		// 其他的节点的 depth 在 add / remove 的不同情况下会发生变化。
		// 如果这里不调整 n.depth 的话，需要返回两个节点(n, leftChild) 用于重新检查。
		n.depth -= 2
		// return next node that need to adjust depth and re-check balance
		return n.leftRightRotation()
	}

	// leftDep-rightDep >= -1 && leftDep-rightDep <= 1
	// 如果节点是 balanced, 继续向上(parent)检查。
	return n.parent
}

func (n *node) whichPos() whichChild {
	if n == n.parent.leftChild {
		return isLeftChild
	}
	return isRightChild
}

// 返回 left, right child 的 depth
func (n *node) childrensDepth() (leftDepth, rightDepth int) {
	if n.leftChild != nil {
		leftDepth = n.leftChild.depth
	}

	if n.rightChild != nil {
		rightDepth = n.rightChild.depth
	}

	return leftDepth, rightDepth
}

// NOTE BST 中，永远不要删除 internal node，
// 找到替代的 leaf node 进行替换，然后删除 leaf node。
// 替换 n & replacer 的 index 和 value，不替换 depth, parent, child 等信息。
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
