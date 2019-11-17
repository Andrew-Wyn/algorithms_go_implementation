package main

import (
	"fmt"
	"math"
)

// Node nodo
type Node struct {
	value float64
	next  *Node
}

type List struct {
	head   *Node
	length int
}

func PrintList(head *Node) {
	if head != nil {
		fmt.Println(head.value)
		PrintList(head.next)
	}
}

func (list *List) getArray() (O []float64) {
	O = make([]float64, list.length)

	appo := list.head
	i := 0
	for appo != nil {
		O[i] = appo.value
		appo = appo.next
		i++
	}

	return

}

func makeListFromArray(A []float64) (list List) {

	list.length = 0
	for i := len(A) - 1; i >= 0; i-- {
		list.Push(A[i])
	}

	return
}

func (list *List) Push(value float64) {
	newNode := new(Node)

	newNode.value = value

	if list.head != nil {
		newNode.next = list.head
		(*list).head = newNode
	} else {
		newNode.next = nil
		(*list).head = newNode
	}

	list.length++
}

func (list *List) Pop() (out float64) {

	out = 3

	if list.head != nil {
		out = list.head.value
		list.head = list.head.next
		list.length--
	}

	return
}

// m sottoinsiemi

func BucketSort(A []float64, m int) {

	B := make([]List, m)

	// creo i bucket O(n)
	for i := 0; i < len(A); i++ {
		bucket := int(math.Floor((A[i] * float64(m))))
		B[bucket].Push(A[i])
	}

	// Ordina i vari bucket -> vedere costo computazionale sul cormen
	for i := 0; i < m; i++ {
		App := B[i].getArray()
		insertionSort(App)
		B[i] = makeListFromArray(App)
	}

	// concatena B[] O(n)
	p := 0
	for i := 0; i < m; i++ {
		PrintList(B[i].head)
		arr := B[i].getArray()
		for j := 0; j < len(arr); j++ {
			A[p] = arr[j]
			p++
		}
	}

}

func insertionSort(A []float64) {
	for i := 1; i < len(A); i++ {
		x := A[i]
		j := i - 1
		for j >= 0 && x < A[j] {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = x
	}
}

func min(a float64, b float64) float64 {
	if a < b {
		return a
	}

	return b
}

func main() {

	fmt.Println("*** try reusing linked list ****")

	A := []float64{0.1, 0.4, 0.3, 0.6, 0.8}

	fmt.Println(A)

	BucketSort(A, 3)

	fmt.Println(A)
}
