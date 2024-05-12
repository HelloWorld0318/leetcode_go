package lru

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func (l *List) RemoveTail() *Node {
	if l.Tail == nil {
		return nil
	}
	tail := l.Tail
	l.Tail = tail.Prev
	l.Remove(tail)
	return tail
}

func (l *List) AddToHead(key, val int) *Node {
	node := &Node{
		Key: key,
		Val: val,
	}
	if l.Tail == nil {
		l.Tail = node
	} else {
		l.Head.Prev = node
		node.Next = l.Head
	}
	l.Head = node
	return node
}

func (l *List) Remove(node *Node) {
	if node == nil {
		return
	}
	prev := node.Prev
	next := node.Next
	node.Next = nil
	node.Prev = nil
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}
	if node == l.Head {
		l.Head = next
	}
	if node == l.Tail {
		l.Tail = prev
	}
}

type LRUCache struct {
	Map  map[int]*Node
	List *List
	Cap  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Map:  map[int]*Node{},
		List: &List{},
		Cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Map[key]
	if !ok || node == nil {
		return -1
	}
	this.List.Remove(node)
	newNode := this.List.AddToHead(node.Key, node.Val)
	this.Map[key] = newNode
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Map[key]; ok {
		node.Val = value
		this.List.Remove(node)
		newNode := this.List.AddToHead(node.Key, node.Val)
		this.Map[key] = newNode
	} else {
		node := this.List.AddToHead(key, value)
		this.Map[key] = node
		for len(this.Map) > this.Cap {
			removeNode := this.List.RemoveTail()
			if removeNode != nil {
				delete(this.Map, removeNode.Key)
			}
		}
	}
}
