package problem0215

import "container/heap"

//创建一个小顶堆
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

func (h *minHeap) Pop() interface{} {
	var val int
	old := *h
	val, *h = old[h.Len()-1], old[:h.Len()-1]
	return val
}

func findKthLargest(nums []int, k int) int {
	h := &minHeap{}
	heap.Init(h)
	for _, val := range nums {
		heap.Push(h, val)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return heap.Pop(h).(int)
}
