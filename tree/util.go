package tree

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 检测用，打印树中的所有节点，表明关系和深度
func PrintAllNode(_node *node) {
	if _node == nil {
		return
	}

	if _node.parent != nil {
		side := "right"
		if _node.order < _node.parent.order {
			side = "left"
		}
		fmt.Printf("parent: %d %s child -> %d, depth: %d\n", _node.parent.order, side, _node.order, _node.depth)
	} else {
		fmt.Printf("The Root of This Tree: %d, depth: %d\n", _node.order, _node.depth)
	}

	PrintAllNode(_node.leftChild)
	PrintAllNode(_node.rightChild)
}
