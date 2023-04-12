package bitree

import l "BinarySearchTrees/list"

func Preorder[T any](node *BiTreeNode[T], list *l.List[T]) int {
	// Load the list with a preorder listing of the tree.
	if node != nil {
		if l.ListInsNext(list, list.Tail, node.Data) != 0 {
			return -1
		}
		if node.Left != nil {
			if Preorder(node.Left, list) != 0 {
				return -1
			}
		}
		if node.Right != nil {
			if Preorder(node.Right, list) != 0 {
				return -1
			}
		}
	}
	return 0
}

func Inorder[T any](node *BiTreeNode[T], list *l.List[T]) int {
	// Load the list with an inorder listing of the tree.
	if node != nil {
		if node.Left != nil {
			if Inorder(node.Left, list) != 0 {
				return -1
			}
		}
		if l.ListInsNext(list, list.Tail, node.Data) != 0 {
			return -1
		}
		if node.Right != nil {
			if Inorder(node.Right, list) != 0 {
				return -1
			}
		}
	}
	return 0
}

func Postorder[T any](node *BiTreeNode[T], list *l.List[T]) int {
	// Load the list with a postorder listing of the tree.
	if node != nil {
		if node.Left != nil {
			if Postorder(node.Left, list) != 0 {
				return 0
			}
		}
		if node.Right != nil {
			if Postorder(node.Right, list) != 0 {
				return -1
			}
		}
		if l.ListInsNext(list, list.Tail, node.Data) != 0 {
			return -1
		}
	}
	return 0
}
