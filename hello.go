package hello

import (
	"errors"
	"fmt"
)

func main_01() {

	//fmt.Println("hello  world")
	//fmt.Println(runtime.GOARCH)
	x := 100

	if x > 0 {
		println("x >0")
	} else if x < 0 {
		println("x <0 ")
	} else {
		println("0")
	}
}

func main_02() {

	var x int

	x = -100

	switch {
	case x > 0:
		println("x>0")
	case x < 0:
		println("x<0")
	default:
		println("x=0")

	}

}

func main_03() {
	var i int
	for i = 0; i < 5; i++ {
		fmt.Printf("i=%d, ", i)
		//println("")

	}
	fmt.Println("")

	for i = 5; i > 0; i-- {
		fmt.Printf("i=%d, ", i)
	}
	fmt.Println("")

}

func main_04() {
	x := 5

	for x > 0 {
		println(x)
		x--
	}
}

func main_05() {
	var i = 0
	for {
		println(i)
		if i == 5 {
			break
		}
		i++
	}
}

func main_06() {
	numbers := []int{100, 101, 102}

	for i, n := range numbers {

		println(i, ":", n)

	}
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main_07() {

	a, b := 10, 0

	c, err := div(a, b)

	fmt.Println(c, err)
}

func decor(x int) func() {
	//闭包

	return func() {
		println(x)
	}
}

func main_08() {

	var x = 200
	f := decor(x)

	f()
}

func test_defer(a, b int) {
	// defer 常用来做释放资源,解除锁定,一些 清理操作等
	defer println("dispose ...")

	println(a / b)
}

func main_09() {
	test_defer(10, 0)
}

func main_10() {
	// 切片用法,动态数组
	x := make([]int, 0, 5)

	for i := 0; i < 10; i++ {
		x = append(x, i)
	}

	fmt.Println(x)
}

func main_11() {
	// map 类型
	d := make(map[string]string)

	d["name"] = "frank"

	//println(d)

	fmt.Println(d)
	x, ok := d["b"]

	fmt.Println(x, ok)

	delete(d, "name")

	fmt.Println(d)

}
