package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(sqrt(-4), sqrt(2))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, limit float64) float64 {
	// initialize v in if loop
	if v := math.Pow(x, n); v < limit {
		return v
	}
	return limit
}

func pow2(x, n, limit float64) float64 {
	// initialize v in if-else loop
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}
	return limit
}
