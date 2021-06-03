package rbtree

// mark deleting node & return double black node
func (t *tree) markNodes(n *node) (doubleBlack *node) {
	// 判断自己是否是 leaf node
	if n.leftChild == nil && n.rightChild == nil {
		// mark deleting node
		t.setDelNode(n)
		if n.color == COLOR_BLK { // double black, root 情况也被包含在里面。
			return n
		}
		// n.color = red，直接删除, 没有 double black node
		return nil
	}

	// 不是 leaf node 的情况：
	// 寻找 successor. (这里也可以选择 predecessor, 下面情况需要镜像处理)
	// successor 自己可以是 leftChild 或者是 rightChild.
	// successor 可能会有 1 个 rightChild，而且这个 rightChild 不会有其他 child。
	// 如果 successor 不存在，那么 n 一定会有一个 leftChild.
	successor := n.successor()

	// successor 不存在，那么 n 一定会有一个 leftChild.
	if successor == nil {
		t.setDelNode(n.leftChild)
		n.replaceNode(n.leftChild)
		if n.leftChild.color == COLOR_BLK { // double black
			return n.leftChild
		}
		return nil // 如果 left Child 是 red, 没有 double black node.
	}

	// successor 不是 nil 的情况
	// 如果 successor 没有 rightChild
	if successor.rightChild == nil {
		t.setDelNode(successor)
		n.replaceNode(successor)
		if successor.color == COLOR_BLK { // double black
			return successor
		}
		return nil // 如果 successor 是 red, 没有 double black node.
	}

	// 如果 successor 有 rightChild,
	// NOTE 这里需要 replace 2次, n -> successor, successor -> successor.rightChild
	t.setDelNode(successor.rightChild)
	n.replaceNode(successor)
	successor.replaceNode(successor.rightChild)
	if successor.rightChild.color == COLOR_BLK { // double black
		return successor.rightChild
	}
	return nil // 如果 successor.rightChild 是 red, 没有 double black node
}

func resolveDoubleBlack(dbNode *node) {
	loop := dbNode
	for loop != nil {
		// 循环检查直到 root
		loop = checkDoubleBlackAndRebalance(loop)
	}
}

/* NOTE
1. 如果一个 leaf node 如果是 black，那么他一定会有 sibling,
否则不满足红黑树的每个分支 black node 数量相同规则。
2. 不会出现某个 node 有一个 nil child 和 一个 black child 的情况，
因为不满足红黑树的每个分支 black node 数量相同规则。
3. 如果 sibling 是 red，则 parent 和 sibling 的两个 children
一定是 black(nil), 否则会造成 red red conflict。

db node situation：
1. db is root.
2. sibling == red
3. sibling == black
	3.1 both sibling's children are black(nil)
		3.1.1 parent == red
		3.1.2 parent == black
	3.2 far side of sibling's child is red.
		NOTE this situation includes: 2 red children OR 1 black(nil) 1 red child.
	3.3 near side of sibling's child is red, far side child is black(nil).
*/
func checkDoubleBlackAndRebalance(dbNode *node) (next *node) {
	// NOTE 这里没考虑出现 nil 的情况。

	// 1. double black is root, just remove double black。
	if dbNode == dbNode.tree.root {
		return nil // stop double black check.
	}

	parent := dbNode.parent
	sibling := dbNode.sibling()

	// 2. sibling == red
	// parent 和 sibling 的两个 children 一定是 black(nil),
	// 否则会造成 red-red conflict。
	if sibling.color == COLOR_RED {
		// recolor
		parent.reColor()
		sibling.reColor()
		// NOTE parent 进行 rotation，DB node 的位置决定 rotation 的方向。
		if dbNode == parent.leftChild { // left child
			parent.rightRightRotation()
		} else { // right child
			parent.leftLeftRotation()
		}
		// rotation 结束后，自己还是 double black，返回继续检查。
		return dbNode
	}

	// 3. sibling == black

	// 3.1 both sibling's children are black(nil)
	// NOTE if 内的判断顺序很重要，否则会造成 nil pointer 的 panic.
	if (sibling.leftChild == nil || sibling.leftChild.color == COLOR_BLK) &&
		(sibling.rightChild == nil || sibling.rightChild.color == COLOR_BLK) {
		// sibling 需要 recolor
		sibling.reColor()

		// 3.1.1 parent == red
		if parent.color == COLOR_RED {
			parent.reColor() // change parent to black
			return nil       // db resolved
		}

		// 3.1.2 parent == black
		return parent // parent becomes double black.
	}

	// sibling 的 children 中至少有一个是 red 的情况.
	// includes: 2 red child; 1 red & 1 black(nil) child.

	// db node 位置不同进行不同的操作。
	if dbNode == parent.leftChild {
		// 3.2 far side of sibling's child is red.
		// NOTE if 内的判断顺序很重要，否则会造成 nil pointer 的 panic.
		if sibling.rightChild != nil && sibling.rightChild.color == COLOR_RED {
			// NOTE 这里使用的 swapColor 而不是 recolor，因为 parent 的颜色不同采用相同处理方式。
			parent.swapColor(sibling)
			sibling.rightChild.reColor()
			parent.rightRightRotation()
			return nil // when red nephew is far away, db resolved.
		}

		// 3.3 near side of sibling's child is red, far side child is black(nil).
		sibling.leftChild.reColor()
		sibling.reColor()
		sibling.leftLeftRotation()
		return dbNode // when red nephew is near, re-check dbnode.
	}

	// db node is right child.
	// 3.2 far side of sibling's child is red.
	// NOTE if 内的判断顺序很重要，否则会造成 nil pointer 的 panic.
	if sibling.leftChild != nil && sibling.leftChild.color == COLOR_RED {
		// NOTE 这里使用的 swapColor 而不是 recolor，因为 parent 的颜色不同采用相同处理方式。
		parent.swapColor(sibling)
		sibling.leftChild.reColor()
		parent.leftLeftRotation()
		return nil // when red nephew is far away, db resolved.
	}

	// 3.3 near side of sibling's child is red, far side child is black(nil).
	sibling.rightChild.reColor()
	sibling.reColor()
	sibling.rightRightRotation()
	return dbNode // when red nephew is near, re-check dbnode.
}

// 缓存需要删除的 node.
func (t *tree) setDelNode(n *node) {
	t.cacheDelNode = n
}

// 实际删除被缓存的 node.
func (t *tree) deleteCachedNode() {
	if t.cacheDelNode != nil {
		if t.cacheDelNode.parent == nil { // root
			t.root = nil
		} else {
			if t.cacheDelNode == t.cacheDelNode.parent.leftChild { // left child
				t.cacheDelNode.parent.leftChild = nil
			} else { // right child
				t.cacheDelNode.parent.rightChild = nil
			}
		}

		t.cacheDelNode = nil // 重置为 nil
	}
}
