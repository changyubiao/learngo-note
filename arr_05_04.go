package main

import "fmt"

func main() {

	a := [...]string{"AAA", "BBBB", "CCCC", "DDDD"}

	//	这里是值传递
	b := a

	b[0] = "Frank"

	fmt.Println("a=", a)
	fmt.Println("b=", b)

}
