package main

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

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	oldHeap := *h
	oldLength := len(oldHeap)
	lastElement := oldHeap[oldLength-1]
	*h = oldHeap[0 : oldLength-1]
	return lastElement
}
