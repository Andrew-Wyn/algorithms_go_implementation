// heap sort, go implementation

package main

import (
	"fmt"
)

func Parent(i int) int {
	i++
	return (i >> 1) - 1
}

func Left(i int) int {
	i++
	return (i << 1) - 1
}

func Right(i int) int {
	i++
	return (i << 1)
}

type Heap struct {
	value    []int
	heapSize int
}

// O(log n)
func (c *Heap) MaxHeapify(i int) {
	l := Left(i)
	r := Right(i)

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
		c.Change(i, massimo)
		c.MaxHeapify(massimo)
	}
}

// O(n)
func (c *Heap) BuildMaxHeap() {
	c.heapSize = len(c.value)

	for i := int(len(c.value)/2) - 1; i >= 0; i-- {
		c.MaxHeapify(i)
	}
}

func (c *Heap) Sort() {
	c.BuildMaxHeap()

	for i := len(c.value) - 1; i >= 1; i-- {
		c.Change(0, i)
		c.heapSize--
		c.MaxHeapify(0)
	}
}

func (c *Heap) Change(a int, b int) {
	temp := c.value[a]
	c.value[a] = c.value[b]
	c.value[b] = temp
}

func main() {
	fmt.Println("*** HEAP-SORT ***")

	var heap Heap = Heap{value: []int{4, 2, 3, 4, 2, 3}, heapSize: 0}

	fmt.Println(heap)

	heap.Sort()

	fmt.Println(heap)

}
