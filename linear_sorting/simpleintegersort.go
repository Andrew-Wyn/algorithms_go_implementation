package main

import "fmt"

// valori di K da 0 a len(K)-1

func simpleIntegerSort(K []int) {
	p := 1
	n := len(K)

	for p <= n {

		fmt.Println(K)

		for p <= n && K[p-1] == p {
			p++
		}

		q := K[p-1]

		for q != p {
			h := K[q-1]
			K[q-1] = K[p-1]
			q = h
		}

		p++

	}
}

func main() {
	fmt.Println("*** simple integer sort ***")

	K := []int{1, 3, 2}

	simpleIntegerSort(K)

	fmt.Println(K)

}
