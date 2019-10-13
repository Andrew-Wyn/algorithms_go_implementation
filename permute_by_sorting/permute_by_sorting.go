
/*

	utilizzare il merge sort per creare una permutazione 

	Permute By Sorting, Go Implementation


	utilizzo il merge sorting come algoritmo di ordinamento di default in quanto è un algoritmo ottimo, 
	complessita di merge_sort == complessita intrinseca del problema

	il costo complessivo di tale procedura è la stessa del merge sort

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pbs_merge(a []int, v []int, p int, q int, r int) {
	n1 := q-p+1
	n2 := r-q

	left :=  make([]int, n1)
	right := make([]int, n2)

	a_left := make([]int, n1)
	a_right := make([]int, n2)

	for i := 0; i < n1; i++ {
		left[i] = v[p+i]
		a_left[i] = a[p+i]
	}

	for i := 0; i < n2; i++ {
		right[i] = v[q+i+1]
		a_right[i] = a[q+i+1]
	}

	var i int = 0
	var j int = 0

	// rimanere nel bound delle pile, assicurarsi di non sforare con gli indici
	var lEmpty bool = false
	var rEmpty bool = false

	for k := p; k <= r; k++ {
		if len(left) == i && !rEmpty{
			v[k] = right[j]
			a[k] = a_right[j]
			j++
			lEmpty = true
		}

		if len(right) == j && !lEmpty{
			v[k] = left[i]
			a[k] = a_left[i]
			i++
			rEmpty = true
		}

		if !lEmpty && !rEmpty {			
			if left[i] <= right[j] {
				v[k] = left[i]
				a[k] = a_left[i]
				i++
			} else {
				v[k] = right[j]
				a[k] = a_right[j]
				j++
			}
		}
	}

}

func pbs_merge_sort(a []int, v []int, p int, r int) {
	if p < r {
		var q int = (p+r)/2
		
		pbs_merge_sort(a, v, p, q)
		pbs_merge_sort(a, v, q+1, r)
		pbs_merge(a, v, p, q, r)
	}
}

func random_inizializer(len int) []int {
	
	out := make([]int, len)
	
    r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	for i:=len-1; i>=0; i-- {
		out[i] = r1.Intn(int(len))
	}

	return out
}

func main() {
	fmt.Println("*** Permute By Sorting ***")
	
	var a = []int {4,5,6}

	vector := random_inizializer(len(a)) // inizializzare vector con numeri casuali da 1 a n-i

	fmt.Println(a)

	pbs_merge_sort(a, vector, 0, len(vector)-1)

	fmt.Println(a)
}