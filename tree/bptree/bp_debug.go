package bptree

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
)

func (t *tree) printTree() {
	var countNodes int
	var loop = []iNode{t.root}
	for loop != nil {
		var tmp []iNode
		for _, v := range loop {
			countNodes++
			if v.getParent() != nil {
				fmt.Printf("%v -> %v, %s", v.getParent().Keys(), v.Keys(), v.Type().String())
				if v.Link() != nil {
					fmt.Printf(", link -> %v\n", v.Link().Keys())
				} else {
					fmt.Printf("\n")
				}
			} else {
				fmt.Printf("root -> %v\n", v.Keys())
			}

			tmp = append(tmp, v.Children()...)
		}
		loop = tmp
	}
	fmt.Println("total Nodes:", countNodes)
}

type checkResult struct {
	index  []int64
	reason string
}

func (t *tree) checkAllNodes() []checkResult {
	if t.root == nil {
		return nil
	}

	var result []checkResult

	if t.root.getParent() != nil {
		result = append(result, checkResult{nil, "root parent is not nil"})
	}

	// 检查 size 是否正确
	if t.Size() != t.len() {
		result = append(result, checkResult{[]int64{int64(t.Size()), int64(t.len())}, "tree Size != tree Len"})
	}

	nodeList := []iNode{t.root}
	for nodeList != nil {
		var loop []iNode
		for _, n := range nodeList {
			// 检查 node cap
			ki, ci := n.cap()
			if ki != t.maxKey {
				result = append(result, checkResult{n.Keys(), "key cap error: " + strconv.Itoa(ki)})
			}

			if n.Type() == Internal && ci != t.maxKey+1 {
				result = append(result, checkResult{n.Keys(), "children cap error: " + strconv.Itoa(ci)})
			}

			// 检查 node 中的排序
			for i := 0; i < n.Size()-1; i++ {
				if n.Key(i) >= n.Key(i+1) {
					result = append(result, checkResult{n.Keys(), "sort error"})
				}
			}

			if n.Children() != nil {
				loop = append(loop, n.Children()...)

				if len(n.Keys())+1 != len(n.Children()) {
					result = append(result, checkResult{n.Keys(), "len children != keys+1"})
				}

				for i, v := range n.Keys() {
					// 检查 left child < key
					if n.Child(i).Key(n.Child(i).Size()-1) >= v {
						result = append(result, checkResult{n.Keys(), "left child bigger than parent"})
					}

					// 检查 right child >= key
					if n.Child(i+1).Key(0) < v {
						result = append(result, checkResult{n.Keys(), "right child smaller than parent"})
					}

					// 检查 leaf node 的 parent[index] = leaf node[0]
					if n.Child(i+1).Type() == Leaf && n.Child(i+1).Key(0) != v {
						result = append(result, checkResult{n.Keys(), "leaf's first element != parent"})
					}

					// 检查 children parent
					if n.Child(i).getParent() != n {
						result = append(result, checkResult{n.Keys(), "node's children's parent is not itself"})
					}
				}
			}

		}
		nodeList = loop
	}
	return result
}

func traceMemStatsMark(mark string) {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("%s: Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)\n", mark, ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}
