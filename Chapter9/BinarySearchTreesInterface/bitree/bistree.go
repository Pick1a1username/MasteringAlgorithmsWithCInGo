package bitree

import "fmt"

// Define balance factors for AVL trees.
const (
	AvlLftHeavy int = 1
	AvlBalanced int = 0
	AvlRgtHeavy int = -1
)

// Define a structure for nodes in AVL trees.
type AvlNode struct {
	Data   interface{}
	Hidden int
	Factor int
}

type BisTree = BiTree

func rotateLeft(node **BiTreeNode) {
	left := (*node).Left
	if left.Data.(*AvlNode).Factor == AvlLftHeavy {
		// Perform an LL rotation.
		(*node).Left = left.Right
		left.Right = *node
		(*node).Data.(*AvlNode).Factor = AvlBalanced
		left.Data.(*AvlNode).Factor = AvlBalanced
		*node = left
	} else {
		// Perform an LR rotation.
		grandChild := left.Right
		left.Right = grandChild.Left
		grandChild.Left = left
		(*node).Left = grandChild.Right
		grandChild.Right = *node

		switch grandChild.Data.(*AvlNode).Factor {
		case AvlLftHeavy:
			(*node).Data.(*AvlNode).Factor = AvlRgtHeavy
			left.Data.(*AvlNode).Factor = AvlBalanced
		case AvlBalanced:
			(*node).Data.(*AvlNode).Factor = AvlBalanced
			left.Data.(*AvlNode).Factor = AvlBalanced
		case AvlRgtHeavy:
			(*node).Data.(*AvlNode).Factor = AvlBalanced
			left.Data.(*AvlNode).Factor = AvlLftHeavy
		default:
			panic(fmt.Sprintf("unexpected Factor: %d", grandChild.Data.(*AvlNode).Factor))
		}
		grandChild.Data.(*AvlNode).Factor = AvlBalanced
		*node = grandChild
	}
}

func rotateRight(node **BiTreeNode) {
	right := (*node).Right
	if right.Data.(*AvlNode).Factor == AvlRgtHeavy {
		// Perform an RR rotation.
		(*node).Right = right.Left
		right.Left = *node
		(*node).Data.(*AvlNode).Factor = AvlBalanced
		right.Data.(*AvlNode).Factor = AvlBalanced
		*node = right
	} else {
		// Perform an RL rotation.
		grandChild := right.Left
		right.Left = grandChild.Right
		grandChild.Right = right
		(*node).Right = grandChild.Left
		grandChild.Left = *node

		switch grandChild.Data.(*AvlNode).Factor {
		case AvlLftHeavy:
			(*node).Data.(*AvlNode).Factor = AvlBalanced
			right.Data.(*AvlNode).Factor = AvlRgtHeavy
		case AvlBalanced:
			(*node).Data.(*AvlNode).Factor = AvlBalanced
			right.Data.(*AvlNode).Factor = AvlBalanced
		case AvlRgtHeavy:
			(*node).Data.(*AvlNode).Factor = AvlLftHeavy
			right.Data.(*AvlNode).Factor = AvlBalanced
		default:
			panic(fmt.Sprintf("unexpected Factor: %d", grandChild.Data.(*AvlNode).Factor))
		}
		grandChild.Data.(*AvlNode).Factor = AvlBalanced
		*node = grandChild
	}
}

func destroyLeft(tree *BisTree, node *BiTreeNode) {
	var position **BiTreeNode
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
			tree.Destroy((*position).Data.(*AvlNode).Data)
		}
		// Free the AVL data in the node, then free the node itself.
		(*position).Data = nil
		*position = nil
		// Adjust the size of the tree to account for the destroyed node.
		tree.Size--
	}
}

func destroyRight(tree *BisTree, node *BiTreeNode) {
	var position **BiTreeNode
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
			tree.Destroy((*position).Data.(*AvlNode).Data)
		}
		// Free the AVL data in the node, then free the node itself.
		(*position).Data = nil
		*position = nil
		// Adjust the size of the tree to account for the destroyed node.
		tree.Size--
	}
}

func insert(tree *BisTree, node **BiTreeNode, data interface{}, balanced *int) int {
	// Insert the data into the tree.
	if *node != nil {
		// Handle insertion into an empty tree.
		avlData := AvlNode{
			Factor: AvlBalanced,
			Hidden: 0,
			Data:   data,
		}

		return BiTreeInsLeft(tree, *node, avlData)
	} else {
		// Todo
	}
	return 0
}
