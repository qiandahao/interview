/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func deepestLeavesSum(root *TreeNode) int {
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    sum := 0
    for len(queue) != 0 {
        sum = 0
        index := len(queue)
        for i:=0; i < index; i++ {
            temp := queue[i]
            sum += temp.Val
            if temp.Left != nil {
                queue = append(queue, temp.Left)
            }
            if temp.Right != nil{
                queue = append(queue, temp.Right)
            }
        }
        n := len(queue)
        queue = queue[index:n]
    }
    return sum
}