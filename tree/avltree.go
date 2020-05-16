package tree

import (
	"errors"
)

const (
	ExistNodeErr    = "the node is already exist"
	NotExistNodeErr = "the node is not in the tree"
)

type node struct {
	parent                *node       // 上级节点
	leftChild, rightChild *node       // 左右节点
	depth                 int         // 自己的深度，最下层默认为1
	value                 interface{} // 内容
	order                 int         // 排序号码
	tree                  *AVLTree    // 树
}

type AVLTree struct {
	root   *node
	length int
	flag   bool // 是否允许 duplicate order
}

// duplicate value is not allowed here
func NewAVLTree(dupl ...bool) *AVLTree {
	var flag bool
	if dupl != nil {
		flag = dupl[0]
	}
	return &AVLTree{flag: flag}
}

// add node
func (avl *AVLTree) Add(order int, value interface{}) error {
	// 添加第一个节点
	if avl.root == nil {
		avl.root = &node{
			parent:     nil,
			leftChild:  nil,
			rightChild: nil,
			depth:      1,
			value:      value,
			order:      order,
			tree:       avl,
		}
		avl.length = 1
		return nil
	}

	// whose child
	parent, isLeftChild, err := avl.whoseChild(order)
	if err != nil {
		return err
	}

	// add node
	parent.addNode(value, order, isLeftChild)
	avl.length++

	for parent != nil { // TODO 可以优化不用一直检测到root
		// balance factor
		err = parent.balanceFactor(true)
		if err != nil {
			return err
		}

		parent = parent.updateDepth()
	}

	return nil
}

// true - add to leftChild , false add to rightChild
func (avl *AVLTree) whoseChild(order int) (*node, bool, error) {
	var result *node
	var isLeftNode bool

	loop := avl.root
	for loop != nil {
		if order == loop.order {
			return nil, false, errors.New(ExistNodeErr)
		}

		if order < loop.order { // left
			result = loop
			loop = loop.leftChild
			isLeftNode = true
		} else { // right
			result = loop
			loop = loop.rightChild
			isLeftNode = false
		}
	}
	return result, isLeftNode, nil
}

// could be nil if the order is not exist
func (avl *AVLTree) Find(order int) *node {
	var result *node

	for result = avl.root; result != nil && result.order != order; {
		if order < result.order { // 左边
			result = result.leftChild
		} else if order > result.order { // 右边
			result = result.rightChild
		}
	}

	return result
}

// delete node
func (avl *AVLTree) Delete(order int) error {
	delNode := avl.Find(order)
	if delNode == nil {
		return errors.New(NotExistNodeErr)
	}

	switch {
	case delNode.leftChild == nil && delNode.rightChild == nil: // has no child
		parent := delNode.parent

		if parent == nil { // root
			avl.root = nil
			delNode = nil
			return nil
		}

		// 删除自己
		if delNode.isLeftChild() {
			parent.leftChild = nil
		} else {
			parent.rightChild = nil
		}
		delNode = nil

		// 重新balance
		for parent != nil {
			// balance factor
			err := parent.balanceFactor(false)
			if err != nil {
				return err
			}

			parent = parent.updateDepth()
		}
	case delNode.leftChild != nil && delNode.rightChild == nil: // has left child
		parent := delNode.parent

		if parent == nil { // root
			avl.root = delNode.leftChild
			delNode = nil
			return nil
		}

		reBoundNodes(parent, delNode.leftChild, delNode.isLeftChild())
		delNode = nil

		for parent != nil {
			// balance factor
			err := parent.balanceFactor(false)
			if err != nil {
				return err
			}

			parent = parent.updateDepth()
		}

	case delNode.leftChild == nil && delNode.rightChild != nil: // has right child
		parent := delNode.parent

		if parent == nil { // root
			avl.root = delNode.rightChild
			delNode = nil
			return nil
		}

		reBoundNodes(parent, delNode.rightChild, delNode.isLeftChild())
		delNode = nil

		for parent != nil {
			// balance factor
			err := parent.balanceFactor(false)
			if err != nil {
				return err
			}

			parent = parent.updateDepth()
		}

	case delNode.leftChild != nil && delNode.rightChild != nil: // has both child
		// find replace
		replaceNode := delNode.LargestLeftTree()

		// 替换value，order，不替换 depth，left，right child
		delNode.order = replaceNode.order
		delNode.value = replaceNode.value

		// 删除 replaceNode
		parent := replaceNode.parent
		if replaceNode.isLeftChild() {
			parent.leftChild = nil
		} else {
			parent.rightChild = nil
		}
		replaceNode = nil

		for parent != nil {
			// balance factor
			err := parent.balanceFactor(false)
			if err != nil {
				return err
			}

			parent = parent.updateDepth()
		}
	}

	return nil
}
