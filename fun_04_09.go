package main

import "fmt"

func main() {

	myfunc := Counter()
	fmt.Printf("%T\n", myfunc)

	fmt.Println(myfunc())
	fmt.Println(myfunc())
	fmt.Println(myfunc())

	// 创建新的函数, nextNumber1, 并查看结果
	//myfunc1 := Counter()
	//fmt.Println("myfunc1", myfunc1)
	//
	//fmt.Println(myfunc1())
	//fmt.Println(myfunc1())
	//fmt.Println(myfunc1())

}

func Counter() func() int {
	i := 0
	res := func() int {
		i += 1
		return i
	}

	//fmt.Println("Counter res=", res)
	fmt.Printf("%T,%v\n", res, res)
	return res
}
