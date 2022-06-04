package problem0225

type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.queue1 = append(this.queue1, x)
}

func (this *MyStack) Pop() int {
	for len(this.queue1) > 1 {
		temp := this.queue1[0]
		this.queue1 = this.queue1[1:]
		this.queue2 = append(this.queue2, temp)
	}
	var res int
	res, this.queue1 = this.queue1[0], this.queue1[1:]
	this.queue1, this.queue2 = this.queue2, this.queue1
	return res
}

func (this *MyStack) Top() int {
	res := this.Pop()
	this.Push(res)
	return res
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}
