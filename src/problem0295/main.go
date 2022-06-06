package problem0295

import "container/heap"

type MedianFinder struct {
	min *minHeap
	max *maxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		min: &minHeap{},
		max: &maxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.max.Len() != 0 && (*this.max)[0] >= num {
		heap.Push(this.max, num)
	} else {
		heap.Push(this.min, num)
	}
	minLen, maxLen := this.min.Len(), this.max.Len()
	if maxLen != minLen && maxLen-1 != minLen {
		if maxLen < minLen {
			val := heap.Pop(this.min)
			heap.Push(this.max, val)
		} else {
			val := heap.Pop(this.max)
			heap.Push(this.min, val)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.min.Len() == 0 && this.max.Len() == 0 {
		return 0
	}
	if this.min.Len() == this.max.Len() {
		return (float64((*this.min)[0]) + float64((*this.max)[0])) / 2
	} else {
		return float64((*this.max)[0])
	}
}

//小顶堆
type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	val := x.(int)
	*h = append(*h, val)
}
func (h *minHeap) Pop() (x interface{}) {
	var val int
	old := *h
	val, *h = old[h.Len()-1], old[:h.Len()-1]
	return val
}

//大顶堆
type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *maxHeap) Push(x interface{}) {
	val := x.(int)
	*h = append(*h, val)
}
func (h *maxHeap) Pop() (x interface{}) {
	var val int
	old := *h
	val, *h = old[h.Len()-1], old[:h.Len()-1]
	return val
}
