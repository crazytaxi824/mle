package tree

// return node's value
func (n *node) Value() interface{} {
	return n.value
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

// node 所属 tree
func (n *node) Tree() *AVLTree {
	return n.tree
}

// left child -> right child -> right child -> right child...
func (n *node) LargestLeftTree() *node {
	var result *node
	for loop := n.leftChild; loop != nil; loop = loop.rightChild {
		result = loop
	}
	return result
}

// right child -> left child -> left child -> left child...
func (n *node) SmallestRightTree() *node {
	var result *node
	for loop := n.rightChild; loop != nil; loop = loop.leftChild {
		result = loop
	}
	return result
}
