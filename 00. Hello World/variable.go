package main

import "fmt"

var c, java, golang bool

func main() {

	var a int
	fmt.Println(a, c, java, golang)

	var b, c = true, "hello"
	fmt.Println(b, c)

	d, e := 4, "world"
	fmt.Println(d, e)
}
