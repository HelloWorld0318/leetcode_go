package problem0142

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle1(head *ListNode) *ListNode {
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

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	var meetNode *ListNode
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if slow == fast {
			meetNode = slow
			break
		}
	}
	if meetNode == nil {
		return nil
	}
	for head != nil && meetNode != nil {
		if head == meetNode {
			return head
		}
		head = head.Next
		meetNode = meetNode.Next
	}
	return nil
}
