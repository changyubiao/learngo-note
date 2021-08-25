package main

import (
	"fmt"
	"strconv"
)

func add1(a, b int) int {
	return a + b
}

func sum() int {
	ret := 0
	for i := 1; i <= 100; i++ {
		ret += i
		fmt.Println("ret:", ret)
		//strconv.
		strconv.Atoi("1")
	}
	return ret
}
func main() {
	//go add1(1, 2)
	//fmt.Println(add1(1, 3))
	//fmt.Println(sum())
	i, err := strconv.Atoi("1000")
	if err != nil {
		panic("errror")
	}
	fmt.Println(i)
}
