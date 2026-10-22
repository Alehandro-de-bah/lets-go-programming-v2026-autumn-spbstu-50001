package main

import (
	"container/heap"
	"fmt"
)

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

const (
	minDishesNumber     = 1
	maxDishesNumber     = 10000
	minDishesRating     = -10000
	maxDishesRating     = 10000
	minPreferenceNumber = 1
)

func main() {
	var dishesNumber int
	if _, err := fmt.Scan(&dishesNumber); err != nil {
		fmt.Printf("Invalid number of dishes: %s\n", err.Error())

		return
	}

	if dishesNumber < minDishesNumber || dishesNumber > maxDishesNumber {
		fmt.Printf("Number of dishes must be in range [%d, %d]\n", minDishesNumber, maxDishesNumber)

		return
	}

	rating := make([]int, dishesNumber)
	for i := range rating {
		if _, err := fmt.Scan(&rating[i]); err != nil {
			fmt.Printf("Invalid rating of dish: %s\n", err.Error())

			return
		}

		if rating[i] < minDishesRating || rating[i] > maxDishesRating {
			fmt.Printf("Rating of dish must be in range [%d, %d]\n", minDishesRating, maxDishesRating)

			return
		}
	}

	var preferenceNumber int
	if _, err := fmt.Scan(&preferenceNumber); err != nil {
		fmt.Printf("Invalid preference number: %s\n", err.Error())

		return
	}

	if preferenceNumber < minPreferenceNumber || preferenceNumber > dishesNumber {
		fmt.Printf("Preference number must be in range [%d, %d]\n", minPreferenceNumber, dishesNumber)

		return
	}

	dishHeap := &minHeap{}
	heap.Init(dishHeap)

	for _, value := range rating {
		heap.Push(dishHeap, value)

		if dishHeap.Len() > preferenceNumber {
			heap.Pop(dishHeap)
		}
	}

	if dishHeap.Len() > 0 {
		fmt.Println((*dishHeap)[0])

		return
	}
	fmt.Println("dishHeap is empty!")
}
