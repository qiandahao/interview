// https://leetcode.com/problems/reverse-linked-list-ii/
package maxInt

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	start, _ := &ListNode{}, &ListNode{}
	start.Next = head

	temp := start
	begin, end, res, prev := temp, temp, temp, temp
	count := 0
	for curr := temp; curr != nil; {
		next := curr.Next
		if count == left {
			begin = prev
			end = curr
			res = curr
		} else if count > left && count < right {
			curr.Next = res
			res = curr
		} else if count == right {
			end.Next = curr.Next
			curr.Next = res
			begin.Next = curr
			break
		}

		prev = curr
		curr = next
		count++
	}
	return start.Next
}
