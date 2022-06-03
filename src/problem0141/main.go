package problem0141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	hasVisitNode := make(map[*ListNode]bool)
	for head != nil {
		if hasVisitNode[head] {
			return true
		}
		hasVisitNode[head] = true
		head = head.Next
	}
	return false
}
