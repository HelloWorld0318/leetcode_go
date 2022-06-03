package problem0138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList1(head *Node) *Node {
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

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*Node]int)
	nodeList := make([]*Node, 0)
	cur, index := head, 0
	for cur != nil {
		nodeMap[cur] = index
		nodeList = append(nodeList, &Node{Val: cur.Val})
		index++
		cur = cur.Next
	}

	cur = head
	i := 0
	for cur != nil {
		if i+1 < len(nodeList) {
			nodeList[i].Next = nodeList[i+1]
		}
		if cur.Random != nil {
			index := nodeMap[cur.Random]
			nodeList[i].Random = nodeList[index]
		}
		i++
		cur = cur.Next
	}
	return nodeList[0]
}