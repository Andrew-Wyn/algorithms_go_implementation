// heap sort, go implementation

package main

import (
	"fmt"
)

func parent(i int) int {
	return (i >> 1)
}

func left(i int) int {
	return (i << 1)
}

func right(i int) int {
	return (i << 1) ^ 1
}

type Heap struct {
	value    []int
	heapSize int
}

// O(log n)
func (c *Heap) maxHeapify(i int) {
	l := left(i)
	r := right(i)

	var massimo int = -1

	if l < c.heapSize && c.value[l] > c.value[i] {
		massimo = l
	} else {
		massimo = i
	}

	if r < c.heapSize && c.value[r] > c.value[massimo] {
		massimo = r
	}

	if massimo != i {
		c.change(i, massimo)
		c.maxHeapify(massimo)
	}
}

// O(n)
func (c *Heap) buildMaxHeap() {
	c.heapSize = len(c.value)

	for i := int(len(c.value)/2) - 1; i >= 0; i-- {
		c.maxHeapify(i)
	}
}

func (c *Heap) sort() {
	c.buildMaxHeap()

	for i := len(c.value) - 1; i >= 1; i-- {
		c.change(0, i)
		c.heapSize--
		c.maxHeapify(0)
	}
}

func (c *Heap) change(a int, b int) {
	temp := c.value[a]
	c.value[a] = c.value[b]
	c.value[b] = temp
}

func main() {
	fmt.Println("*** HEAP-SORT ***")

	var heap Heap = Heap{value: []int{4, 2, 3, 4, 2, 3}, heapSize: 0}

	fmt.Println(heap)

	heap.sort()

	fmt.Println(heap)

}
