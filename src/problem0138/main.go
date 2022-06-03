package problem0138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	curNode := head
	for curNode != nil {
		node := &Node{Val: curNode.Val}
		node.Next = curNode.Next
		curNode.Next = node
		curNode = node.Next
	}

	curNode = head
	for curNode != nil {
		if curNode.Random != nil {
			curNode.Next.Random = curNode.Random.Next
		}
		curNode = curNode.Next.Next
	}

	dummy := &Node{Val: -1}

	curNode = head
	node := dummy
	for curNode != nil {
		node.Next = curNode.Next
		node = node.Next

		curNode.Next = node.Next
		curNode = curNode.Next
	}
	return dummy.Next
}
