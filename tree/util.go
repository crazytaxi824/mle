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

// print all element
func PrintAllNode(_node *node) {
	if _node == nil {
		return
	}

	if _node.parent != nil {
		side := "right"
		if _node.isLeftChild {
			side = "left"
		}
		fmt.Printf("parent: %d %s child -> %d, depth: %d\n", _node.parent.order, side, _node.order, _node.depth)
	} else {
		fmt.Printf("tree Root: %d, depth: %d\n", _node.order, _node.depth)
	}

	PrintAllNode(_node.leftChild)
	PrintAllNode(_node.rightChild)
}
