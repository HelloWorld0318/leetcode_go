package problem0141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle1(head *ListNode) bool {
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

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if slow == fast {
			return true
		}
	}
	return false
}
