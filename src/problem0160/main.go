package problem0160

type ListNode struct {
	Val  int
	Next *ListNode
}

//采用集合的方法
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	headAListNodeMap := make(map[*ListNode]bool)
	for headA != nil {
		headAListNodeMap[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if headAListNodeMap[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := getListLen(headA), getListLen(headB)
	delta := lenA - lenB
	if lenA > lenB {
		headA = moveForward(headA, delta)
	} else {
		delta = -delta
		headB = moveForward(headB, delta)
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func getListLen(head *ListNode) int {
	listNode := head
	len := 0
	for listNode != nil {
		len++
		listNode = listNode.Next
	}
	return len
}

func moveForward(head *ListNode, len int) *ListNode {
	for len > 0 && head != nil {
		len--
		head = head.Next
	}
	return head
}
