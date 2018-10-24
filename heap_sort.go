package main

import (
	"flag"
	"fmt"
	"strconv"
)

type Heap struct {
	source []int
	length int
	size   int
}

func (heap Heap) parent(index int) int {
	if index == 0 {
		return -1
	}
	parentIndex := (index+1)/2 - 1
	return parentIndex
}

func (heap Heap) left(index int) int {
	leftIndex := 2*index + 1
	if leftIndex >= heap.size {
		return -1
	}
	return leftIndex
}

func (heap Heap) right(index int) int {
	rightIndex := 2*index + 2
	if rightIndex >= heap.size {
		return -1
	}
	return rightIndex
}

func (heap Heap) maxHeapify(index int, show bool) {
	largestIndex := index
	leftIndex := heap.left(index)
	rightIndex := heap.right(index)
	if show {
		fmt.Println("max heapifing: index", index, "left index", leftIndex, "right index", rightIndex)
	}

	if leftIndex != -1 && leftIndex < heap.size && heap.source[leftIndex] > heap.source[index] {
		if show {
			fmt.Println("max heapifing: left value", heap.source[leftIndex], "is more than this value", heap.source[index])
		}
		largestIndex = leftIndex
	}
	if rightIndex != -1 && rightIndex < heap.size && heap.source[rightIndex] > heap.source[largestIndex] {
		if show {
			fmt.Println("max heapifing: right value", heap.source[rightIndex], "is more than this value", heap.source[largestIndex])
		}
		largestIndex = rightIndex
	}
	if largestIndex != index {
		if show {
			fmt.Println("max heapifing: source", heap.source)
			fmt.Println("max heapifing: exchange value", heap.source[largestIndex], "and", heap.source[index])
		}
		heap.source[index], heap.source[largestIndex] = heap.source[largestIndex], heap.source[index]
		heap.maxHeapify(largestIndex, show)
		if show {
			fmt.Println("max heapifing: exchange result", heap.source)
		}
	} else {
		if show {
			fmt.Println("max heapifing: nothing to do")
		}
	}
}

func (heap Heap) minHeapify(index int, show bool) {
	smallestIndex := index
	leftIndex := heap.left(index)
	rightIndex := heap.right(index)
	if show {
		fmt.Println("min heapifing: index", index, "left index", leftIndex, "right index", rightIndex)
	}

	if leftIndex != -1 && leftIndex < heap.size && heap.source[leftIndex] < heap.source[index] {
		if show {
			fmt.Println("min heapifing: left value", heap.source[leftIndex], "is less than this value", heap.source[index])
		}
		smallestIndex = leftIndex
	}
	if rightIndex != -1 && rightIndex < heap.size && heap.source[rightIndex] < heap.source[smallestIndex] {
		if show {
			fmt.Println("min heapifing: right value", heap.source[rightIndex], "is less than this value", heap.source[smallestIndex])
		}
		smallestIndex = rightIndex
	}
	if smallestIndex != index {
		if show {
			fmt.Println("min heapifing: source", heap.source)
			fmt.Println("min heapifing: exchange value", heap.source[smallestIndex], "and", heap.source[index])
		}
		heap.source[index], heap.source[smallestIndex] = heap.source[smallestIndex], heap.source[index]
		heap.minHeapify(smallestIndex, show)
		if show {
			fmt.Println("min heapifing: exchange result", heap.source)
		}
	} else {
		if show {
			fmt.Println("min heapifing: nothing to do")
		}
	}
}

func buildMaxHeap(source []int, show bool) Heap {
	if show {
		fmt.Println("build: start build a max heap", source)
	}
	heap := Heap{source: source, length: len(source), size: len(source)}
	maxParentIndex := heap.size/2 - 1
	for i := maxParentIndex; i >= 0; i-- {
		heap.maxHeapify(i, show)
	}
	if show {
		fmt.Println("build: build a new max heap", heap.source)
	}
	return heap
}

func buildMinHeap(source []int, show bool) Heap {
	if show {
		fmt.Println("build: start build a min heap", source)
	}
	heap := Heap{source: source, length: len(source), size: len(source)}
	minParentIndex := heap.size/2 - 1
	for i := minParentIndex; i >= 0; i-- {
		heap.minHeapify(i, show)
	}
	if show {
		fmt.Println("build: build a new min heap", heap.source)
	}
	return heap
}

func heapSort(source []int, show, desc bool) {
	if desc {
		heap := buildMinHeap(source, show)
		for i := heap.length - 1; i > 0; i-- {
			if show {
				fmt.Println("sort: exchange first value", heap.source[0], "and this value", heap.source[i])
			}
			heap.source[0], heap.source[i] = heap.source[i], heap.source[0]
			heap.size--
			if show {
				fmt.Println("sort: source", heap.source)
			}
			heap.minHeapify(0, show)
		}
	} else {
		heap := buildMaxHeap(source, show)
		for i := heap.length - 1; i > 0; i-- {
			if show {
				fmt.Println("sort: exchange first value", heap.source[0], "and this value", heap.source[i])
			}
			heap.source[0], heap.source[i] = heap.source[i], heap.source[0]
			heap.size--
			if show {
				fmt.Println("sort: source", heap.source)
			}
			heap.maxHeapify(0, show)
		}
	}
}

func main() {
	show := flag.Bool("s", false, "show detail flag")
	desc := flag.Bool("d", false, "desc flag")
	flag.Parse()

	source := []int{}
	for _, arg := range flag.Args() {
		i, _ := strconv.Atoi(arg)
		source = append(source, i)
	}

	heapSort(source, *show, *desc)
	fmt.Println(source)
}
