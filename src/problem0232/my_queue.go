package problem0232

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) != 0 {
			var temp int
			temp, this.stack1 = this.stack1[len(this.stack1)-1], this.stack1[:len(this.stack1)-1]
			this.stack2 = append(this.stack2, temp)
		}
	}
	var res int
	this.stack2, res = this.stack2[:len(this.stack2)-1], this.stack2[len(this.stack2)-1]
	return res
}

func (this *MyQueue) Peek() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) != 0 {
			var temp int
			temp, this.stack1 = this.stack1[len(this.stack1)-1], this.stack1[:len(this.stack1)-1]
			this.stack2 = append(this.stack2, temp)
		}
	}
	return this.stack2[len(this.stack2)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}
