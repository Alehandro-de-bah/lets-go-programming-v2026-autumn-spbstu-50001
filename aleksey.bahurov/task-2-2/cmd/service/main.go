package main

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		return
	}
	*h = append(*h, value)
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

func readDishesNumber() (int, error) {
	var dishesNumber int
	if _, err := fmt.Scan(&dishesNumber); err != nil {
		return 0, fmt.Errorf("Invalid number of dishes: %w", err)
	}

	if dishesNumber < minDishesNumber || dishesNumber > maxDishesNumber {
		return 0, fmt.Errorf("Number of dishes must be in range [%d, %d]", minDishesNumber, maxDishesNumber)
	}

	return dishesNumber, nil
}

func readRatings(dishesNumber int) ([]int, error) {
	rating := make([]int, dishesNumber)

	for index := range rating {
		if _, err := fmt.Scan(&rating[index]); err != nil {
			return nil, fmt.Errorf("Invalid rating of dish: %w", err)
		}

		if rating[index] < minDishesRating || rating[index] > maxDishesRating {
			return nil, fmt.Errorf("Rating of dish must be in range [%d, %d]", minDishesRating, maxDishesRating)
		}
	}

	return rating, nil
}

func main() {
	dishesNumber, err := readDishesNumber()
	if err != nil {
		fmt.Printf("%s\n", err.Error())

		return
	}

	rating, err := readRatings(dishesNumber)
	if err != nil {
		fmt.Printf("%s\n", err.Error())

		return
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
