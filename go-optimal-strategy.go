package main

import (
	"fmt"
	"math"
)

func optimalStrategy(coins []float64) float64 {
	return optimalStrategy_recursive(coins, 0, len(coins)-1)
}
func optimalStrategy_recursive(coins []float64, i int, j int) float64 {
	if i >= len(coins) || j < 0 {
		return 0
	}
	if j-i == 1 {
		return math.Max(coins[i], coins[j])
	}

	A := coins[i] + math.Min(optimalStrategy_recursive(coins, i+1, j-1), optimalStrategy_recursive(coins, i+2, j))
	B := coins[j] + math.Min(optimalStrategy_recursive(coins, i+1, j-1), optimalStrategy_recursive(coins, i, j-2))
	C := math.Max(A, B)
	return C
}

func main() {
	var test1 = []float64{8, 15, 3, 7}
	fmt.Println(optimalStrategy(test1))
	var test2 = []float64{5, 15, 5, 1}
	fmt.Println(optimalStrategy(test2))
}
