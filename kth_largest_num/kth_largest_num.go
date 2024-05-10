package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

// Let's build a max-heap!
//
// formula to represent in an array
// 2i + 1
// 2i + 2

type Heap struct {
	HeapArray []int
	LastItemIndex int
}

func (h *heap) New(length int) {
	h.HeapArray = make([]int, length)
	h.LastItemIndex = 0
}

func (h *Heap) Insert(item int) {
	h.HeapArray[h.LastItemIndex + 1] = item
	h.LastItemIndex++
}

func (h *Heap) Delete(item int) int {
	toPop := h.HeapArray[0]
	return toPop
}

func findKthLargest(nums []int, k int) int {
	return 0
}
