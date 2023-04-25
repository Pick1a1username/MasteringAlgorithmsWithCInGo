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
// I had to make another struct because ...
// type BisTree[T any] BiTree[T]
type BisTree[T any] struct {
	Size    int
	Compare CompareFunc[T]
	Destroy DestroyFunc[T]
	Root    *BiTreeNode[AvlNode[T]]
}

func rotateLeft[T any](node **BiTreeNode[AvlNode[T]]) {
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

func rotateRight[T any](node **BiTreeNode[AvlNode[T]]) {
	right := (*node).Right
	if right.Data.Factor == AvlRgtHeavy {
		// Perform an RR rotation.
		(*node).Right = right.Left
		right.Left = *node
		(*node).Data.Factor = AvlBalanced
		right.Data.Factor = AvlBalanced
		*node = right
	} else {
		// Perform an RL rotation.
		grandChild := right.Left
		right.Left = grandChild.Right
		grandChild.Right = right
		(*node).Right = grandChild.Left
		grandChild.Left = *node

		switch grandChild.Data.Factor {
		case AvlLftHeavy:
			(*node).Data.Factor = AvlBalanced
			right.Data.Factor = AvlRgtHeavy
		case AvlBalanced:
			(*node).Data.Factor = AvlBalanced
			right.Data.Factor = AvlBalanced
		case AvlRgtHeavy:
			(*node).Data.Factor = AvlLftHeavy
			right.Data.Factor = AvlBalanced
		default:
			panic(fmt.Sprintf("unexpected Factor: %d", grandChild.Data.Factor))
		}
		grandChild.Data.Factor = AvlBalanced
		*node = grandChild
	}
}

func destroyLeft[T any](tree *BisTree[T], node *BiTreeNode[AvlNode[T]]) {
	var position **BiTreeNode[AvlNode[T]]
	// Do not allow destruction of an empty tree.
	if tree.Size == 0 {
		return
	}
	// Determine where to destroy nodes.
	if node == nil {
		position = &tree.Root
	} else {
		position = &node.Left
	}
	// Destroy the nodes.
	if *position != nil {
		destroyLeft(tree, *position)
		destroyRight(tree, *position)
		if tree.Destroy != nil {
			// Call a user-defined function to free dynamically allocated data.
			tree.Destroy((*position).Data.Data)
		}
		// Free the AVL data in the node, then free the node itself.
		(*position).Data = nil
		*position = nil
		// Adjust the size of the tree to account for the destroyed node.
		tree.Size--
	}
}

func destroyRight[T any](tree *BisTree[T], node *BiTreeNode[AvlNode[T]]) {
	var position **BiTreeNode[AvlNode[T]]
	// Do not allow destruction of an empty tree.
	if tree.Size == 0 {
		return
	}
	// Determine where to destroy nodes.
	if node == nil {
		position = &tree.Root
	} else {
		position = &node.Right
	}
	// Destroy the nodes.
	if *position != nil {
		destroyLeft(tree, *position)
		destroyRight(tree, *position)
		if tree.Destroy != nil {
			// Call a user-defined function to free dynamically allocated data.
			tree.Destroy((*position).Data.Data)
		}
		// Free the AVL data in the node, then free the node itself.
		(*position).Data = nil
		*position = nil
		// Adjust the size of the tree to account for the destroyed node.
		tree.Size--
	}
}

// Todo
func insert[T any](tree *BisTree[T], node **BiTreeNode[AvlNode[T]], data *T, balanced *int) int {
	// Insert the data into the tree.
	if *node != nil {
		// Handle insertion into an empty tree.
		avlData := AvlNode[T]{
			Factor: AvlBalanced,
			Hidden: 0,
			Data:   data,
		}

		return BiTreeInsLeft[T](tree, *node, avlData)

	}
	return 0
}
