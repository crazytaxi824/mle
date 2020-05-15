package tree

import (
	"errors"
)

const (
	ExistNodeErr = "the node is already exist"
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
	whose, isLeftChild, err := avl.whoseChild(order)
	if err != nil {
		return err
	}

	// add node
	whose.addNode(value, order, isLeftChild)
	avl.length++

	for whose != nil {
		// balance factor
		err = whose.balanceFactor()
		if err != nil {
			return err
		}

		whose = whose.updateDepth()
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
