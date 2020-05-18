package rbtree

import (
	"errors"
)

const (
	ErrNodeExist    = "the node is already exist"
	ErrNodeNotExist = "the node is not in the tree"

	RED   = true
	BLACK = false
)

const (
	doNothing byte = iota // do nothing
	reColor               // needs to re-color
	rrRotation
	rlRotation
	llRotation
	lrRotation
)

type node struct {
	parent                *node       // 上级节点
	leftChild, rightChild *node       // 左右节点
	value                 interface{} // 内容
	order                 int         // 排序号码
	color                 bool        // 颜色, RED - true / BLACK - false
	tree                  *rbTree     // 所属树
}

type rbTree struct {
	root   *node
	length int
}

func NewTree() *rbTree {
	return &rbTree{}
}

func (t *rbTree) Add(order int, value interface{}) error {
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
	t.addNodeReColorAndRotation(newNode)

	return nil
}

// true - add to leftChild , false add to rightChild
func (t *rbTree) whoseChild(order int) (*node, bool, error) {
	var result *node
	var isLeftNode bool

	for loop := t.root; loop != nil; {
		if order == loop.order {
			return nil, false, errors.New(ErrNodeExist)
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

// recolor and rotation when add note to the tree
func (t *rbTree) addNodeReColorAndRotation(_node *node) {
	for loop := _node; loop != nil; {
		switch loop.checkWhatToDo() {
		case reColor:
			loop.parent.color = !loop.parent.color
			loop.parent.sibling().color = !loop.parent.sibling().color
			if !loop.parent.parent.isRootNode() {
				loop.parent.parent.color = !loop.parent.parent.color
			}

			loop = loop.parent.parent

		case llRotation:
			loop.parent.color = !loop.parent.color
			loop.parent.parent.color = !loop.parent.parent.color
			loop.parent.parent.leftLeftRotate()
			loop = nil

		case rrRotation:
			loop.parent.color = !loop.parent.color
			loop.parent.parent.color = !loop.parent.parent.color
			loop.parent.parent.rightRightRotate()
			loop = nil

		case lrRotation:
			loop.color = !loop.color
			loop.parent.parent.color = !loop.parent.parent.color
			loop.parent.parent.leftRightRotate()
			loop = nil

		case rlRotation:
			loop.color = !loop.color
			loop.parent.parent.color = !loop.parent.parent.color
			loop.parent.parent.rightLeftRotate()
			loop = nil

		default: // Do Nothing
			loop = nil
		}
	}
}

// find node from order number, could be nil if the order is not exist
func (t *rbTree) Find(order int) *node {
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
func (t *rbTree) DeleteFromOrder(order int) error {
	// find the node needs to be deleted
	delNode := t.Find(order)
	if delNode == nil {
		return errors.New(ErrNodeNotExist)
	}

	// deletion cases
	for loop := delNode; loop != nil; {
		// find a predecessor or successor,
		// if it's nil means no child
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

			// DOUBLE BLACK situation
			t.casesOfDoubleBlackSituation(loop)

			// delete node
			if loop.isLeftChild() {
				loop.parent.leftChild = nil
			} else {
				loop.parent.rightChild = nil
			}
			loop = nil

		case ps != nil: // ps is not nil, means ps is the replace node to be deleted
			loop.value = ps.value
			loop.order = ps.order

			loop = ps
		}
	}

	t.length--
	return nil
}

func (t *rbTree) Delete(n *node) error {
	return t.DeleteFromOrder(n.order)
}

// 6 cases of the DOUBLE BLACK situation
func (t *rbTree) casesOfDoubleBlackSituation(blackLeaf *node) {
	for doubleBlack := blackLeaf; doubleBlack != nil && !doubleBlack.isRootNode(); {
		sibling := doubleBlack.sibling()

		switch {
		// case doubleBlack.isRootNode(): // double black is root node
		// 	doubleBlack = nil

		case doubleBlack.parent.color == RED &&
			sibling.color == BLACK && sibling.bothChildrenAreBlack():

			// swap color
			doubleBlack.parent.color, sibling.color = sibling.color, doubleBlack.parent.color

			doubleBlack = nil

		case doubleBlack.parent.color == BLACK &&
			sibling.color == BLACK && sibling.bothChildrenAreBlack():
			// recolor
			sibling.color = RED

			doubleBlack = doubleBlack.parent

		case doubleBlack.parent.color == BLACK &&
			sibling.color == RED && sibling.bothChildrenAreBlack():

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

// total number of nodes in the tree
func (t *rbTree) Size() int {
	return t.length
}

// root node of the tree
func (t *rbTree) Root() *node {
	return t.root
}

// smallest node in the tree
func (t *rbTree) Smallest() *node {
	var smallest *node
	for loop := t.root; loop != nil; loop = loop.leftChild {
		smallest = loop
	}
	return smallest
}

// biggest node in the tree
func (t *rbTree) Biggest() *node {
	var biggest *node
	for loop := t.root; loop != nil; loop = loop.rightChild {
		biggest = loop
	}
	return biggest
}

// sort the nodes in ASC order
func (t *rbTree) Sort() []*node {
	result := make([]*node, 0, t.length)
	smallest := t.Smallest()

	// s -> small right tree
	for loop := smallest; loop != nil; {
		result = append(result, loop)
		if loop.Successor() != nil { // 先找自己 smallest right
			loop = loop.Successor()
		} else { // 再找 parent
			if loop.parent != nil && loop.order < loop.parent.order { // 在左边
				loop = loop.parent
			} else if loop.parent != nil && loop.order >= loop.parent.order {
				// 如果自己是 right child 继续向上找一直到 left child
				loop = loop.findLeftParent()
			} else {
				break
			}
		}
	}

	return result
}

// for sort，一直寻找 parent，直到自己是 parent 的 left child，返回 parent，
// 如果自己是 parent 的 right child，继续向上寻找。
func (n *node) findLeftParent() *node {
	for loop := n; ; {
		if loop.parent == nil {
			return nil
		}
		if loop.order >= loop.parent.order {
			loop = loop.parent
		} else {
			return loop.parent
		}
	}
}
