package rbtree

import (
	"log"
	"runtime"
)

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
				log.Printf("%d color: %v -> left %d color:%v\n",
					n.index, n.color, n.leftChild.index, n.leftChild.color)
				loop = append(loop, n.leftChild)
			}

			if n.rightChild != nil {
				log.Printf("%d color: %v -> right %d color:%v\n",
					n.index, n.color, n.rightChild.index, n.rightChild.color)
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

func (t *tree) checkAllNodes() []CheckResult {
	if t.root == nil {
		return nil
	}

	var result []CheckResult

	// 检查 root parent
	if t.root.parent != nil {
		result = append(result, CheckResult{t.root.index, "root parent is not nil"})
	}

	// 检查 root delnode 是否为 nil
	if t.cacheDelNode != nil {
		result = append(result, CheckResult{t.root.index, "tree's delnode is not nil"})
	}

	// 检查 Size() 是否正确
	var nodesCount int

	nodeList := []*node{t.root}
	for nodeList != nil {
		// 检查 Size() 是否正确
		nodesCount += len(nodeList)

		var loop []*node
		for _, n := range nodeList {
			// check left right child
			if n.leftChild != nil {
				loop = append(loop, n.leftChild)
				// 检查 left index 大小是否正确
				if n.leftChild.index >= n.index {
					result = append(result, CheckResult{n.index, "leftChild >= n"})
				}
				// 检查 child parent
				if n.leftChild.parent != n {
					result = append(result, CheckResult{n.index, "n's left child's parent is not n"})
				}
				// 检查 red red conflict
				if n.color == COLOR_RED && n.leftChild.color == COLOR_RED {
					result = append(result, CheckResult{n.index, "n's left child - Red Red Conflict"})
				}
			}

			if n.rightChild != nil {
				loop = append(loop, n.rightChild)
				// 检查 right index 大小是否正确
				if n.rightChild.index <= n.index {
					result = append(result, CheckResult{n.index, "rightChild <= n"})
				}
				// 检查 child parent
				if n.rightChild.parent != n {
					result = append(result, CheckResult{n.index, "n's right child's parent is not n"})
				}
				// 检查 red red conflict
				if n.color == COLOR_RED && n.rightChild.color == COLOR_RED {
					result = append(result, CheckResult{n.index, "n's right child - Red Red Conflict"})
				}
			}

			// 检查 Predecessor & Successor
			predecessor, successor := n.predecessor(), n.successor()
			if predecessor != nil && predecessor.index >= n.index {
				result = append(result, CheckResult{n.index, "predecessor >= n"})
			}
			if successor != nil && successor.index <= n.index {
				result = append(result, CheckResult{n.index, "successor <= n"})
			}

			// NOTE 检查每一个节点的两边分支的 BLACK node 数量是相同的
			var leftBlackCount, rightBlackCount int
			for leftLoop := n.leftChild; leftLoop != nil; leftLoop = leftLoop.leftChild {
				if leftLoop.color == COLOR_BLK {
					leftBlackCount++
				}
			}
			for rightLoop := n.rightChild; rightLoop != nil; rightLoop = rightLoop.rightChild {
				if rightLoop.color == COLOR_BLK {
					rightBlackCount++
				}
			}

			if leftBlackCount != rightBlackCount {
				result = append(result, CheckResult{n.index, "black node is unbalanced"})
			}
		}
		nodeList = loop
	}

	// 检查 Size() 是否正确
	if nodesCount != t.Size() {
		result = append(result, CheckResult{-1, "nodes count != tree Size()"})
	}

	return result
}

func traceMemStatsMark(mark string) {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("%s: Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)\n", mark, ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}
