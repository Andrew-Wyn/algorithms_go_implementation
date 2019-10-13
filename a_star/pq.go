package main

// PQ priority queue
type PQ struct {
	heap Heap
}

func (p *PQ) empty() bool {
	return p.heap.heapSize == 0
}

func (p *PQ) put(obj dic) {
	p.heap.MinHeapInsert(obj)
}

func (p *PQ) get() point {
	return p.heap.HeapExtractMin()
}
