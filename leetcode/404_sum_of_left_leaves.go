func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if isLeaf(root.Left) {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}

	if isLeaf(root.Right) {
		return sumOfLeftLeaves(root.Left) + 0
	}

	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func isLeaf(node *TreeNode) bool {
	if node != nil && node.Left == nil && node.Right == nil {
		return true
	}
	return false
}