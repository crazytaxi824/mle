/*
b+tree 中 data 只储存在 leaf node, internal node 不储存data.
*/

package bptree

import "log"

type iNode interface {
	Tree() *tree                    // return tree which this node belongs to.
	Size() int                      // return number of keys of this node.
	Key(int) int64                  // return key of specific index.
	Child(int) iNode                // return child of specific index.
	Value(int) (interface{}, error) // leaf node return value, internal node return error.
	Keys() []int64                  // return keys of the node.
	Children() []iNode              // return children of the node.
	Type() NodeType                 // return type of the node.
	Link() iNode                    // leaf node link to the right side of the node.

	// internal function
	setKey(key int64, index int) // set new key to node index.
	getParent() *internalNode    // return parent of the node.
	setParent(*internalNode)     // set parent of the node.
	setLink(iNode)               // set link to a new node.
	addItem(int64, interface{}, iNode) (int64, iNode)
	removeItem(keyIndex, childIndex int) (nDelNode iNode, nKey, nChild int)
	removeKVByIndex(index int)
	removeChildByIndex(index int)
	contents() []KV // leaf node return all data, internal node return nil.
	content(int) KV
	appendContents([]KV)           // leaf node use to append content
	findKey(key int64) (int, bool) // true = found exact key
	delete()                       // delete node self

	// DEBUG use only
	// return key cap & child cap. use to check the cap of the slice.
	cap() (keyCap int, childCap int)
}

type baseNode struct {
	parent *internalNode
	typ    NodeType
	tree   *tree
}

func (bn *baseNode) Tree() *tree {
	if bn.tree == nil {
		return nil
	}
	return bn.tree
}

func (bn *baseNode) Type() NodeType {
	return bn.typ
}

func (bn *baseNode) getParent() *internalNode {
	return bn.parent
}

func (bn *baseNode) setParent(n *internalNode) {
	bn.parent = n
}

type internalNode struct {
	baseNode
	keys     []int64 // key 数量 >= minKey, key 数量 <= maxKey
	children []iNode // children 数量 = key 数量 + 1
}

// debug use only,
// 查看 key & children 的 cap 长度是否有变化, 用于排查他们是不是生成了新的底层 array.
func (in *internalNode) cap() (keys, children int) {
	return cap(in.keys), cap(in.children)
}

func (in *internalNode) Keys() []int64 {
	return in.keys
}

func (in *internalNode) Key(index int) int64 {
	return in.keys[index]
}

// 返回 key 的长度, children 的长度应该是 +1.
func (in *internalNode) Size() int {
	return len(in.keys)
}

func (in *internalNode) Child(i int) iNode {
	if in.children[i] == nil {
		return nil
	}
	return in.children[i]
}

func (in *internalNode) Children() []iNode {
	return in.children
}

// 情况1, len(ln.data) < tree.maxKey
// 情况2, len(ln.data) == tree.maxKey, split & new leaf node
// 情况2.1, pos = tree.middleIndex, 返回自己的 key
// 情况2.2, pos < tree.middleIndex, 返回 data[middleIndex-1].key
// 情况2.3, pos > tree.middleIndex, 返回 data[middleIndex].key
// NOTE return:
// 如果发生了 split node 操作, 至少会返回 right node. (left node = nil).
// 如果 root 节点发生了 split node 操作, 会返回 left & right node (都不为 nil).
// 如果没有发生 split node 操作, 返回的 left & right node 都为 nil.
func (in *internalNode) addItem(key int64, value interface{}, rightChild iNode) (upKey int64, newRightNode iNode) {
	// find pos
	pos := in.findPlaceForNewKey(key)

	// 情况1, len(ln.data) < tree.maxKey
	if len(in.keys) < in.tree.maxKey {
		in.insertKey(key, pos)
		in.insertChild(rightChild, pos+1)
		return -1, nil
	}

	// 情况2, len(ln.data) == tree.maxKey, split & new leaf node
	if pos == in.tree.middleKeyIndex { // new leaf need to go up
		// 情况2.1, pos = tree.middleIndex, 返回自己的 key
		return in.newKeyIsAtMiddleIndex(key, pos, rightChild)
	} else if pos < in.tree.middleKeyIndex {
		// 情况2.2, pos < tree.middleIndex, 返回 data[middleIndex-1].key
		return in.newKeyIsLeftSideToMiddleIndex(key, pos, rightChild)
	}

	// 情况2.3, pos > tree.middleIndex, 返回 data[middleIndex].key
	return in.newKeyIsRightSideToMiddleIndex(key, pos, rightChild)
}

// 自己正好在 middle key index 的位置.
func (in *internalNode) newKeyIsAtMiddleIndex(key int64, pos int, rightChild iNode) (upKey int64, newRightNode iNode) {
	// 情况2.1, pos = tree.middleIndex, 返回自己的 key
	newRightINode := in.tree.newInternal()

	// NOTE internode 不需要保留自己的 key 在右侧的 child 的第一个位置中。
	// node 分开后重新分配 key
	newRightINode.keys = append(newRightINode.keys, in.keys[pos:]...)
	in.keys = in.keys[:pos]

	// node 分开后重新绑定 children 关系.
	newRightINode.children = append(newRightINode.children, rightChild)
	newRightINode.children = append(newRightINode.children, in.children[pos+1:]...)

	// NOTE for GC release memory
	l := len(in.children)
	for i := pos + 1; i < l; i++ {
		in.children[i] = nil
	}

	in.children = in.children[:pos+1]

	// new right node 的 children 重新绑定 panrent
	for _, ch := range newRightINode.children {
		ch.setParent(newRightINode)
	}

	// 如果 node 没有 parent，说明 node 是 root 的情况
	if in.parent == nil {
		in.tree.genNewRootNode(key, in, newRightINode)
		return -1, nil
	}

	// 绑定 parent 关系
	newRightINode.setParent(in.parent)

	// 如果 node 已经有 parent，返回 newRightINode
	return key, newRightINode
}

// 自己在 middle key index 的左边.
func (in *internalNode) newKeyIsLeftSideToMiddleIndex(key int64, pos int, rightChild iNode) (upKey int64, newRightNode iNode) {
	// 情况2.2, pos < tree.middleIndex, 返回 data[middleIndex-1].key
	newRightINode := in.tree.newInternal()
	up := in.keys[in.tree.middleKeyIndex-1]

	// NOTE internode 不需要保留自己的 key 在右侧的 child 的第一个位置中。
	// node 分开后重新分配 key
	newRightINode.keys = append(newRightINode.keys, in.keys[in.tree.middleKeyIndex:]...)
	in.keys = in.keys[:in.tree.middleKeyIndex-1]
	in.insertKey(key, pos)

	// node 分开后重新绑定 children 关系.
	newRightINode.children = append(newRightINode.children, in.children[in.tree.middleKeyIndex:]...)

	// NOTE for GC release memory
	l := len(in.children)
	for i := in.tree.middleKeyIndex; i < l; i++ {
		in.children[i] = nil
	}

	in.children = in.children[:in.tree.middleKeyIndex]
	in.insertChild(rightChild, pos+1)

	// new right node 的 children 需要重新绑定 panrent
	for _, ch := range newRightINode.children {
		ch.setParent(newRightINode)
	}

	// 如果 node 没有 parent，说明 node 是 root 的情况
	if in.parent == nil {
		in.tree.genNewRootNode(up, in, newRightINode)
		return -1, nil
	}

	// 绑定 parent 关系
	newRightINode.setParent(in.parent)

	// 如果 node 已经有 parent，返回 newRightINode
	return up, newRightINode
}

// 自己在 middle key index 的右边.
func (in *internalNode) newKeyIsRightSideToMiddleIndex(key int64, pos int, rightChild iNode) (upKey int64, newRightNode iNode) {
	// 情况2.3, pos > tree.middleIndex, 返回 data[middleIndex].key
	newRightINode := in.tree.newInternal()
	up := in.keys[in.tree.middleKeyIndex]

	// NOTE internode 不需要保留自己的 key 在右侧的 child 的第一个位置中。
	// node 分开后重新分配 key
	newRightINode.keys = append(newRightINode.keys, in.keys[in.tree.middleKeyIndex+1:pos]...)
	newRightINode.keys = append(newRightINode.keys, key)
	newRightINode.keys = append(newRightINode.keys, in.keys[pos:]...)
	in.keys = in.keys[:in.tree.middleKeyIndex]

	// node 分开后重新绑定 children 关系.
	newRightINode.children = append(newRightINode.children, in.children[in.tree.middleKeyIndex+1:pos+1]...)
	newRightINode.children = append(newRightINode.children, rightChild)
	newRightINode.children = append(newRightINode.children, in.children[pos+1:]...)

	// NOTE for GC release memory
	l := len(in.children)
	for i := in.tree.middleKeyIndex + 1; i < l; i++ {
		in.children[i] = nil
	}

	in.children = in.children[:in.tree.middleKeyIndex+1]

	// new right node 的 children 需要重新绑定 panrent
	for _, ch := range newRightINode.children {
		ch.setParent(newRightINode)
	}

	// 如果 node 没有 parent，说明 node 是 root 的情况
	if in.parent == nil { // in 是 root 的情况
		in.tree.genNewRootNode(up, in, newRightINode)
		return -1, nil
	}

	// 绑定 parent 关系
	newRightINode.setParent(in.parent)

	// 如果 node 已经有 parent，返回 newRightINode
	return up, newRightINode
}

// 向 in.keys 中的指定位置(index) 插入 key.
func (in *internalNode) insertKey(newKey int64, index int) {
	in.keys = append(in.keys, 0)
	l := len(in.keys)
	for i := l - 1; i > index; i-- {
		in.keys[i] = in.keys[i-1]
	}
	in.keys[index] = newKey
}

// 向 in.children 中的指定位置(index) 插入 child.
func (in *internalNode) insertChild(newChild iNode, index int) {
	in.children = append(in.children, nil)
	l := len(in.children)
	for i := l - 1; i > index; i-- {
		in.children[i] = in.children[i-1]
	}
	in.children[index] = newChild
}

// 返回 key 应该放在 node 的那个位置上 (index).
func (in *internalNode) findPlaceForNewKey(key int64) int {
	for i, v := range in.keys {
		if key < v {
			return i
		}
	}
	return len(in.keys)
}

// 删除指定位置(index)的 key
func (in *internalNode) removeKVByIndex(index int) {
	keyTmp := in.keys
	in.keys = in.keys[:index]
	in.keys = append(in.keys, keyTmp[index+1:]...)
}

// 删除指定位置(index)的 child
// NOTE slice 中内存释放的问题。GC 后底层 array 并没有释放内存。
func (in *internalNode) removeChildByIndex(index int) {
	childTmp := in.children
	in.children = in.children[:index]
	in.children = append(in.children, childTmp[index+1:]...)
	// NOTE for GC release memory
	in.children = append(in.children, nil)
	in.children = in.children[:len(in.children)-1]
}

// 删除 Key 和 child.
func (in *internalNode) removeItem(keyIndex, childIndex int) (next iNode, nextDelKeyIndex, nextDelChildIndex int) {
	// root 情况
	if in == in.tree.root {
		in.removeKVByIndex(keyIndex)
		in.removeChildByIndex(childIndex)

		// 如果删除了 root 最后一个元素
		if in.Size() == 0 {
			in.tree.root = in.Child(0) // root = 唯一的 child
			in.Child(0).setParent(nil) // 清除新 root 的 parent 信息
			// delete node
			in.delete()
		}
		return nil, -1, -1
	}

	// 删除元素前获取 parent index，否则 in.keys 可能变为 nil
	parentIndex, _ := in.parent.findKey(in.keys[0])

	in.removeKVByIndex(keyIndex)
	in.removeChildByIndex(childIndex)

	return in.checkBalance(parentIndex)
}

// borrow right sibling's first key form parent.
// 将 parent.key[pi+1] 借给自己，将 right sibling 的第一个 key 给到 parent[pi+1]
// right sibling 的第一个 child 变成自己的最后一个 child.
// 如果自己的 child 是 leaf node, 则需要修改借到的 key 的值为 child 的第一个 key.
func (in *internalNode) borrowFromRightSibling(pi int, rightSibling iNode) {
	rightSiblingFirstChild := rightSibling.Child(0)

	if rightSiblingFirstChild.Type() == Leaf {
		// 如果 child 是 leaf node，borrow 的 key
		// 设置为 right Sibling First Child first Key
		in.keys = append(in.keys, rightSiblingFirstChild.Key(0))
	} else {
		// 如果 child 不是 leaf node，borrow 的 key = parent.keys[pi+1].
		in.keys = append(in.keys, in.parent.keys[pi+1])
	}

	// parent.keys[pi+1] change to right sibling's first key.
	in.parent.keys[pi+1] = rightSibling.Key(0)

	// node append right sibling's first child.
	in.children = append(in.children, rightSiblingFirstChild)

	// reset right sibling's first child parent to node.
	rightSiblingFirstChild.setParent(in)

	// delete right sibling's first key, and first child.
	rightSibling.removeKVByIndex(0)
	rightSibling.removeChildByIndex(0)
}

// borrow left sibling's last key form parent.
// 将 parent.key[pi] 借给自己，将 left sibling 的最后一个 Key 给到 parent[pi].
// 将 left sibling 的最后一个 child 变成自己的第一个 child.
// 如果自己的 child 是 leaf node, 则需要修改借到的 key 的值为 child 的第一个 key.
func (in *internalNode) borrowFromLeftSibling(pi int, leftSibling iNode) {
	lastIndex := leftSibling.Size() - 1

	leftSiblingLastChild := leftSibling.Child(lastIndex + 1)
	inFirstChild := in.children[0]

	if inFirstChild.Type() == Leaf {
		in.insertKey(inFirstChild.Key(0), 0)
	} else {
		in.insertKey(in.parent.keys[pi], 0)
	}

	// parent.keys[pi] change to left sibling's last key.
	in.parent.keys[pi] = leftSibling.Key(lastIndex)

	// insert left sibling's last child to node's first index.
	in.insertChild(leftSiblingLastChild, 0)

	// reset left sibling's last child's parent need to node.
	leftSiblingLastChild.setParent(in)

	// delete left sibling's last key and last child.
	leftSibling.removeKVByIndex(lastIndex)
	leftSibling.removeChildByIndex(lastIndex + 1)
}

// 将 right node merge 到 in node 中，删除 right node.
func (in *internalNode) mergeRightSibling(pi int, rightSibling iNode) {
	// node append parent.keys[pi+1], append all right sibling keys.
	in.keys = append(in.keys, in.parent.keys[pi+1])
	in.keys = append(in.keys, rightSibling.Keys()...)

	// node append all right sibling's children.
	// reset right sibling's children's parent need to node.
	for _, child := range rightSibling.Children() {
		child.setParent(in)
		in.children = append(in.children, child)
	}

	// delete right sibling
	rightSibling.delete()
}

// 将 left sibling merge 到 in node 中，删除 left node.
func (in *internalNode) mergeLeftSibling(pi int, leftSibling iNode) {
	// left sibling append parent.keys[pi], append all node's children.
	lk := leftSibling.Keys()
	lk = append(lk, in.parent.keys[pi])
	lk = append(lk, in.keys...)

	// left sibling append all node's children
	lc := leftSibling.Children()
	lc = append(lc, in.children...)

	// 将所有数据赋值到 in node 中.
	in.keys = lk
	in.children = lc

	// reset left sibling's children's parent need to node.
	for _, v := range leftSibling.Children() {
		v.setParent(in)
	}

	// delete left sibling
	leftSibling.delete()
}

// 4 种情况：
// 1. borrow from right sibling
// 2. borrow from left sibling
// 3. merge right sibling
// 4. merge left sibling
// NOTE internal Node merge 中 parent 需要一起参与 merge.
// NOTE internal Node borrow 需要调整 parent 的 key.
func (in *internalNode) checkBalance(pi int) (next iNode, nextDelIndex, nextDelChildIndex int) {
	// 自己的 keys >= minkey
	if in.Size() >= in.tree.minKey {
		return nil, 0, -1 // 返回 nil，表示不需要再继续检查了
	}

	// 自己不是 root，同时 key < minkey
	// 先检查两边 sibling，看是否能够 borrow, 先右再左
	leftSibling, rightSibling := siblings(in.parent, pi)
	if rightSibling != nil && rightSibling.Size() > in.tree.minKey { // 如果右边可以 borrow
		in.borrowFromRightSibling(pi, rightSibling)

		// no need to check again
		return nil, -1, -1
	} else if leftSibling != nil && leftSibling.Size() > in.tree.minKey {
		in.borrowFromLeftSibling(pi, leftSibling)

		// no need to check again
		return nil, -1, -1
	} else if rightSibling != nil && rightSibling.Size() <= in.tree.minKey {
		in.mergeRightSibling(pi, rightSibling)

		// NOTE 删除 parent 的元素时 删除 children
		return in.parent, pi + 1, pi + 2
	} else if leftSibling != nil && leftSibling.Size() <= in.tree.minKey {
		in.mergeLeftSibling(pi, leftSibling)

		// NOTE 删除 parent 的元素时 删除 children
		return in.parent, pi, pi
	}

	// DEBUG leftSibling == nil && rightSibling == nil 的情况不存在
	log.Println("remove item error")
	return nil, 0, -1
}

// NOTE internal node 中 key 不会在上层重复。
// 返回的 index 可能为 -1. 如果为 -1 说明自己是最左侧的 child.
// 如果返回的是 len(in.keys)-1 说明自己是最右侧的 child.
// true = found exact key
func (in *internalNode) findKey(key int64) (int, bool) {
	for i, v := range in.keys {
		if key == v {
			return i, true // return true = found exact key
		} else if key < v {
			return i - 1, false
		}
	}
	return len(in.keys) - 1, false
}

// 修改 node 中指定位置(index) key 的值。
func (in *internalNode) setKey(key int64, index int) {
	in.keys[index] = key
}

// NOTE 切断所有指针，用于 GC。否则 node 可能不会被回收。
func (in *internalNode) delete() {
	in.parent = nil
	in.children = nil
	in.tree = nil
}

func (in *internalNode) Link() iNode                          { return nil }
func (in *internalNode) setLink(newNode iNode)                {}
func (in *internalNode) contents() []KV                       { return nil }
func (in *internalNode) content(int) KV                       { return nil }
func (in *internalNode) appendContents(lc []KV)               {}
func (in *internalNode) Value(index int) (interface{}, error) { return nil, ErrInternalValue }
