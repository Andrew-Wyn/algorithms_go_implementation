package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readMatrix(path string) [][]int {

	var matrix [][]int

	file, err := os.Open("matrix.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	fmt.Println()

	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println(n)

	matrix = make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	i := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		for j := range line {
			nij, _ := strconv.Atoi(line[j])
			matrix[i][j] = nij
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matrix
}

func submatrix(matrix [][]int, ri int, rj int) [][]int {
	out := make([][]int, len(matrix)-1)
	for i := 0; i < len(matrix)-1; i++ {
		out[i] = make([]int, len(matrix)-1)
	}

	ii := 0
	for i := 0; i < len(matrix); i++ {
		jj := 0
		if i != ri {
			for j := 0; j < len(matrix); j++ {
				if j != rj {
					out[ii][jj] = matrix[i][j]
					jj++
				}
			}
			ii++
		}
	}
	return out
}

func determinant(matrix [][]int) int { // il bello che faccio ancora errori da superiori, saro mai un programmatore decente ?

	if len(matrix) == 1 {
		return matrix[0][0]
	} else {
		det := 0

		for i := 0; i < len(matrix); i++ {
			det += int(math.Pow(float64(-1), float64(2+i))) * matrix[i][0] * determinant(submatrix(matrix, i, 0))
		}

		return det
	}

}

func invertible(matrix [][]int) bool {
	return determinant(matrix) != 0
}

func gaussFactorization(matrix [][]int) ([][]int, [][]int) {
	fmt.Println("---- factorizing ----")

	n := len(matrix)

	U := make([][]int, n)
	L := make([][]int, n)
	for i := 0; i < n; i++ {
		U[i] = make([]int, n)
		copy(U[i], matrix[i])
		L[i] = make([]int, n)
		L[i][i] = 1
	}

	for k := 0; k < n-1; k++ {
		for i := k + 1; i < n; i++ {
			L[i][k] = U[i][k] / U[k][k]
			for j := k + 1; j < n; j++ {
				U[i][j] = U[i][j] - L[i][k]*U[k][j]
			}
			U[i][k] = 0
		}
	}

	return L, U

}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func matrixMultiplication(L [][]int, U [][]int) [][]int {

	n := len(L)

	A := make([][]int, n)

	for i := 0; i < n; i++ {
		A[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				A[i][j] += L[i][k] * U[k][j]
			}
		}
	}

	return A
}

func main() {

	matrix := readMatrix("matrix.txt")

	fmt.Println("---- MATRIX ----")
	printMatrix(matrix)
	fmt.Println(invertible(matrix)) // invertibile se il determinante è diverso da 0

	L, U := gaussFactorization(matrix)

	fmt.Println("---- L ----")
	printMatrix(L)
	fmt.Println()
	fmt.Println("---- U ----")
	printMatrix(U)
	fmt.Println()
	fmt.Println("---- A ----")
	printMatrix(matrixMultiplication(L, U))
	fmt.Println()
}
