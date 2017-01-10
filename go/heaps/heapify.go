package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Heapify struct {
	Arr []int
}

//intialisation
func (Heapify) new(size int) Heapify {
	heap := &Heapify{}
	heap.Arr = randomIntSlice(size)
	return *heap
}

//swap values
func (heap *Heapify) swap(max int, val int) {
	heap.Arr[max], heap.Arr[val] = heap.Arr[val], heap.Arr[max]
}

//heapify nodes recursively till max heap is true
func (heap *Heapify) maxHeapify(size int, index int) {
	lc := 2 * index
	rc := 2*index + 1
	max := index

	if lc < size && heap.Arr[index] < heap.Arr[lc] {
		max = lc
	}
	if rc < size && heap.Arr[max] < heap.Arr[rc] {
		max = rc
	}

	if max != index {
		heap.swap(max, index)
		heap.maxHeapify(size, max)
	}
}

//heapify nodes recursively till min heap is true
func (heap *Heapify) minHeapify(size int, index int) {
	lc := 2 * index
	rc := 2*index + 1
	min := index

	if lc < size && heap.Arr[index] > heap.Arr[lc] {
		min = lc
	}
	if rc < size && heap.Arr[min] > heap.Arr[rc] {
		min = rc
	}

	if min != index {
		heap.swap(min, index)
		heap.minHeapify(size, min)
	}
}

//build a max heap
func (heap *Heapify) buildMaxHeap() {

	for i := len(heap.Arr)/2 + 1; i > -1; i-- {
		heap.maxHeapify(len(heap.Arr), i)
	}
}

//build a min heap
func (heap *Heapify) buildMinHeap() {

	for i := len(heap.Arr)/2 + 1; i > -1; i-- {
		heap.minHeapify(len(heap.Arr), i)
	}
}

//sort the heap in asc order by creating a maxHeap
func (heap *Heapify) sortAsc() {
	end := len(heap.Arr)
	heap.buildMaxHeap()
	for i := end - 1; i >= 0; i-- {
		heap.swap(i, 0)
		heap.maxHeapify(i, 0)
	}
}

//sort the heap in desc order by creating a minHeap
func (heap *Heapify) sortDesc() {
	end := len(heap.Arr)
	heap.buildMinHeap()
	for i := end - 1; i >= 0; i-- {
		heap.swap(i, 0)
		heap.minHeapify(i, 0)
	}
}

//generate a random interger
func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

//generate a random slice of integers
func randomIntSlice(size int) []int {
	randomNumbers := make(map[int]bool)
	fmt.Println("generating random number list")
	for len(randomNumbers) < size {
		i := randomInt(1, size*6)
		if !randomNumbers[i] {
			randomNumbers[i] = true
		}
	}

	count := 0
	arr := make([]int, size)
	for k := range randomNumbers {
		arr[count] = k
		count++
	}

	return arr
}

//generate random seed
func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	size := 10

	test := &Heapify{}
	test.new(size)

	fmt.Println("-----------------------------------------")
	fmt.Printf("array before maxheap: %v\n", test.Arr)
	test.buildMaxHeap()
	fmt.Printf("array after maxheap: %v\n", test.Arr)
	fmt.Println("-----------------------------------------")
	fmt.Printf("array before minheap: %v\n", test.Arr)
	test.buildMinHeap()
	fmt.Printf("array after minheap: %v\n", test.Arr)
	fmt.Println("-----------------------------------------")
	test.sortDesc()
	fmt.Printf("array after sort desc: %v\n", test.Arr)
	fmt.Println("-----------------------------------------")
	test.sortAsc()
	fmt.Printf("array after sort asc: %v\n", test.Arr)
}
