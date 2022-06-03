package problem0023

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeKListsRange(lists, 0, len(lists)-1)
}

func mergeKListsRange(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	mid := (start + end) / 2
	left := mergeKListsRange(lists, start, mid)
	rigit := mergeKListsRange(lists, mid+1, end)
	return mergeTwoList(left, rigit)
}

func mergeTwoList(headA *ListNode, headB *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	node := dummy
	for headA != nil || headB != nil {
		if headA == nil || (headB != nil && headB.Val < headA.Val) {
			node.Next = headB
			headB = headB.Next
		} else {
			node.Next = headA
			headA = headA.Next
		}
		node = node.Next
	}
	return dummy.Next
}
