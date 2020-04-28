package main

import "fmt"

func swap(a *int, b *int) {
	var tmp int = *a
	*a = *b
	*b = tmp

}

func main() {
	var a = 10
	var b = 20
	fmt.Printf("a=%d,b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("a=%d,b=%d\n", a, b)
}
