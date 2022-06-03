package problem0206

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//非递归的方法
func reverseListMethod1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dumpy := &ListNode{Val: -1}
	dumpy.Next = head
	cur, next := head, head.Next
	for next != nil {
		cur.Next = next.Next
		next.Next = dumpy.Next
		dumpy.Next = next

		next = cur.Next
	}
	return dumpy.Next
}

//非递归的方法
func reverseListMethod2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead, next := head, head.Next
	for next != nil {
		head.Next = next.Next
		next.Next = newHead
		newHead = next

		next = head.Next
	}
	return newHead
}
