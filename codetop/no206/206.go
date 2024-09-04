package no206

type ListNode struct {
	Val  int
	Next *ListNode
}

// 增加一个新的头部，然后开始反转
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
