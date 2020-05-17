package rbtree

import (
	"errors"
)

const (
	ExistNodeErr    = "the node is already exist"
	NotExistNodeErr = "the node is not in the tree"

	RED   = true
	BLACK = false
)

const (
	DONothing byte = iota // do nothing
	RECOLOR               // needs to re-color
	RRRotation
	RLRotation
	LLRotation
	LRRotation
)

type node struct {
	parent                *node       // 上级节点
	leftChild, rightChild *node       // 左右节点
	value                 interface{} // 内容
	order                 int         // 排序号码
	color                 bool        // 颜色, RED - true / BLACK - false
	tree                  *RBTree     // 所属树
}

type RBTree struct {
	root   *node
	length int
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (t *RBTree) Add(order int, value interface{}) error {
	// empty tree
	if t.root == nil {
		t.root = &node{
			parent:     nil,
			leftChild:  nil,
			rightChild: nil,
			value:      value,
			order:      order,
			color:      BLACK,
			tree:       t,
		}
		t.length = 1
		return nil
	}

	// find where to insert the node
	parent, isLeftChild, err := t.whoseChild(order)
	if err != nil {
		return err
	}

	// 添加到 parent 下
	newNode := parent.addNewChild(value, order, isLeftChild)
	t.length++

	// check color
	t.reColorAndRotation(newNode)

	return nil
}

// true - add to leftChild , false add to rightChild
func (t *RBTree) whoseChild(order int) (*node, bool, error) {
	var result *node
	var isLeftNode bool

	for loop := t.root; loop != nil; {
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

// if parent color is BLACK, do nothing
func (t *RBTree) reColorAndRotation(_node *node) {
	for loop := _node; loop != nil; {
		switch loop.checkWhatToDo() {
		case RECOLOR:
			loop.parent.color = !loop.parent.color
			loop.parent.sibling().color = !loop.parent.sibling().color
			if !loop.parent.parent.isRootNode() {
				loop.parent.parent.color = !loop.parent.parent.color
			}

			loop = loop.parent.parent

		case LLRotation:
			loop.parent.color = !loop.parent.color
			loop.parent.parent.color = !loop.parent.parent.color
			loop.parent.parent.leftLeftRotate()
			loop = nil

		case RRRotation:
			loop.parent.color = !loop.parent.color
			loop.parent.parent.color = !loop.parent.parent.color
			loop.parent.parent.rightRightRotate()
			loop = nil

		case LRRotation:
			loop.color = !loop.color
			loop.parent.parent.color = !loop.parent.parent.color
			loop.parent.parent.leftRightRotate()
			loop = nil

		case RLRotation:
			loop.color = !loop.color
			loop.parent.parent.color = !loop.parent.parent.color
			loop.parent.parent.rightLeftRotate()
			loop = nil
		default:
			loop = nil
		}
	}
}

// find node from order number, could be nil if the order is not exist
func (t *RBTree) Find(order int) *node {
	var result *node
	for result = t.root; result != nil && result.order != order; {
		if order < result.order { // 左边
			result = result.leftChild
		} else if order > result.order { // 右边
			result = result.rightChild
		}
	}

	return result
}

// delete node
func (t *RBTree) DeleteFromOrder(order int) {
	// delete red no problem

	// delete black node / RR conflict
}

// delete RedNode
// TODO
func (t *RBTree) deleteRedNode(redNode *node) {
	switch {
	case redNode.leftChild == nil && redNode.rightChild == nil: // no child
		if redNode.isLeftChild() {
			redNode.parent.leftChild = nil
		} else {
			redNode.parent.rightChild = nil
		}
	case redNode.leftChild != nil && redNode.rightChild == nil: // has left child

	}

	redNode = nil
}
