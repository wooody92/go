package main

import "fmt"

func main() {

	delayPrint()
	calculateFirstPrintLater(10)
}

func delayPrint() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func calculateFirstPrintLater(max int) {
	fmt.Println("counting")

	for i := 0; i < max; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
