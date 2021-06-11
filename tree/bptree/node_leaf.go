package bptree

import "log"

type KV interface {
	Key() int64
	Value() interface{}
	SetValue(interface{})
}

type LeafContent struct {
	key   int64
	value interface{}
}

func (lc *LeafContent) Key() int64             { return lc.key }
func (lc *LeafContent) Value() interface{}     { return lc.value }
func (lc *LeafContent) SetValue(v interface{}) { lc.value = v }

type leafNode struct {
	baseNode
	data []KV  // key 数量
	link iNode // link to the right side leaf node
}

func (ln *leafNode) cap() (keys, children int) {
	return cap(ln.data), -1
}

func (ln *leafNode) Keys() []int64 {
	r := make([]int64, 0, ln.tree.maxKey)
	for _, v := range ln.data {
		r = append(r, v.Key())
	}
	return r
}

func (ln *leafNode) Key(index int) int64 {
	return ln.data[index].Key()
}

func (ln *leafNode) contents() []KV {
	return ln.data
}

func (ln *leafNode) content(index int) KV {
	return ln.data[index]
}

func (ln *leafNode) Value(index int) (interface{}, error) {
	return ln.data[index].Value(), nil
}

func (ln *leafNode) Size() int {
	return len(ln.data)
}

// 情况1, len(ln.data) < tree.maxKey
// 情况1.1，pos < len(ln.data) 插入
// 情况1.2，pos == len(ln.data), append
// 情况2, len(ln.data) == tree.maxKey, split & new leaf node
// 情况2.1, pos = tree.middleIndex, 返回自己的 key
// 情况2.2, pos < tree.middleIndex, 返回 data[middleIndex-1].key
// 情况2.3, pos > tree.middleIndex, 返回 data[middleIndex].key
// 返回需要向上传递的 key 和 iNode, 如果不需要向上传递 index，返回 nil iNode
func (ln *leafNode) addItem(key int64, value interface{}, rightChild iNode) (upKey int64, newRightNode iNode) {
	// find place
	pos := ln.findPlaceForNewItem(key)

	// insert index and value into data
	newData := &LeafContent{
		key:   key,
		value: value,
	}

	// 情况1, len(ln.data) < tree.maxKey
	if len(ln.data) < ln.tree.maxKey {
		// 情况1.1，pos < len(ln.data) 插入
		ln.insertData(newData, pos)
		return -1, nil
	}

	// 情况2, len(ln.data) == tree.maxKey, split & new leaf node
	if pos == ln.tree.middleKeyIndex { // new leaf need to go up
		// 情况2.1, pos = tree.middleIndex, 返回自己的 key
		return ln.newKeyIsAtMiddleIndex(newData, pos)
	} else if pos < ln.tree.middleKeyIndex {
		// 情况2.2, pos < tree.middleIndex, 返回 data[middleIndex-1].key
		return ln.newKeyIsLeftSideToMiddleIndex(newData, pos)
	}

	// 情况2.3, pos > tree.middleIndex, 返回 data[middleIndex].key
	return ln.newKeyIsRightSideToMiddleIndex(newData, pos)
}

// 自己正好在 middle key index 的位置.
func (ln *leafNode) newKeyIsAtMiddleIndex(newData *LeafContent, pos int) (upKey int64, newRightNode iNode) {
	// 情况2.1, pos = tree.middleIndex, 返回自己的 key
	newRightLeaf := ln.tree.newLeaf()

	// node 分开后重新分配 data
	newRightLeaf.data = append(newRightLeaf.data, newData)
	newRightLeaf.data = append(newRightLeaf.data, ln.data[pos:]...)

	// NOTE for GC release memory
	l := len(ln.data)
	for i := pos; i < l; i++ {
		ln.data[i] = nil
	}

	ln.data = ln.data[:pos]

	// 重新连接 link
	newRightLeaf.link = ln.link
	ln.link = newRightLeaf

	// 如果 leaf 没有 parent，说明 leaf 是 root 的情况
	if ln.parent == nil {
		ln.tree.genNewRootNode(newData.key, ln, newRightLeaf)
		return -1, nil
	}

	// 绑定 parent 关系
	newRightLeaf.setParent(ln.parent)

	// 如果 leaf 已经有 parent，返回 newRightLeaf
	return newData.key, newRightLeaf
}

// 自己在 middle key index 的左边.
func (ln *leafNode) newKeyIsLeftSideToMiddleIndex(newData *LeafContent, pos int) (upKey int64, newRightNode iNode) {
	// 情况2.2, pos < tree.middleIndex, 返回 data[middleIndex-1].key
	newRightLeaf := ln.tree.newLeaf()
	up := ln.data[ln.tree.middleKeyIndex-1].Key()

	// node 分开后重新分配 data
	newRightLeaf.data = append(newRightLeaf.data, ln.data[ln.tree.middleKeyIndex-1:]...)

	// NOTE for GC release memory
	l := len(ln.data)
	for i := ln.tree.middleKeyIndex - 1; i < l; i++ {
		ln.data[i] = nil
	}

	ln.data = ln.data[:ln.tree.middleKeyIndex-1]
	ln.insertData(newData, pos)

	// 重新连接 link
	newRightLeaf.link = ln.link
	ln.link = newRightLeaf

	// 如果 leaf 没有 parent，说明 leaf 是 root 的情况
	if ln.parent == nil {
		ln.tree.genNewRootNode(up, ln, newRightLeaf)
		return -1, nil
	}

	// 绑定 parent 关系
	newRightLeaf.setParent(ln.parent)

	// 如果 leaf 已经有 parent，返回 newRightLeaf
	return up, newRightLeaf
}

// 自己在 middle key index 的右边.
func (ln *leafNode) newKeyIsRightSideToMiddleIndex(newData *LeafContent, pos int) (upKey int64, newRightNode iNode) {
	// 情况2.3, pos > tree.middleIndex, 返回 data[middleIndex].key
	newRightLeaf := ln.tree.newLeaf()
	up := ln.data[ln.tree.middleKeyIndex].Key()

	// node 分开后重新分配 data
	newRightLeaf.data = append(newRightLeaf.data, ln.data[ln.tree.middleKeyIndex:pos]...)
	newRightLeaf.data = append(newRightLeaf.data, newData)
	newRightLeaf.data = append(newRightLeaf.data, ln.data[pos:]...)

	// NOTE for GC release memory
	l := len(ln.data)
	for i := ln.tree.middleKeyIndex; i < l; i++ {
		ln.data[i] = nil
	}

	ln.data = ln.data[:ln.tree.middleKeyIndex]

	// 重新连接 link
	newRightLeaf.link = ln.link
	ln.link = newRightLeaf

	// 如果 leaf 没有 parent，说明 leaf 是 root 的情况
	if ln.parent == nil {
		ln.tree.genNewRootNode(up, ln, newRightLeaf)
		return -1, nil
	}

	// 绑定 parent 关系
	newRightLeaf.setParent(ln.parent)

	// 如果 leaf 已经有 parent，返回 newRightLeaf
	return up, newRightLeaf
}

// 向 ln.data 中的指定位置(index) 插入 KV.
func (ln *leafNode) insertData(newData KV, index int) {
	// insert Data Into Leaf
	ln.data = append(ln.data, nil)
	l := len(ln.data)
	for i := l - 1; i > index; i-- {
		ln.data[i] = ln.data[i-1]
	}
	ln.data[index] = newData
}

// 返回 key 应该放在 node 的那个位置上 (index).
func (ln *leafNode) findPlaceForNewItem(key int64) int {
	for i, v := range ln.data {
		if key < v.Key() {
			return i
		}
	}
	return len(ln.data)
}

// 左侧 left leaf node -> right leaf node.
func (ln *leafNode) Link() iNode {
	if ln.link == nil {
		return nil
	}
	return ln.link
}

// 重新绑定 link 关系.
func (ln *leafNode) setLink(newNode iNode) {
	ln.link = newNode
}

// 在 node 中查找 key 的位置，如果 key 不存在返回 false
// 如果存在，返回 index & true.
func (ln *leafNode) findKey(key int64) (int, bool) {
	for i, v := range ln.data {
		if key == v.Key() {
			return i, true
		}
	}
	return 0, false
}

// 删除指定位置(index)的 KV
func (ln *leafNode) removeKVByIndex(index int) {
	tmp := ln.data
	ln.data = ln.data[:index]
	ln.data = append(ln.data, tmp[index+1:]...)
	// NOTE for GC release memory
	ln.data = append(ln.data, nil)
	ln.data = ln.data[:len(ln.data)-1]
}

// 获取 node 两边的 sibling. 如果不存在则返回 nil.
func siblings(parent iNode, parentIndex int) (left, right iNode) {
	// NOTE 这里没考虑 root 的情况，前面有判断
	if parentIndex == -1 { // first item
		return nil, parent.Child(parentIndex + 2)
	} else if parentIndex == parent.Size()-1 { // last item
		return parent.Child(parentIndex), nil
	}
	return parent.Child(parentIndex), parent.Child(parentIndex + 2)
}

// 根据 index 删除 item
func (ln *leafNode) removeItem(index, ci int) (next iNode, nextDelIndex, nextCi int) {
	// 如果自己是 root
	if ln == ln.tree.root {
		// 删除 key
		ln.removeKVByIndex(index)

		// 如果删除后 data 已经为 nil
		if len(ln.data) == 0 {
			ln.tree.root = nil
			// delete node
			ln.delete()
		}
		return nil, 0, -1 // 返回 nil，表示不需要再继续检查了
	}

	// 获取 parent index, pi 可能为 -1，如果 ln 是第一个 child 的时候
	// 这里不会返回 error, 只是为了实现 interface
	pi, _ := ln.parent.findKey(ln.data[0].Key())

	// 删除 key
	ln.removeKVByIndex(index)

	// check balance
	return ln.checkBalance(index, pi)
}

// 从右侧节点借 KV, 直接借右侧 right sibling 的第一个 KV, 需要更改 parent.key[pi+1]
// NOTE 如果借过来的 KV 变成自己的第一个 KV (自己是空节点的情况). 需要更改 parent.key[pi]
func (ln *leafNode) borrowFromRightSibling(pi int, rightSibling iNode) {
	// node append new KV
	lc := rightSibling.content(0)
	ln.data = append(ln.data, lc)

	// borrow 的 item 变成自己的 first item, 需要设置 parent.keys[pi]
	if pi != -1 && ln.Size() == 1 {
		ln.parent.keys[pi] = lc.Key()
	}

	// set parent.keys[pi+1] to right sibling's second key.
	key1 := rightSibling.Key(1)
	ln.parent.keys[pi+1] = key1

	// remove right sibling's first item
	rightSibling.removeKVByIndex(0)
}

// 从左侧节点借 KV, 直接借左侧 left sibling 的最后一个节点, 需要更改 parent.key[pi]
func (ln *leafNode) borrowFromLeftSibling(pi int, leftSibling iNode) {
	lastIndex := leftSibling.Size() - 1

	// node insert new KV to first index.
	lc := leftSibling.content(lastIndex)
	ln.insertData(lc, 0) // 新的 item 添加到最前面

	// parent key 需要设置
	ln.parent.keys[pi] = lc.Key()

	// 删除 leftSibling 的最后一个元素
	leftSibling.removeKVByIndex(lastIndex)
}

// 与右侧节点合并, 这里是将右侧 right sibling 节点合并到自己节点中. 以自己的节点为主.
func (ln *leafNode) mergeRightSibling(rightSibling iNode) {
	// relink
	ln.link = rightSibling.Link()

	// node append right sibling's data
	ln.data = append(ln.data, rightSibling.contents()...)
}

// 与左侧节点合并, 这里是将自己合并到 left sibling 中. 以 left sibling 为主.
func (ln *leafNode) mergeLeftSibling(leftSibling iNode) {
	// relink
	leftSibling.setLink(ln.link)

	// left node append right node
	leftSibling.appendContents(ln.data)
}

// 4 种情况：
// 1. borrow from right sibling
// 2. borrow from left sibling
// 3. merge right sibling
// 4. merge left sibling
func (ln *leafNode) checkBalance(index, pi int) (next iNode, nextDelIndex, nextCi int) {
	// 自己的 keys >= minkey
	if ln.Size() >= ln.tree.minKey {
		// 如果被删除的 item 是 leaf 的第一个 item，同时不在 parent 最左边
		// 需要改变 parent 的 key
		if index == 0 && pi != -1 {
			ln.parent.keys[pi] = ln.data[0].Key()
		}
		return nil, 0, -1 // 返回 nil，表示不需要再继续检查了
	}

	// 自己不是 root，同时 key < minkey
	// 先检查两边 sibling，看是否能够 borrow, 先右再左
	// NOTE 这里不能直接用 link，查找 sibling，因为 link 的 node 可能不是同一个 parent.
	leftSibling, rightSibling := siblings(ln.parent, pi)
	if leftSibling != nil && leftSibling.Size() > ln.tree.minKey {
		// 如果左边 left sibling 可以 borrow 的情况
		ln.borrowFromLeftSibling(pi, leftSibling)
		return nil, -1, -1 // 返回 nil，表示不需要再继续检查了
	} else if rightSibling != nil && rightSibling.Size() > ln.tree.minKey {
		// 如果右边 right sibling 可以 borrow 的情况
		ln.borrowFromRightSibling(pi, rightSibling)
		return nil, -1, -1 // 返回 nil，表示不需要再继续检查了
	} else if rightSibling != nil && rightSibling.Size() <= ln.tree.minKey {
		// 将 right sibling 合并到 ln(自己)的节点中
		ln.mergeRightSibling(rightSibling)

		// delete right sibling
		parent := ln.parent
		rightSibling.delete()

		// 删除对应的 parent 的 key & children
		return parent, pi + 1, pi + 2
	} else if leftSibling != nil && leftSibling.Size() <= ln.tree.minKey {
		// 将 ln(自己)合并到左侧节点 left sibling
		ln.mergeLeftSibling(leftSibling)

		// 删除自己
		parent := ln.parent
		ln.delete()

		// 删除对应的 parent 的 key & children
		return parent, pi, pi + 1
	}

	// DEBUG leftSibling == nil && rightSibling == nil 的情况不存在
	log.Println("remove item error")
	return nil, 0, -1
}

func (ln *leafNode) appendContents(lc []KV) {
	ln.data = append(ln.data, lc...)
}

// NOTE 切断所有指针，用于 GC。否则 node 可能不会被回收。
func (ln *leafNode) delete() {
	ln.parent = nil
	ln.link = nil
	ln.tree = nil
	ln.data = nil
}

func (ln *leafNode) Child(i int) iNode            { return nil }
func (ln *leafNode) Children() []iNode            { return nil }
func (ln *leafNode) removeChildByIndex(index int) {}
func (ln *leafNode) setKey(key int64, index int)  {}
