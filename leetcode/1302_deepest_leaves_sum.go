/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func deepestLeavesSum(root *TreeNode) int {
	queueP := make([]*TreeNode, 0)
	queueQ := make([]*TreeNode, 0)
	queueP = append(queueP, root)
	clearP := true
	sum, res := 0, 0
	for len(queueP) != 0 || len(queueQ) != 0 {
		if clearP {
			if len(queueP) == 0 {
				clearP = false
				if sum != 0 {
					res = sum
				}
				sum = 0
			} else {
				temp := queueP[0]
				sum += temp.Val
				n := len(queueP)
				queueP = queueP[1:n]
				if temp.Left != nil {
                    queueQ = append(queueQ, temp.Left)
				}
				if temp.Right != nil {
                    queueQ = append(queueQ, temp.Right)
				}
			}
		} else {
			if len(queueQ) == 0 {
				clearP = true
				if sum != 0 {
					res = sum
				}
				sum = 0
			} else {
				temp := queueQ[0]
				n := len(queueQ)
				sum += temp.Val
				queueQ = queueQ[1:n]
				if temp.Left != nil {
					queueP = append(queueP, temp.Left)
				}
				if temp.Right != nil {
					queueP = append(queueP, temp.Right)
				}
			}
		}
		
	}
     if sum != 0 {
         return sum
     } else {
         return res
     }
}