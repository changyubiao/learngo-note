package main

import (
	"fmt"
)

/*
可变参数


一个函数 只能有一个可变参数
若参数列表中  还有其他类型的参数 , 则可变参数 写在所有参数的最后



*/

func main_01() {

	sum, avg, count := GetScore(90, 50.5, 60.2, 62.9)

	fmt.Printf("count=%d, sum = %.2f, avg = %.2f", count, sum, avg)

	fmt.Println()

	scores := []float64{90, 50.5, 60.2, 62.9}

	sum, avg, count = GetScore(scores...)
	fmt.Printf("count=%d, sum = %.2f, avg = %.2f\n", count, sum, avg)

}

func GetScore(scores ...float64) (sum, avg float64, count int) {

	for _, value := range scores {
		sum += value
		count++
	}

	avg = sum / float64(count)

	return
	//return sum, avg, count
}

func Mysum(a, b int, scores ...int) int {
	sum := 0
	for _, score := range scores {
		sum += score

	}
	return sum + a + b
}

func main() {
	a, b := 1, 2
	scores := []int{10, 20, 30}
	fmt.Println(Mysum(a, b, scores...))
}
