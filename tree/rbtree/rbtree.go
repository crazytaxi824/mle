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
// TODO
func (t *RBTree) DeleteFromOrder(order int) error {
	delNode := t.Find(order)
	if delNode == nil {
		return errors.New(NotExistNodeErr)
	}

	// deletion cases
	for loop := delNode; loop != nil; {
		ps := loop.predecessorOrSuccessor()
		switch {
		case ps == nil && loop.color == RED: // red leaf - means no child
			// delete node
			if loop.isLeftChild() {
				loop.parent.leftChild = nil
			} else {
				loop.parent.rightChild = nil
			}
			loop = nil

		case ps == nil && loop.color == BLACK: // black leaf - means no child
			if loop.isRootNode() { // root situation
				t.root = nil
				t.length--
				return nil
			}

			// double black situation
			t.recolorOrRotateDoubleBlackLeaf(loop)

			// delete node
			if loop.isLeftChild() {
				loop.parent.leftChild = nil
			} else {
				loop.parent.rightChild = nil
			}
			loop = nil

		case ps != nil:
			loop.value = ps.value
			loop.order = ps.order

			loop = ps
		}
	}

	t.length--
	return nil
}

func (t *RBTree) recolorOrRotateDoubleBlackLeaf(blackLeaf *node) {
	for doubleBlack := blackLeaf; doubleBlack != nil; {
		sibling := doubleBlack.sibling()

		switch {
		case doubleBlack.isRootNode(): // double black is root node
			doubleBlack = nil

		case doubleBlack.parent.color == RED &&
			sibling.color == BLACK && sibling.childrenAreBlack():

			// swap color
			doubleBlack.parent.color, sibling.color = sibling.color, doubleBlack.parent.color

			doubleBlack = nil

		case doubleBlack.parent.color == BLACK &&
			sibling.color == BLACK && sibling.childrenAreBlack():
			// recolor
			sibling.color = RED

			doubleBlack = doubleBlack.parent

		case doubleBlack.parent.color == BLACK &&
			sibling.color == RED && sibling.childrenAreBlack():

			// swap color
			doubleBlack.parent.color, sibling.color = sibling.color, doubleBlack.parent.color

			// rotation
			if doubleBlack.isLeftChild() {
				doubleBlack.parent.rightRightRotate()
			} else {
				doubleBlack.parent.leftLeftRotate()
			}
			// doubleBlack = doubleBlack, double black is still here, apply other cases

		case doubleBlack.parent.color == BLACK &&
			sibling.color == BLACK &&
			(doubleBlack.farSideOfTheNephew() == nil || doubleBlack.farSideOfTheNephew().color == BLACK) &&
			(doubleBlack.nearSideOfTheNephew() != nil && doubleBlack.nearSideOfTheNephew().color == RED):
			// recolor
			doubleBlack.nearSideOfTheNephew().color = BLACK
			sibling.color = RED

			// rotation
			if doubleBlack.isLeftChild() {
				sibling.leftLeftRotate()
			} else {
				sibling.rightRightRotate()
			}
			// doubleBlack = doubleBlack, double black is still here, apply other cases

		case sibling.color == BLACK &&
			doubleBlack.farSideOfTheNephew() != nil && doubleBlack.farSideOfTheNephew().color == RED:

			// recolor far side of its nephew
			doubleBlack.farSideOfTheNephew().color = BLACK

			// swap color
			doubleBlack.parent.color, sibling.color = sibling.color, doubleBlack.parent.color

			// rotation
			if doubleBlack.isLeftChild() {
				doubleBlack.parent.rightRightRotate()
			} else {
				doubleBlack.parent.leftLeftRotate()
			}

			doubleBlack = nil
		}
	}
}

// children are black color
func (n *node) childrenAreBlack() bool {
	return (n.leftChild == nil || n.leftChild.color == BLACK) &&
		(n.rightChild == nil || n.rightChild.color == BLACK)
}

func (n *node) farSideOfTheNephew() *node {
	if n.isLeftChild() {
		return n.sibling().rightChild
	}
	return n.sibling().leftChild
}

func (n *node) nearSideOfTheNephew() *node {
	if n.isLeftChild() {
		return n.sibling().leftChild
	}
	return n.sibling().rightChild
}
