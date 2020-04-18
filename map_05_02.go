package main

import "fmt"

/*

获取 map 对应的值

value ,ok =  countryMap[key]
根据 ok 的值 来判断 map 中有没有对应的 key

	// 如果 key存在会返回  ok = true, 返回对应的值
	// 如果key 不存在 ok = false ,返回对应类型的默认值





删除 map 中的元素 delete
delete(map,key) 如果 key 不存在也不会报错，并且 delete 没有任何返回值

清空map 可以重新生成一个 map ,这里go语言 并没有提供一个清空map的方法



*/

func main_01() {

	// 3.   先 创建，后赋值
	countryMap := make(map[string]string)

	countryMap["China"] = "Beijing"
	countryMap["Janpan"] = "Tokyo"
	countryMap["India"] = "New Delhi"
	countryMap["France"] = "Pairs"
	countryMap["Italy"] = "Rome"

	// 遍历 map
	//for k, v := range countryMap {
	//	fmt.Println("k=", k, ", v=", v)
	//}

	value, ok := countryMap["china"]
	fmt.Printf("%q ,%t \n", value, ok)

	// 如果 key存在会返回  ok = true, 返回对应的值
	// 如果key 不存在 ok = false ,返回对应类型的默认值
	value2, ok := countryMap["China"]
	fmt.Printf("%q ,%t \n", value2, ok)
	//fmt.Printf("%q", countryMap["aaaa"])

	// 更加 go 的写法
	value3, ok := countryMap["China11"]

	if ok {
		fmt.Println("capital=", value3)
	} else {
		fmt.Println("Not exist key ")
	}

	// 这种写法 更加 go
	if value, ok := countryMap["China"]; ok {
		fmt.Println("capital=", value)
	} else {
		fmt.Println("Not exist key ")
	}

}

func main_02() {
	//  先 创建，后赋值
	countryMap := make(map[string]string)

	countryMap["China"] = "Beijing"
	countryMap["Janpan"] = "Tokyo"
	countryMap["India"] = "New Delhi"
	countryMap["France"] = "Pairs"
	countryMap["Italy"] = "Rome"

	delete(countryMap, "aaa")

	printMap(countryMap)

	fmt.Println("======= delete  Italy  France ======")
	delete(countryMap, "Italy")
	delete(countryMap, "France")
	printMap(countryMap)

	if _, ok := countryMap["Janpan"]; ok {
		delete(countryMap, "Janpan")
	}
	fmt.Println("=======  delete Janpan =====")
	printMap(countryMap)

	// clear map
	fmt.Println("======= clear  countryMap  =====")

	//countryMap = map[string]string{}
	countryMap = make(map[string]string)
	printMap(countryMap)
}

func main() {

	person := map[string]string{

		"name":   "frank",
		"gender": "male",
		"hobby":  "swimming",
	}

	// 语法不通过
	//r := delete(person, "aaaa")
	//fmt.Println(r)

	delete(person, "name")

	printMap(person)
}

func printMap(x map[string]string) {
	for k, v := range x {
		fmt.Printf("%q=%q \n", k, v)
		//fmt.Printf("%s=%s \n", k, v)
	}

}
