package main

import "fmt"

/*
map  是无序的

map 的长度是不固定， map  不能用 cap() 计算容量，   len() 可以用来计算 map 的长度

map 中的key 的值要是唯一的，
key 的数据类型 必须是 可以参与比较运算的类型，也是就是 支持 == ，!= 操作类型，
	比如   bool  int   float  string  array  这些都是可以作为key


	像   slice  func   等引用类型 是不能作为key 的，因为他们 不可比较，或者说 不可哈希

	map中key可以是：int、float、bool、string、数组
	一定不可以是：切片、函数、map


*/
func main01() {
	// 定义一个 一个map     map[key_type]value_type
	var info map[string]string

	info2 := make(map[string]string)

	fmt.Printf("%T, %T\n", info, info2)

}

func main_02() {
	fmt.Println("go run map0501.go")

	// 1 声明，然后初始化一个map
	var numbers = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(numbers)

	// 2. 短变量 声明方式
	rates := map[string]int{
		"go":     1,
		"c++":    2,
		"python": 3,
		"java":   4,
	}

	fmt.Println(rates)

	// 3.   先 声明map，后赋值
	countryMap := make(map[string]string)

	countryMap["China"] = "Beijing"
	countryMap["Janpan"] = "Tokyo"
	countryMap["India"] = "New Delhi"
	countryMap["France"] = "Pairs"
	countryMap["Italy"] = "Rome"

	// 遍历 map
	for k, v := range countryMap {

		fmt.Println("k=", k, ", v=", v)
	}

	for k, _ := range countryMap {
		fmt.Println("k=", k)
	}
	for _, v := range countryMap {
		fmt.Println("v=", v)
	}

	fmt.Println("==================================")
	for k := range countryMap {
		fmt.Println("k=", k, ", v=", countryMap[k])

	}

	fmt.Printf("len(countryMap)=%d \n", len(countryMap))
}

func main() {

	person := map[string]string{

		"name":   "frank",
		"gender": "male",
		"hobby":  "swimming",
	}

	fmt.Printf("len(person)=%d \n", len(person))
	//fmt.Printf("cap(person)=%d \n", cap(person))  // error

}
