// piority queue, go implementation

package main

import (
	"fmt"
)

// MaxInt the maximum integer
const MaxInt = int(^uint(0) >> 1)

// MinInt the minimum integer
const MinInt = -MaxInt - 1

type dic struct {
	key   int
	value string
}

// Heap is a simple heap struct
type Heap struct {
	value    []dic
	heapSize int
}

func parent(i int) int {
	return (i >> 1)
}

func left(i int) int {
	return (i << 1)
}

func right(i int) int {
	return (i << 1) ^ 1
}

// O(log n)
func (c *Heap) maxHeapify(i int) {
	l := left(i)
	r := right(i)

	massimo := -1

	if l < c.heapSize && c.value[l].key > c.value[i].key {
		massimo = l
	} else {
		massimo = i
	}

	if r < c.heapSize && c.value[r].key > c.value[massimo].key {
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

func (c *Heap) change(a int, b int) {
	temp := c.value[a]
	c.value[a] = c.value[b]
	c.value[b] = temp
}

func (c *Heap) heapExtractMax() string {
	if c.heapSize < 1 {
		panic("underflow dell'heap")
	}
	max := c.value[1].value
	c.value[0] = c.value[c.heapSize-1]
	c.heapSize--
	c.maxHeapify(0)
	return max
}

func (c *Heap) heapIncreaseKey(i int, key int) {
	if key < c.value[i].key {
		panic("la nuova chiave è piu piccola di quella corrente")
	}

	c.value[i].key = key

	for i > 0 && c.value[parent(i)].key < c.value[i].key { // for is golang's while
		c.change(i, parent(i))
		i = parent(i)
	}
}

func (c *Heap) maxHeapInsert(obj dic) {

	if c.heapSize == len(c.value) {
		panic("maximum lenght of the heap reached")
	}

	c.heapSize++
	c.value[c.heapSize-1].key = MinInt
	c.value[c.heapSize-1].value = obj.value
	c.heapIncreaseKey(c.heapSize-1, obj.key)
}

func (c *Heap) heapMaximum() string {
	return c.value[0].value
}

func main() {
	fmt.Println("*** PRIORITY-QUEUE ***")

	var heap Heap = Heap{value: []dic{{key: 1, value: "a"}, {key: 2, value: "b"}}, heapSize: 0}

	fmt.Println(heap)

	heap.buildMaxHeap()

	fmt.Println(heap)

	fmt.Println(heap.heapMaximum())
	fmt.Println(heap.heapExtractMax())

	heap.maxHeapInsert(dic{key: 3, value: "il potere della fica mi attanaglia"})

	fmt.Println(heap)
}
