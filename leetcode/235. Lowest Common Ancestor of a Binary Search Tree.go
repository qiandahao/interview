func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val == root.Val {
		return p
	} else if q.Val == root.Val {
		return q
	} else if p.Val < root.Val && q.Val > root.Val {
		return root
	} else if p.Val > root.Val && q.Val < root.Val {
		return root
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
}