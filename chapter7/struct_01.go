package main

import "fmt"

type Teacher struct {
	name   string
	age    int
	gender byte
}

/*
结构体的初始化
*/
func main() {
	var t1 Teacher
	fmt.Printf("t1:%T,%v,%q \n", t1, t1, t1)
	t1.name = "frank"
	t1.age = 19
	t1.gender = 0
	fmt.Printf("t1:%T,%v,%q \n", t1, t1, t1)

	// 2
	t2 := Teacher{
		name:   "frank2",
		age:    20,
		gender: 1,
	}
	//fmt.Printf("t1:%T,%v,\n",t1,t1)
	printTeacher(t2)

	// 3
	t3 := Teacher{}
	t3.name = "laoda"
	t3.age = 28
	t3.gender = 1
	printTeacher(t3)

	// 4
	t4 := Teacher{"lile", 29, 1}
	printTeacher(t4)

	//5、创建指针类型的结构体
	t5 := new(Teacher)
	fmt.Printf("t5:%T , %v , %p \n", t5, t5, t5)
	//(*t5).name = "Running"
	//(*t5).age = 31
	//(*t5).sex = 0

	//语法简写形式——语法糖
	t5.name = "Running2"
	t5.age = 31
	t5.gender = 0
	//printTeacher(t5)
	fmt.Println(t5)
	//fmt.Println("-------------------")

}

func printTeacher(t Teacher) {

	fmt.Printf("t:%T,%v \n", t, t)

}
