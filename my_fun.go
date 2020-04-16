package main

import "fmt"

func add(a, b int) (int, int) {
	var c int

	c = a + b
	d := 10
	//return d, c
	return c, d
}

func main() {
	fmt.Println(add(3, 5))
}
