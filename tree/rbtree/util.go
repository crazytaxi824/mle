package rbtree

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// print all the nodes with relationships in the tree.
// this function is for test only.
func PrintAllNode(_node *node) {
	if _node == nil {
		return
	}

	if _node.parent != nil {
		side := "right"
		if _node.order < _node.parent.order {
			side = "left"
		}

		color := "BLACK"
		if _node.color {
			color = "RED"
		}

		fmt.Printf("parent: %d %s child -> %d, color: %s\n", _node.parent.order, side, _node.order, color)
	} else {
		fmt.Printf("The Root of This Tree: %d, color: %s\n", _node.order, "BLACK")
	}

	PrintAllNode(_node.leftChild)
	PrintAllNode(_node.rightChild)
}
