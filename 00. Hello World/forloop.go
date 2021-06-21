package main

import "fmt"

func main() {

	sum(10)
	sum2(10)
	sum3(10)
	infiniteLoop()
}

func sum(max int) {
	sum := 0
	for i := 0; i < max; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func sum2(max int) {
	sum := 1
	for sum < max {
		sum += sum
	}
	fmt.Println(sum)
}

func sum3(max int) {
	// while
	sum := 1
	for sum < max {
		sum += sum
	}
	fmt.Println(sum)
}

func infiniteLoop() {
	for {
	}
}
