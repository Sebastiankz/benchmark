package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func spiralMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	dx, dy := 0, 1
	x, y := 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[x][y] = i
		if matrix[(x+dx+n)%n][(y+dy+n)%n] != 0 {
			dx, dy = dy, -dx
		}
		x += dx
		y += dy
	}
	return matrix
}

func addPrimeNumbers(matrix [][]int) int64 {
	var sum int64
	for _, row := range matrix {
		for _, num := range row {
			if isPrime(num) {
				sum += int64(num)
			}
		}
	}
	return sum
}

func main() {

	start := time.now()

	N := 1000
	createSpiralMatrix := spiralMatrix(N)
	fmt.Println("Matriz Espiral de 1000x1000:")
	for i := 0; i < 10; i++ {
		fmt.Println(createSpiralMatrix[i])
	}

	suma := addPrimeNumbers(createSpiralMatrix)
	fmt.Printf("Suma de los números primos de la matriz: %d\n", suma)

	executionTime := time.Since(start).Milliseconds()
	fmt.Printf("Tiempo de ejecución: %d ms\n", executionTime)

}