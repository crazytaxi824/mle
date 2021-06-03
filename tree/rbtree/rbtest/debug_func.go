package rbtest

import (
	"fmt"

	"github.com/crazytaxi824/mle/tree/rbtree"
)

func PrintALL(t rbtree.Tree) {
	if t.Root() == nil {
		fmt.Println("nil tree")
		return
	}

	fmt.Printf("tree root: %d size: %d\n", t.Root().Index(), t.Size())

	nodeList := []rbtree.Node{t.Root()}
	for nodeList != nil {
		var loop []rbtree.Node
		for _, n := range nodeList {
			if n.LeftChild() != nil {
				fmt.Printf("%d color: %s -> left %d color:%s\n",
					n.Index(), n.Color().String(), n.LeftChild().Index(), n.LeftChild().Color().String())
				loop = append(loop, n.LeftChild())
			}

			if n.RightChild() != nil {
				fmt.Printf("%d color: %s -> right %d color:%s\n",
					n.Index(), n.Color().String(), n.RightChild().Index(), n.RightChild().Color().String())
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

func CheckAllNodes(t rbtree.Tree) []CheckResult {
	if t.Root() == nil {
		return nil
	}

	var result []CheckResult

	// 检查 root parent
	if t.Root().Parent() != nil {
		result = append(result, CheckResult{t.Root().Index(), "root parent is not nil"})
	}

	nodeList := []rbtree.Node{t.Root()}
	for nodeList != nil {
		var loop []rbtree.Node
		for _, n := range nodeList {
			// check left right child
			if n.LeftChild() != nil {
				loop = append(loop, n.LeftChild())
				// 检查 left index 大小是否正确
				if n.LeftChild().Index() >= n.Index() {
					result = append(result, CheckResult{n.Index(), "leftChild >= n"})
				}
				// 检查 child parent
				if n.LeftChild().Parent() != n {
					result = append(result, CheckResult{n.Index(), "n's left child's parent is not n"})
				}
				// 检查 color
				if n.Color() == rbtree.COLOR_RED && n.LeftChild().Color() == rbtree.COLOR_RED {
					result = append(result, CheckResult{n.Index(), "n's left child - Red Red Conflict"})
				}
			}

			if n.RightChild() != nil {
				loop = append(loop, n.RightChild())
				// 检查 right index 大小是否正确
				if n.RightChild().Index() <= n.Index() {
					result = append(result, CheckResult{n.Index(), "rightChild <= n"})
				}
				// 检查 child parent
				if n.RightChild().Parent() != n {
					result = append(result, CheckResult{n.Index(), "n's right child's parent is not n"})
				}
				// 检查 color
				if n.Color() == rbtree.COLOR_RED && n.RightChild().Color() == rbtree.COLOR_RED {
					result = append(result, CheckResult{n.Index(), "n's right child - Red Red Conflict"})
				}
			}

			// 检查 Predecessor & Successor
			predecessor, successor := n.Predecessor(), n.Successor()
			if predecessor != nil && predecessor.Index() >= n.Index() {
				result = append(result, CheckResult{n.Index(), "predecessor >= n"})
			}
			if successor != nil && successor.Index() <= n.Index() {
				result = append(result, CheckResult{n.Index(), "successor <= n"})
			}

			// NOTE 检查每一个节点的两边分支的 BLACK node 数量是相同的
			var leftBlackCount, rightBlackCount int
			for leftLoop := n.LeftChild(); leftLoop != nil; leftLoop = leftLoop.LeftChild() {
				if leftLoop.Color() == rbtree.COLOR_BLK {
					leftBlackCount++
				}
			}
			for rightLoop := n.RightChild(); rightLoop != nil; rightLoop = rightLoop.RightChild() {
				if rightLoop.Color() == rbtree.COLOR_BLK {
					rightBlackCount++
				}
			}

			if leftBlackCount != rightBlackCount {
				result = append(result, CheckResult{n.Index(), "black node is unbalanced"})
			}
		}
		nodeList = loop
	}

	return result
}
