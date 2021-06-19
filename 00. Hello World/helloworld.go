package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Hello golang")
	fmt.Println("The time is", time.Now())
	fmt.Println("Random number Generate", rand.Intn(10))

	fmt.Printf("Now you have %g problems. \n\n", math.Sqrt(7))
	fmt.Println(math.Pi)
}
