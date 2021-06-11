/*
if Order of a B+ tree is 7,
Internal Node(except root) may have 4~7 children;
Root Node may have 2~7 children;
Leaf Node have no child.

if Order = x
Max key number = x-1
Min key number = (x-1)/2
Max children numer = x
Min children numer = (x+1)/2
middle(index) = x/2(right bias), index从0开始算.

if Order=4; maxKey=3; minKey=(4-1)/2=1; maxChildren=4; minChildren=2; middle=4/2=2
if Order=5; maxKey=4; minKey=(5-1)/2=2; maxChildren=5; minChildren=3; middle=5/2=2
if Order=6; maxKey=5; minKey=(6-1)/2=2; maxChildren=6; minChildren=3; middle=6/2=3
if Order=7; maxKey=6; minKey=(7-1)/2=3; maxChildren=7; minChildren=4; middle=7/2=3
*/

package bptree

import (
	"errors"
)

type Tree interface {
	// return order of the b+ tree.
	Order() int

	// add kv
	Insert(key int64, value interface{}) error

	// delete kv by key
	Delete(key int64) error

	// get KV by key
	Search(key int64) (KV, error)

	// get n number of KVs whose key is greater than start
	// n = 0 means get all the KVs whose key is greater than start.
	SearchGreaterThan(key int64, equal bool, limit, offset int) []KV

	// get n number of KVs whose key is smaller than end
	// n = 0 means get all the KVs whose key is smaller than end.
	SearchLessThan(key int64, equal bool, limit, offset int) []KV

	// get a range of values whose key is greater than start
	// and smaller than end.
	// [start,end), [start,end], (start,end], (start,end)
	SearchFromTo(start int64, startEqual bool, end int64, endEqual bool, limit, offset int) []KV

	// the number of elements in the tree.
	Size() int

	// height of the tree
	Height() int

	// smallest element
	Smallest() KV

	// largest element
	Largest() KV

	// return all elements of the tree, from smallest to largest.
	Sort() []KV

	// Clear the whole tree. delete every node in the tree.
	// NOTE
	// if just set tree.root = nil, GC may not release memory.
	// if you want to release memory, you have to Clear() after
	// using the tree.
	Clear()
}

type tree struct {
	order, minKey, maxKey int // 每个 node 的大小
	middleKeyIndex        int // node split 的时候，哪个 index 的数字向上移动
	root                  iNode
	length                int // 节点总数
}

func NewTree(order int) (Tree, error) {
	return newTree(order)
}

// for internal use only
func newTree(order int) (*tree, error) {
	if order < 3 {
		return nil, errors.New("order too small, must >= 3")
	}
	t := &tree{
		order:          order,
		minKey:         (order - 1) / 2,
		maxKey:         order - 1,
		middleKeyIndex: order / 2,
	}
	return t, nil
}

func (t *tree) newLeaf() *leafNode {
	ln := &leafNode{
		baseNode: baseNode{
			typ:  Leaf,
			tree: t,
		},
		data: make([]KV, 0, t.maxKey),
	}
	// runtime.SetFinalizer(ln, func(p *leafNode) {
	// 	log.Println(p.Type(), "GC ~~~~~~~~~~~~~~~~")
	// })
	return ln
}

func (t *tree) newInternal() *internalNode {
	in := &internalNode{
		baseNode: baseNode{
			typ:  Internal,
			tree: t,
		},
		keys:     make([]int64, 0, t.maxKey),
		children: make([]iNode, 0, t.maxKey+1),
	}
	// runtime.SetFinalizer(in, func(p *internalNode) {
	// 	log.Println(p.Type(), "GC ~~~~~~~~~~~~~~~~")
	// })
	return in
}

func (t *tree) Order() int {
	return t.order
}

func (t *tree) Size() int {
	return t.length
}

func (t *tree) Height() int {
	if t.root == nil {
		return 0
	}

	var h int
	for loop := t.root; loop != nil; loop = loop.Child(0) {
		h++
	}
	return h
}

func (t *tree) Smallest() KV {
	fl := t.firstLeaf()
	if fl == nil {
		return nil
	}

	return fl.content(0)
}

func (t *tree) Largest() KV {
	if t.root == nil {
		return nil
	}

	loop := t.root
	for loop.Children() != nil {
		loop = loop.Child(loop.Size())
	}

	return loop.content(loop.Size() - 1)
}

func (t *tree) Sort() []KV {
	if t.root == nil {
		return nil
	}

	// find most left leaf node
	loop := t.root
	for loop.Children() != nil {
		loop = loop.Child(0)
	}

	// append all data
	var result []KV
	for loop != nil {
		for _, v := range loop.contents() {
			tmp := v
			result = append(result, tmp)
		}
		loop = loop.Link()
	}

	return result
}

func (t *tree) Search(key int64) (KV, error) {
	leaf := t.findLeafNode(key)
	// leaf is nil only when root is nil.
	if leaf == nil {
		return nil, ErrIndexNotExist
	}

	i, ok := leaf.findKey(key)
	if !ok {
		return nil, ErrIndexNotExist
	}

	return leaf.content(i), nil
}

// 从 key 开始读取直到最后.
// equal: true 意思是 >= key; false 意思是 > key.
// limit 限制返回多少条记录. limit = 0, 不限值数据量.
// offset 设置跳过多少条记录开始读取数据.
func (t *tree) SearchGreaterThan(key int64, equal bool, limit, offset int) []KV {
	// start at leaf node
	startNode := t.findLeafNode(key)
	if startNode == nil {
		return nil
	}

	// start position
	startPos, ok := findPos(startNode, key)
	// 如果返回的是 exact key，需要向后移动一位
	if ok && !equal {
		startPos++
	}

	// 如果超出 index 范围需要移到下一个 node.
	if startPos == len(startNode.contents()) {
		startNode = startNode.Link()
		startPos = 0
	}

	// offset 设置
	if offset > 0 {
		startNode, startPos = findOffNode(startNode, startPos, offset)
	}

	// 判断是否有 limit
	var mark bool
	if limit > 0 {
		mark = true // result limited
	}

	var result []KV
	var count = limit // count 还剩多少数据需要 append

	for loop := startNode; loop != nil; loop = loop.Link() {
		tmplc := loop.contents()
		if !mark {
			result = append(result, tmplc[startPos:]...)
			startPos = 0
		} else {
			l := len(tmplc)
			if count >= l-startPos {
				result = append(result, tmplc[startPos:]...)
				count -= (l - startPos)
				startPos = 0
			} else {
				result = append(result, tmplc[:count]...)
				break
			}
		}
	}

	return result
}

// 从第一个 KV 开始读取直到 key 结束.
// equal: true 意思是 <= key; false 意思是 < key.
// limit 限制返回多少条记录. limit = 0, 不限值数据量.
// offset 设置跳过多少条记录开始读取数据.
func (t *tree) SearchLessThan(key int64, equal bool, limit, offset int) []KV {
	// start at leaf node
	startNode := t.firstLeaf()
	if startNode == nil {
		return nil
	}

	// start position
	startPos := 0

	// offset 设置
	if offset > 0 {
		startNode, startPos = findOffNode(startNode, startPos, offset)
	}

	return getResultFromStartToEnd(startNode, startPos, key, equal, limit)
}

// 从 start 开始取值，一直到 end 结束。
// startEqual: true 意思是 >= start; false 意思是 > start.
// endEqual: true 意思是 <= end; false 意思是 < end.
// limit 限制返回多少条记录. limit = 0, 不限值数据量.
// offset 设置跳过多少条记录开始读取数据.
func (t *tree) SearchFromTo(start int64, startEqual bool, end int64, endEqual bool, limit, offset int) []KV {
	// start at leaf node
	startNode := t.findLeafNode(start)
	if startNode == nil {
		return nil
	}

	// start position
	startPos, ok := findPos(startNode, start)
	// 如果返回的是 exact key，需要向后移动一位
	if ok && !startEqual {
		startPos++
	}

	// offset 设置
	if offset > 0 {
		startNode, startPos = findOffNode(startNode, startPos, offset)
	}

	return getResultFromStartToEnd(startNode, startPos, end, endEqual, limit)
}

// 寻找 offset 后的 startNode & startPos
func findOffNode(startNode iNode, startPos, offset int) (offNode iNode, offPos int) {
	offcount := offset
	offPos = startPos
	for loop := startNode; loop != nil; loop = loop.Link() {
		offNode = loop
		tmplc := len(loop.contents())
		if offcount >= tmplc-offPos {
			offcount -= (tmplc - offPos)
			offPos = 0
		} else {
			offPos = offcount
			break
		}
	}
	return
}

// 根据条件获取数据
func getResultFromStartToEnd(startNode iNode, startPos int, endKey int64, endEqual bool, limit int) []KV {
	var mark bool
	if limit > 0 {
		mark = true // 数据集有 limit 限制
	}

	var result []KV
	var count = limit // count 还剩多少条数据需要返回

	for loop := startNode; loop != nil; loop = loop.Link() {
		tmplc := loop.contents() // loop KV 的长度

		// 返回的 endPos 不会 > len(loop.data)
		endPos, ok := findPos(loop, endKey)
		if ok && !endEqual { // 条件设置 < end 的情况，而不是 <= end.
			endPos-- // 向前移动一位
		}

		// 如果获取到的 pos < 0 说明已经没有更多数据了.
		// 如果获取到的 pos = 0, 但是不是 exact key，说明没有更多数据了.
		if endPos < 0 || (endPos == 0 && !ok) {
			break
		}

		if !mark {
			// 如果没有 limit 限制.
			if ok {
				// 如果找到了 exact key，那么应该结束数据查找了.
				result = append(result, tmplc[startPos:endPos+1]...)
				break
			}
			// 如果没有找到 exact key，继续循环添加数据.
			result = append(result, tmplc[startPos:]...)
			startPos = 0
		} else {
			// 如果有 limit 限制.
			if count >= endPos {
				// 先找到 end key，但 limit 还没结束的情况。
				if ok {
					// 如果找到了 exact key，那么应该结束数据查找了.
					result = append(result, tmplc[startPos:endPos+1]...)
					break
				}
				// 如果没有找到 exact key，继续循环添加数据. 同时计算 limit 还剩多少.
				result = append(result, tmplc[startPos:endPos]...)
				count -= (endPos - startPos)
				startPos = 0
			} else {
				// 如果 limit 到 0，结束数据查找.
				result = append(result, tmplc[:count]...)
				break
			}
		}
	}
	return result
}

func (t *tree) Insert(key int64, value interface{}) error {
	leaf := t.findLeafNode(key)
	// leaf is nil only when root is nil.
	if leaf == nil {
		nleaf := t.newLeaf()
		nleaf.addItem(key, value, nil)
		t.root = nleaf
		t.length++
		return nil
	}

	// key 已经存在了
	_, ok := leaf.findKey(key)
	if ok {
		return ErrIndexExist
	}

	// 将 kv 添加到 leaf node
	addItemToLeafNode(leaf, key, value)

	t.length++

	return nil
}

func addItemToLeafNode(leaf iNode, key int64, value interface{}) {
	var (
		loop      iNode = leaf
		loopKey         = key
		rightNode iNode
	)

	for loop != nil {
		loopKey, rightNode = loop.addItem(loopKey, value, rightNode)
		if rightNode != nil {
			loop = loop.getParent()
		} else {
			loop = nil
		}
	}
}

func (t *tree) Delete(key int64) error {
	leaf := t.findLeafNode(key)
	if leaf == nil { // leaf == nil 说明 root == nil
		return ErrIndexNotExist
	}

	// key 不存在
	index, ok := leaf.findKey(key)
	if !ok {
		return ErrIndexNotExist
	}

	removeAndReBalance(leaf, index)

	t.length--

	// NOTE 如果有 index(internal) node 包含需要删除的 key，
	// 找到并替换成该 key 的 successor.
	n, index := t.findIndexNode(key)
	if n != nil {
		n.setKey(findSuccessorKey(n, index), index)
	}

	return nil
}

func removeAndReBalance(n iNode, ki int) {
	loop := n
	loopki := ki
	var loopci int
	for loop != nil {
		loop, loopki, loopci = loop.removeItem(loopki, loopci)
	}
}

// delete all the nodes in the tree. for GC.
func (t *tree) Clear() {
	loop := []iNode{t.root}
	for loop != nil {
		var tmp []iNode
		for _, n := range loop {
			if n.Children() != nil {
				tmp = append(tmp, n.Children()...)
			}
			n.delete()
		}
		loop = tmp
	}

	t.length = 0
	t.root = nil
}

// it will return nil only when root is nil
func (t *tree) findLeafNode(key int64) iNode {
	if t.root == nil {
		return nil
	}

	var leaf iNode
	for loop := t.root; loop != nil; {
		leaf = loop
		index, _ := loop.findKey(key)
		loop = loop.Child(index + 1)
	}
	return leaf
}

// index node is internal node which has same key of a leaf node.
// for deletion use.
func (t *tree) findIndexNode(key int64) (n iNode, index int) {
	loop := t.root
	for loop != nil && loop.Type() == Internal {
		i, ok := loop.findKey(key)
		if ok {
			return loop, i
		}
		loop = loop.Child(i + 1)
	}
	return nil, -1
}

// 这里是找 key 而不是找 node.
// for deletion use.
func findSuccessorKey(n iNode, keyIndex int) int64 {
	loop := n.Child(keyIndex + 1)
	for loop.Children() != nil {
		loop = loop.Child(0)
	}

	return loop.Key(0)
}

// NOTE 当 internal node 被拆开的时候，他的 children 需要重新绑定 parent.
// for insertion use.
func (t *tree) genNewRootNode(key int64, left, right iNode) {
	newRoot := t.newInternal()
	newRoot.keys = append(newRoot.keys, key)
	newRoot.children = append(newRoot.children, left, right)
	left.setParent(newRoot)
	right.setParent(newRoot)
	t.root = newRoot
}

// debug use only
// tree length, count all KVs in the leaf nodes
func (t *tree) len() int {
	// find most left leaf node
	loop := t.root
	for loop != nil && loop.Type() == Internal {
		loop = loop.Child(0)
	}

	// 统计每个 node 中的 KV 数量
	var count int
	for loop != nil {
		count += len(loop.contents())
		loop = loop.Link()
	}

	return count
}

// locate where the key should be in leaf node.
// true means found exact key.
func findPos(leaf iNode, key int64) (int, bool) {
	c := leaf.contents()
	for i, v := range c {
		if key == v.Key() {
			return i, true // exact key
		} else if key < v.Key() {
			return i, false
		}
	}
	return len(c), false
}

// find smallest leaf node
func (t *tree) firstLeaf() iNode {
	if t.root == nil {
		return nil
	}

	loop := t.root
	for loop.Children() != nil {
		loop = loop.Child(0)
	}

	return loop
}
