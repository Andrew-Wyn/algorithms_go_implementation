package main

import "fmt"

func integerSort(K []int, m int) {
	Y := make([]int, m+1)

	for i := 0; i < len(K); i++ {
		Y[K[i]] = Y[K[i]] + 1
	}

	p := 0

	fmt.Println(Y)

	for i := 0; i <= m; i++ {
		for j := 0; j < Y[i]; j++ {
			K[p] = i
			fmt.Println(Y[j])
			p++
		}
	}

}

func main() {
	fmt.Println("*** Integer Sort ***")

	K := []int{0, 12, 1}

	fmt.Println(K)

	integerSort(K, 12)

	fmt.Println(K)

}
