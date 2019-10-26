// MIN_HEAP

package main

// MaxInt the maximum integer
const MaxInt = int(^uint(0) >> 1)

// MinInt the minimum integer
const MinInt = -MaxInt - 1

type dic struct {
	key   float64
	value point
}

type point struct {
	x int
	y int
}

// Heap is a simple heap struct
type Heap struct {
	value    []dic
	heapSize int
}

func parent(i int) int {
	i++
	return (i >> 1) - 1
}

func left(i int) int {
	i++
	return (i << 1) - 1
}

func right(i int) int {
	i++
	return (i << 1)
}

// O(log n)
func (c *Heap) minHeapify(i int) {
	l := left(i)
	r := right(i)

	minimo := -1

	if l < c.heapSize && c.value[l].key < c.value[i].key {
		minimo = l
	} else {
		minimo = i
	}

	if r < c.heapSize && c.value[r].key < c.value[minimo].key {
		minimo = r
	}

	if minimo != i {
		c.change(i, minimo)
		c.minHeapify(minimo)
	}
}

// O(n)
func (c *Heap) buildMinHeap() {
	c.heapSize = len(c.value)

	for i := int(len(c.value)/2) - 1; i >= 0; i-- {
		c.minHeapify(i)
	}
}

func (c *Heap) change(a int, b int) {
	temp := c.value[a]
	c.value[a] = c.value[b]
	c.value[b] = temp
}

// HeapExtractMax -> get
func (c *Heap) HeapExtractMin() point {
	if c.heapSize < 1 {
		panic("underflow dell'heap")
	}

	min := c.value[0].value
	c.value[0] = c.value[c.heapSize-1]
	c.heapSize--
	c.minHeapify(0)

	// avoid pseudo memory leak
	appo := make([]dic, c.heapSize)
	copy(appo, c.value[0:c.heapSize])
	c.value = appo

	return min
}

func (c *Heap) heapIncreaseKey(i int, key float64) {
	if key > c.value[i].key {
		panic("la nuova chiave è piu grande di quella corrente")
	}

	c.value[i].key = key

	for i > 0 && c.value[parent(i)].key > c.value[i].key { // for is golang's while
		c.change(i, parent(i))
		i = parent(i)
	}
}

// MaxHeapInsert -> put
func (c *Heap) MinHeapInsert(obj dic) {

	if c.heapSize == len(c.value) {
		c.value = append(c.value, dic{})
	}

	c.value[c.heapSize].key = float64(MaxInt)
	c.value[c.heapSize].value = obj.value
	c.heapSize++

	c.heapIncreaseKey(c.heapSize-1, obj.key)
}

func (c *Heap) heapMaximum() point {
	return c.value[0].value
}

func (c *Heap) build() {
	c.buildMinHeap()
}
