package main

import "fmt"

func main() {

	fmt.Println(fact(5))

	fmt.Println(getMultiple(5))

}

func fact(n int) int {
	if n == 1 {
		return 1
	}

	return fact(n-1) * n

}

func getMultiple(n int) int {

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result

}
