package rbtree

type Node interface {
	Index() int64       // return index of this node.
	Value() interface{} // return stored value of this node.
	Predecessor() Node  // return Predecessor node, might be nil.
	Successor() Node    // return Successor node, might be nil.
	Color() NodeColor   // return color of this node
	Tree() Tree         // return tree which this node belongs to.
	Parent() Node       // return parent node, it might be nil, if the node is root.
	LeftChild() Node    // return left child node, might be nil.
	RightChild() Node   // return right child node, might be nil.

	// internal use only
	delete()
}

type node struct {
	// rb tree 会按照 index 排序，这个值不能变
	index                         int64
	value                         interface{} // 传入的数据
	parent, leftChild, rightChild *node       // 节点的关系
	color                         NodeColor   // 节点的颜色
	tree                          *tree       // 所属 tree
}

func (n *node) Index() int64 {
	return n.index
}

func (n *node) Value() interface{} {
	return n.value
}

// left child -> right child -> right child -> right child...
// Predecessor is n's largest Left sub tree.
// Predecessor could be nil if it is not exist.
// NOTE Predecessor chould have a leftChild.
// Predecessor might not be leaf node.
// Predecessor could be left OR right child of its parent.
func (n *node) Predecessor() Node {
	r := n.predecessor()
	if r == nil {
		return nil
	}
	return r
}

func (n *node) predecessor() *node {
	var result *node
	for loop := n.leftChild; loop != nil; loop = loop.rightChild {
		result = loop
	}
	return result
}

// right child -> left child -> left child -> left child...
// Successor is n's smallest right sub tree.
// Successor could be nil if it is not exist.
// NOTE Successor chould have a rightChild.
// Successor might not be leaf node.
// Successor could be left OR right child of its parent.
func (n *node) Successor() Node {
	r := n.successor()
	if r == nil {
		return nil
	}
	return r
}

func (n *node) successor() *node {
	var result *node
	for loop := n.rightChild; loop != nil; loop = loop.leftChild {
		result = loop
	}
	return result
}

func (n *node) Color() NodeColor {
	return n.color
}

func (n *node) Tree() Tree {
	return n.tree
}

func (n *node) Parent() Node {
	r := n.parent
	if r == nil {
		return nil
	}
	return r
}

func (n *node) LeftChild() Node {
	r := n.leftChild
	if r == nil {
		return nil
	}
	return r
}

func (n *node) RightChild() Node {
	r := n.rightChild
	if r == nil {
		return nil
	}
	return r
}

func (n *node) delete() {
	n.parent = nil
	n.leftChild = nil
	n.rightChild = nil
	n.tree = nil
}
