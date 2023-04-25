package bitree

type CompareFunc func(key1, key2 interface{}) int

type DestroyFunc func(data interface{})

// Define a structure for binary tree nodes.
type BiTreeNode struct {
	Data  interface{}
	Left  *BiTreeNode
	Right *BiTreeNode
}

// Define a structure for binary trees.
type BiTree struct {
	Size    int
	Compare CompareFunc
	Destroy DestroyFunc
	Root    *BiTreeNode
}

// Initialize the binary tree.
func BiTreeInit(tree *BiTree, destroy DestroyFunc) {
	tree.Size = 0
	tree.Destroy = destroy
	tree.Root = nil
}

func BiTreeRemLeft(tree *BiTree, node *BiTreeNode) {
	var position **BiTreeNode = nil

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

func BiTreeRemRight(tree *BiTree, node *BiTreeNode) {
	var position **BiTreeNode = nil

	// Do not allow removal from an empty tree.
	if tree.Size == 0 {
		return
	}

	// Determine where to remove nodes.
	if node == nil {
		position = &tree.Root
	} else {
		position = &node.Right
	}

	// Remove the nodes.
	if position != nil {
		BiTreeRemLeft(tree, *position)
		BiTreeRemRight(tree, *position)
		if tree.Destroy != nil {
			// Call a user-defined function to free dynamically allocated data.
			tree.Destroy((*position).Data)
		}
		**position = BiTreeNode{}
		*position = nil

		// Adjust the size of the tree to account for the removed node.
		tree.Size--
	}
}

func BiTreeDestroy(tree *BiTree) {
	// Remove all the nodes from the tree
	BiTreeRemLeft(tree, nil)
	// No operations are allowed now, but clear the structure as a precaution.
	*tree = BiTree{}
}

// Return -1 if there's something wrong.
// Return 0 if insertion is successful.
func BiTreeInsLeft(tree *BiTree, node *BiTreeNode, data interface{}) int {
	var position **BiTreeNode = nil

	// Determine where to insert the node.
	if node == nil {
		// Allow insertion at the root only in an empty tree.
		if tree.Size > 0 {
			return -1
		}
		position = &tree.Root
	} else {
		// Normally allow insertion only at the end of a branch.
		if node.Left != nil {
			return -1
		}
		position = &node.Left
	}
	// Allocate storage for the node.
	newNode := &BiTreeNode{}

	// Insert the node into the tree.
	newNode.Data = data
	newNode.Left = nil
	newNode.Right = nil
	*position = newNode

	// Adjust the size of the tree to account for the inserted node.
	tree.Size++

	return 0
}

// Return -1 if there's something wrong.
// Return 0 if insertion is successful.
func BiTreeInsRight(tree *BiTree, node *BiTreeNode, data interface{}) int {
	var position **BiTreeNode = nil

	// Determine where to insert the node.
	if node == nil {
		// Allow insertion at the root only in an empty tree.
		if tree.Size > 0 {
			return -1
		}
		position = &tree.Root
	} else {
		// Normally allow insertion only at the end of a branch.
		if node.Right != nil {
			return -1
		}
		position = &node.Right
	}
	// Allocate storage for the node.
	newNode := &BiTreeNode{}

	// Insert the node into the tree.
	newNode.Data = data
	newNode.Left = nil
	newNode.Right = nil
	*position = newNode

	// Adjust the size of the tree to account for the inserted node.
	tree.Size++

	return 0
}

// Return -1 if there's something wrong.
// Return 0 if insertion is successful.
func BiTreeMerge(merge *BiTree, left *BiTree, right *BiTree, data interface{}) int {
	// Initialize the merged tree.
	BiTreeInit(merge, left.Destroy)
	// Insert the data for the root node of the merged tree.
	if BiTreeInsLeft(merge, nil, data) != 0 {
		BiTreeDestroy(merge)
		return -1
	}
	// Merge the two binary trees into a single binary tree.
	merge.Root.Left = left.Root
	merge.Root.Right = right.Root
	// Adjust the size of the new binary tree.
	merge.Size = merge.Size + left.Size + right.Size
	// Do not let the original trees access the merged nodes.
	left.Root = nil
	left.Size = 0
	right.Root = nil
	right.Size = 0

	return 0
}
