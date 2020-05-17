package tree

// return node's value
func (n *node) Value() interface{} {
	return n.value
}

// change Value
func (n *node) ChangeValue(v interface{}) {
	n.value = v
}

// return node's order
func (n *node) Order() int {
	return n.order
}

// return node's depth
func (n *node) Depth() int {
	return n.depth
}

// could be nil
func (n *node) LeftChild() *node {
	return n.leftChild
}

// could be nil
func (n *node) RightChild() *node {
	return n.rightChild
}

// if node is root , it will return nil
func (n *node) Parent() *node {
	return n.parent
}

// the tree which the node belongs to
func (n *node) Tree() *AVLTree {
	return n.tree
}

// left child -> right child -> right child -> right child...
// largest Left sub tree
func (n *node) Predecessor() *node {
	var result *node
	for loop := n.leftChild; loop != nil; loop = loop.rightChild {
		result = loop
	}
	return result
}

// right child -> left child -> left child -> left child...
// smallest right sub tree
func (n *node) Successor() *node {
	var result *node
	for loop := n.rightChild; loop != nil; loop = loop.leftChild {
		result = loop
	}
	return result
}
