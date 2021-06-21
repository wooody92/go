package main

import "fmt"

func main() {

	pointer()
}

func pointer() {
	i, j := 10, 30

	p := &i
	fmt.Println(p)
	fmt.Println(*p)

	*p = 20
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(p)
	fmt.Println(*p)

	p = &j
	*p = *p / 10
	fmt.Println(j)
}
