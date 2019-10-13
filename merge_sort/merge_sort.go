/*

	Merge sort GO implementation, utilizzo di flag booleani per ovviare alla massimalizzazione tramite infiniti della formalizzazione algoritmica

	vedi libro cormen

*/

/*

	un po di documentazione

	make([]int, n) -> This creates a slice that is associated with an underlying int array of length n. Slices are always associated with some array, and although they can never be longer than the array, they can be smaller. The make function also allows a 3rd parameter:

*/

package main

import (
	"fmt"
)

func merge(v []int, p int, q int, r int) {
	n1 := q-p+1
	n2 := r-q

	left :=  make([]int, n1)
	right := make([]int, n2)

	for i := 0; i < n1; i++ {
		left[i] = v[p+i]
	}

	for i := 0; i < n2; i++ {
		right[i] = v[q+i+1]
	}

	var i int = 0
	var j int = 0

	// rimanere nel bound delle pile, assicurarsi di non sforare con gli indici
	var lEmpty bool = false
	var rEmpty bool = false

	for k := p; k <= r; k++ {
		if len(left) == i && !rEmpty{
			v[k] = right[j]
			j++
			lEmpty = true
		}

		if len(right) == j && !lEmpty{
			v[k] = left[i]
			i++
			rEmpty = true
		}

		if !lEmpty && !rEmpty {			
			if left[i] <= right[j] {
				v[k] = left[i]
				i++
			} else {
				v[k] = right[j]
				j++
			}
		}
	}

}

func merge_sort(v []int, p int, r int) {
	if p < r {
		var q int = (p+r)/2
		
		merge_sort(v, p, q)
		merge_sort(v, q+1, r)
		merge(v, p, q, r)
	}
}

func main() {
	fmt.Println("merge sort")
	
	var vector = []int {2,5,1,4} 

	fmt.Println(vector)

	merge_sort(vector, 0, len(vector)-1)

	fmt.Println(vector)
}