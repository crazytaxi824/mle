package avltree

import "log"

func (t *tree) printTree() {
	if t.root == nil {
		log.Println("nil tree")
		return
	}

	log.Printf("tree root: %d size: %d\n", t.root.index, t.length)

	nodeList := []*node{t.root}
	for nodeList != nil {
		var loop []*node
		for _, n := range nodeList {
			if n.leftChild != nil {
				log.Printf("%d depth: %d -> left %d depth:%d\n",
					n.index, n.depth, n.leftChild.index, n.leftChild.depth)
				loop = append(loop, n.leftChild)
			}

			if n.rightChild != nil {
				log.Printf("%d depth: %d -> right %d depth:%d\n",
					n.index, n.depth, n.rightChild.index, n.rightChild.depth)
				loop = append(loop, n.rightChild)
			}
		}
		nodeList = loop
	}
}

type CheckResult struct {
	index  int64
	reason string
}

func (t *tree) checkAllNodes() []CheckResult { // nolint
	if t.root == nil {
		return nil
	}

	var result []CheckResult

	// 检查 root parent
	if t.root.parent != nil {
		result = append(result, CheckResult{t.root.index, "root parent is not nil"})
	}

	nodeList := []*node{t.root}
	for nodeList != nil {
		var loop []*node
		for _, n := range nodeList {
			// check left right child
			var leftDep, rightDep int
			if n.leftChild != nil {
				leftDep = n.leftChild.depth
				loop = append(loop, n.leftChild)
				// 检查 left index 大小是否正确
				if n.leftChild.index >= n.index {
					result = append(result, CheckResult{n.index, "leftChild >= n"})
				}
				// 检查 child 的 parent 是不是自己
				if n != n.leftChild.parent {
					result = append(result, CheckResult{n.index, "n's left child's parent is not n"})
				}
			}

			if n.rightChild != nil {
				rightDep = n.rightChild.depth
				loop = append(loop, n.rightChild)
				// 检查 right index 大小是否正确
				if n.rightChild.index <= n.index {
					result = append(result, CheckResult{n.index, "rightChild <= n"})
				}
				// 检查 child 的 parent 是不是自己
				if n != n.rightChild.parent {
					result = append(result, CheckResult{n.index, "n's right child's parent is not n"})
				}
			}

			// 检查是否 balance
			if leftDep-rightDep > 1 || leftDep-rightDep < -1 {
				result = append(result, CheckResult{n.index, "is unbalanced"})
			}

			// 检查 depth 数据是否正确
			if n.depth != max(leftDep, rightDep)+1 {
				result = append(result, CheckResult{n.index, "wrong depth"})
			}

			// 检查 Predecessor & Successor
			predecessor, successor := n.predecessor(), n.successor()
			if predecessor != nil && predecessor.index >= n.index {
				result = append(result, CheckResult{n.index, "predecessor >= n"})
			}
			if successor != nil && successor.index <= n.index {
				result = append(result, CheckResult{n.index, "successor <= n"})
			}
		}
		nodeList = loop
	}

	return result
}
