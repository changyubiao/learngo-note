package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []float64{1, 9, 16, 25, 36}

	f := func(v float64) {
		v = math.Sqrt(v)
		fmt.Printf("%.2f \n", v)
	}
	visit(arr, f)

	visit(arr, func(v float64) {
		v = math.Pow(v, 2)
		fmt.Printf("%.2f \n", v)

	})

	fmt.Println("36*36 = ", math.Pow(36.0, 2))

}

//  定义一个函数 ,遍历切片 元素, 对每个元素进行处理
func visit(list []float64, f func(float64)) {

	for _, value := range list {

		f(value)

	}

}
