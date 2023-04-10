package bitree

type CompareFunc[T any] func(key1, key2 *T) int

type DestroyFunc[T any] func(data *T)

// Define a structure for binary tree nodes.
type BiTreeNode[T any] struct {
	Data  *T
	Left  *BiTreeNode[T]
	Right *BiTreeNode[T]
}

// Define a structure for binary trees.
type BiTree[T any] struct {
	Size    int
	Compare CompareFunc[T]
	Destroy DestroyFunc[T]
	Root    *BiTreeNode[T]
}

// Initialize the binary tree.
func BiTreeInit[T any](tree *BiTree[T], destroy DestroyFunc[T]) {
	tree.Size = 0
	tree.Destroy = destroy
	tree.Root = nil
}

func BiTreeRemLeft[T any](tree *BiTree[T], node *BiTreeNode[T]) {
	var position **BiTreeNode[T] = nil

	// Do not allow removal from an empty tree.
	if tree.Size == 0 {
		return
	}

	// Determine where to remove nodes.
	if node == nil {
		position = &tree.Root
	} else {
		position = &node.Left
	}

	// Remove the nodes.
	if position != nil {
		BiTreeRemLeft(tree, *position)
		BiTreeRemRight(tree, *position)
		if tree.Destroy != nil {
			// Call a user-defined function to free dynamically allocated data.
			tree.Destroy((*position).Data)
		}
		// something like free() doesn't exist in Golang.
		// free(position)
		*position = nil

		// Adjust the size of the tree to account for the removed node.
		tree.Size--
	}
}

func BiTreeRemRight[T any](tree *BiTree[T], node *BiTreeNode[T]) {
	// Todo
}

func BiTreeDestroy(tree *BiTree[any]) {
	// Remove all the nodes from the tree
	BiTreeRemLeft(tree, nil)
}
