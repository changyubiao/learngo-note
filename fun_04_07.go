package main

import "fmt"

func main_01() {

	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d \t", i)

		fmt.Println(add2(i))
	}

}

func add2(x int) int {
	sum := 0
	sum += x
	return sum
}

/*  ======= 闭包演示 =========   */
func adder() func(int) int {
	sum := 0

	f := func(x int) int {
		fmt.Printf("sum1 = %d\t", sum)

		sum += x
		fmt.Printf("sum2 = %d\t", sum)
		return sum

	}
	return f
}

func main() {
	pos := adder()

	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d \t", i)
		fmt.Println(pos(i))
	}

	fmt.Println("--------------------------")
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d \t", i)
		fmt.Println(pos(i))
	}
}
