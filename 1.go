package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printMatrix(mat [][]int) {
	for i := range mat {
		for j := range mat[i] {
			fmt.Print(mat[i][j])
			fmt.Print(" ")
		}

		fmt.Println()
	}
}

func getMultiplication1(mat1 [][]int, mat2 [][]int) [][]int {
	m := len(mat1)
	n := len(mat1[0])
	o := len(mat2[0])

	mat3 := make([][]int, m)
	for i := range mat3 {
		mat3[i] = make([]int, o)
		for j := range mat3[i] {
			mat3[i][j] = 0
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < o; j++ {
			for k := 0; k < n; k++ {
				mat3[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}

	return mat3
}

// This one takes less than half the time of getMultiplication1 in row major mode because of cache locality
func getMultiplication2(mat1 [][]int, mat2 [][]int) [][]int {
	m := len(mat1)
	n := len(mat1[0])
	o := len(mat2[0])

	mat3 := make([][]int, m)
	for i := range mat3 {
		mat3[i] = make([]int, o)
		for j := range mat3[i] {
			mat3[i][j] = 0
		}
	}

	for i := 0; i < m; i++ {
		for k := 0; k < n; k++ {
			for j := 0; j < o; j++ {
				mat3[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}

	return mat3
}

func main() {
	fmt.Println("Started!!!")

	// m x n, n x o --> m x o

	m := 4096 / 4
	n := 4096 / 4
	o := 4096 / 4

	mat1 := make([][]int, m)
	for i := range mat1 {
		mat1[i] = make([]int, n)
		for j := range mat1[i] {
			mat1[i][j] = rand.Intn(10) // random values between 0 and 9
		}
	}
	// fmt.Println("Matrix 1")
	// printMatrix(mat1)
	// fmt.Println()

	mat2 := make([][]int, n)
	for i := range mat2 {
		mat2[i] = make([]int, o)
		for j := range mat2[i] {
			mat2[i][j] = rand.Intn(10) // random values between 0 and 9
		}
	}

	timeStart1 := time.Now()
	getMultiplication1(mat1, mat2)
	timeEnd1 := time.Now()
	fmt.Println("Time taken by getMultiplication1: ", timeEnd1.Sub(timeStart1))

	timeStart2 := time.Now()
	getMultiplication2(mat1, mat2)
	timeEnd2 := time.Now()
	fmt.Println("Time taken by getMultiplication2: ", timeEnd2.Sub(timeStart2))
}
