package main

import "fmt"

/*


map 是引用类型  将一个map 赋值给另外一个变量的时候， 他们 会指向同一块内存



*/

func main() {

	personSalary := map[string]float32{
		"Steven": 10000,
		"June":   20000,
		"Charly": 30000,
	}

	//fmt.Println(personSalary)
	printPerson(personSalary)
	newPerson := personSalary
	// 添加一个 人
	newPerson["Frank"] = 100

	printPerson(newPerson)
	printPerson(personSalary)

	//fmt.Println("newPerson=", newPerson)
	//fmt.Println("person=", personSalary)

}

func printPerson(x map[string]float32) {

	fmt.Printf("address:%p \n", x)
	for k, v := range x {
		fmt.Printf(" %q=%v \n", k, v)
		//fmt.Printf("%s=%s \n", k, v)
	}

}
