package problem0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, dummy := 0, &ListNode{Val: -1}
	cur := dummy
	for l1 != nil || l2 != nil || carry > 0 {
		valueOfL1, valueOfL2 := 0, 0
		if l1 != nil {
			valueOfL1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			valueOfL2 = l2.Val
			l2 = l2.Next
		}
		sum := valueOfL1 + valueOfL2 + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else if carry > 1 {
			carry = 0
		}
		node := &ListNode{
			Val: sum,
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}
