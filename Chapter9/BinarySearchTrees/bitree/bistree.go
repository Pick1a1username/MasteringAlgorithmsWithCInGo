package bitree

import "fmt"

// Define balance factors for AVL trees.
const (
	AvlLftHeavy int = 1
	AvlBalanced int = 0
	AvlRgtHeavy int = -1
)

// Define a structure for nodes in AVL trees.
type AvlNode[T any] struct {
	Data   *T
	Hidden int
	Factor int
}

// Implement binary search trees as binary trees.
type BisTree[T any] BiTree[T]

// The following signature doesn't work.
// func rotateLeft[T AvlNode[any]](node **BiTreeNode[T]) {
func rotateLeft[T AvlNode[any]](node **BiTreeNode[AvlNode[any]]) {
	left := (*node).Left
	if left.Data.Factor == AvlLftHeavy {
		// Perform an LL rotation.
		(*node).Left = left.Right
		left.Right = *node
		(*node).Data.Factor = AvlBalanced
		left.Data.Factor = AvlBalanced
		*node = left
	} else {
		// Perform an LR rotation.
		grandChild := left.Right
		left.Right = grandChild.Left
		grandChild.Left = left
		(*node).Left = grandChild.Right
		grandChild.Right = *node

		switch grandChild.Data.Factor {
		case AvlLftHeavy:
			(*node).Data.Factor = AvlRgtHeavy
			left.Data.Factor = AvlBalanced
		case AvlBalanced:
			(*node).Data.Factor = AvlBalanced
			left.Data.Factor = AvlBalanced
		case AvlRgtHeavy:
			(*node).Data.Factor = AvlBalanced
			left.Data.Factor = AvlLftHeavy
		default:
			panic(fmt.Sprintf("unexpected Factor: %d", grandChild.Data.Factor))
		}
		grandChild.Data.Factor = AvlBalanced
		*node = grandChild
	}
}
