package example

import (
	"fmt"
	"local/src/avltree"
)

func PrintTree(t avltree.Tree) {
	if t.Root() == nil {
		fmt.Println("nil tree")
		return
	}

	fmt.Printf("tree root: %d size: %d\n", t.Root().Index(), t.Size())

	nodeList := []avltree.Node{t.Root()}
	for nodeList != nil {
		var loop []avltree.Node
		for _, n := range nodeList {
			if n.LeftChild() != nil {
				fmt.Printf("%d depth: %d -> left %d depth:%d\n",
					n.Index(), n.Depth(), n.LeftChild().Index(), n.LeftChild().Depth())
				loop = append(loop, n.LeftChild())
			}

			if n.RightChild() != nil {
				fmt.Printf("%d depth: %d -> right %d depth:%d\n",
					n.Index(), n.Depth(), n.RightChild().Index(), n.RightChild().Depth())
				loop = append(loop, n.RightChild())
			}
		}
		nodeList = loop
	}
}

type CheckResult struct {
	index  int64
	reason string
}

func CheckAllNodes(t avltree.Tree) []CheckResult { // nolint
	if t.Root() == nil {
		fmt.Println("nil tree")
		return nil
	}

	var result []CheckResult
	// 检查 root parent
	if t.Root().Parent() != nil {
		result = append(result, CheckResult{t.Root().Index(), "root parent is not nil"})
	}

	nodeList := []avltree.Node{t.Root()}
	for nodeList != nil {
		var loop []avltree.Node
		for _, n := range nodeList {
			var leftDep, rightDep int
			if n.LeftChild() != nil {
				leftDep = n.LeftChild().Depth()
				loop = append(loop, n.LeftChild())
				// 检查 left index 大小是否正确
				if n.LeftChild().Index() >= n.Index() {
					result = append(result, CheckResult{n.Index(), "leftChild >= n"})
				}
				// 检查 child 的 parent 是不是自己
				if n != n.LeftChild().Parent() {
					result = append(result, CheckResult{n.Index(), "n's left child's parent is not n"})
				}
			}

			if n.RightChild() != nil {
				rightDep = n.RightChild().Depth()
				loop = append(loop, n.RightChild())
				// 检查 right index 大小是否正确
				if n.RightChild().Index() <= n.Index() {
					result = append(result, CheckResult{n.Index(), "rightChild <= n"})
				}
				// 检查 child 的 parent 是不是自己
				if n != n.RightChild().Parent() {
					result = append(result, CheckResult{n.Index(), "n's right child's parent is not n"})
				}
			}

			// 检查是否 balance
			if leftDep-rightDep > 1 || leftDep-rightDep < -1 {
				result = append(result, CheckResult{n.Index(), "is unbalanced"})
			}

			// 检查 depth 数据是否正确
			if n.Depth() != max(leftDep, rightDep)+1 {
				result = append(result, CheckResult{n.Index(), "wrong depth"})
			}

			// 检查 Predecessor & Successor
			predecessor, successor := n.Predecessor(), n.Successor()
			if predecessor != nil && predecessor.Index() >= n.Index() {
				result = append(result, CheckResult{n.Index(), "predecessor >= n"})
			}
			if successor != nil && successor.Index() <= n.Index() {
				result = append(result, CheckResult{n.Index(), "successor <= n"})
			}
		}
		nodeList = loop
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
