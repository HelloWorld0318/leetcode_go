package problem0142

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	hasVisitNode := make(map[*ListNode]bool)
	for head != nil {
		if hasVisitNode[head] {
			return head
		}
		hasVisitNode[head] = true
		head = head.Next
	}
	return nil
}
