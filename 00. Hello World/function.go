package main

import "fmt"

func main() {

	fmt.Println(add(1, 2))

	fmt.Println(add2(1, 2))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x // multiple return values
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}
