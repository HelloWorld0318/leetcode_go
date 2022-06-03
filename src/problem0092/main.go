package problem0092

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//1.记录翻转部分头节点
	var preHead *ListNode
	changeLen := right - left + 1
	result := head
	for left > 1 && head != nil {
		preHead = head
		head = head.Next
		left--
	}

	//2.记录翻倍部分的尾部节点
	newListTail := head
	var newHead *ListNode

	//3.翻转
	for changeLen > 0 && head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head

		head = next
		changeLen--
	}

	//4.拼接链表
	newListTail.Next = head
	if preHead == nil {
		result = newHead
	} else {
		preHead.Next = newHead
	}
	return result
}
