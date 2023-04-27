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
	if *node == nil {
		// Handle insertion into an empty tree.
		avlData := AvlNode{
			Factor: AvlBalanced,
			Hidden: 0,
			Data:   data,
		}

		return BiTreeInsLeft(tree, *node, avlData)
	} else {
		// Handle insertion into a tree that is not empty.
		cmpVal := tree.Compare(data, (*node).Data)

		if cmpVal < 0 {
			// Move to the left.
			if (*node).Left == nil {
				avlData := AvlNode{
					Factor: AvlBalanced,
					Hidden: 0,
					Data:   data,
				}
				if BiTreeInsLeft(tree, *node, avlData) != 0 {
					return -1
				}
				*balanced = 0
			} else {
				if retVal := insert(tree, &(*node).Left, data, balanced); retVal != 0 {
					return retVal
				}
			}
			// Ensure that the tree remains balanced.
			if *balanced == 0 {
				switch (*node).Data.(*AvlNode).Factor {
				case AvlLftHeavy:
					rotateLeft(node)
					*balanced = 1
				case AvlBalanced:
					(*node).Data.(*AvlNode).Factor = AvlLftHeavy
				case AvlRgtHeavy:
					(*node).Data.(*AvlNode).Factor = AvlBalanced
					*balanced = 1
				default:
					panic(fmt.Sprintf("unexpected Factor: %d", (*node).Data.(*AvlNode).Factor))
				}
			}
		} else if cmpVal > 0 {
			// Move to the right.
			if (*node).Right == nil {
				avlData := AvlNode{
					Factor: AvlBalanced,
					Hidden: 0,
					Data:   data,
				}
				if BiTreeInsRight(tree, *node, avlData) != 0 {
					return -1
				}
				*balanced = 0
			} else {
				if retVal := insert(tree, &(*node).Right, data, balanced); retVal != 0 {
					return retVal
				}
			}
			// Ensure that the tree remains balanced.
			if *balanced == 0 {
				switch (*node).Data.(*AvlNode).Factor {
				case AvlLftHeavy:
					(*node).Data.(*AvlNode).Factor = AvlBalanced
					*balanced = 1
				case AvlBalanced:
					(*node).Data.(*AvlNode).Factor = AvlLftHeavy
				case AvlRgtHeavy:
					rotateRight(node)
					*balanced = 1
				default:
					panic(fmt.Sprintf("unexpected Factor: %d", (*node).Data.(*AvlNode).Factor))
				}
			}
		} else {
			// Handle finding a copy of the data.
			if (*node).Data.(*AvlNode).Hidden == 0 {
				// Do nothing since the data is in the tree and not hidden.
				return -1
			} else {
				// Insert the new data and mark it as not hidden.
				if tree.Destroy != nil {
					// Destroy the hidden data since it is being replaced.
					tree.Destroy((*node).Data.(*AvlNode))
				}
				(*node).Data.(*AvlNode).Data = data
				(*node).Data.(*AvlNode).Hidden = 0
				// Do not rebalance because the tree structure is unchanged.
				*balanced = -1
			}
		}
	}
	return 0
}

func hide(tree *BisTree, node *BiTreeNode, data interface{}) int {
	if node == nil {
		// Return that the data was not found.
		return -1
	}
	cmpVal := tree.Compare(data, node.Data.(*AvlNode))
	if cmpVal < 0 {
		// Move to the left.
		return hide(tree, node.Left, data)
	} else if cmpVal > 0 {
		// Move to the right.
		return hide(tree, node.Right, data)
	} else {
		// Mark the node as hidden.
		node.Data.(*AvlNode).Hidden = 1
		return 0
	}
}

func lookup(tree *BisTree, node *BiTreeNode, data *interface{}) int {
	if node == nil {
		// Return that the data was not found.
		return -1
	}
	cmpVal := tree.Compare(*data, node.Data.(*AvlNode))
	if cmpVal < 0 {
		// Move to the left.
		return lookup(tree, node.Left, data)
	} else if cmpVal > 0 {
		// Move to the right.
		return lookup(tree, node.Right, data)
	} else {
		if node.Data.(*AvlNode).Hidden == 0 {
			// Pass back the data from the tree.
			// Is it ok?
			*data = node.Data.(*AvlNode).Data
			return 0
		} else {
			// Return that the data was not found.
			return -1
		}
	}
}

func BisTreeInit(tree *BisTree, compare *CompareFunc, destroy *DestroyFunc) {
	// Initialize the tree.
	BiTreeInit(tree, *destroy)
	tree.Compare = compare
}

func BisTreeDestroy(tree *BisTree) {
	// Todo
}

func BisTreeInsert(tree *BisTree, data interface{}) int {
	// Todo
	return 0
}

func BisTreeRemove(tree *BisTree, data interface{}) int {
	// Todo
	return 0
}

func BisTreeLookup(tree *BisTree, data *interface{}) int {
	// Todo
	return 0
}
