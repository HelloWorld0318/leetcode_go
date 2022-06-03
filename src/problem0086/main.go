package problem0086

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	lessThanHead, moreThanHead := &ListNode{Val: 0}, &ListNode{Val: 0}
	lessThan, moreThan := lessThanHead, moreThanHead
	for head != nil {
		if head.Val < x {
			lessThan.Next = head
			lessThan = head
		} else {
			moreThan.Next = head
			moreThan = head
		}
		head = head.Next
	}
	lessThan.Next = moreThanHead.Next
	moreThan.Next = nil
	return lessThanHead.Next
}
