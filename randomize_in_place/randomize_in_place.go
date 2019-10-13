package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomizeInPlace(a []int) {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < len(a); i++ {
		nRandom := r1.Int()%(len(a)-i) + i
		temp := a[i]
		a[i] = a[nRandom]
		a[nRandom] = temp

		/* va e non va capire perche
		a[i] = a[i] ^ a[nRandom] // a+b
		a[nRandom] = a[nRandom] ^ a[i] // b+a+b = a
		a[i] = a[i] ^ a[nRandom] // a+b+ = b*/
	}
}

func main() {
	fmt.Println("*** Randomize in Place ***")

	v := []int{1, 2, 3, 1, 2, 3, 1, 2, 3}

	fmt.Println(v)

	randomizeInPlace(v)

	fmt.Println(v)
}
