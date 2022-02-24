/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    res := make([]int, 0)
    curr, jump := head, head
    for curr != nil && jump.Next != nil {
        jump = jump.Next.Next
        if jump == nil {
            res = append(res, curr.Val)
            curr = curr.Next
            break
        }
        res = append(res, curr.Val)
        curr = curr.Next
        
    }
    if jump !=nil && jump.Next == nil {
        curr = curr.Next
    }
    
    for i := len(res)-1; i>=0; i-- {
        if res[i] != curr.Val {
            return false
        }
        curr = curr.Next
    }
    return true
}