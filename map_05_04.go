package main

import "fmt"

/*


map 初始化 ，赋值 遍历 等操作



*/

func main() {
	//1、声明时同时初始化
	var countryMap = map[string]string{
		"China":  "Beijing",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
		"France": "Paris",
		"Italy":  "Rome",
	}
	printCountry(countryMap)

	//	3、遍历map（无序）
	//	(1)、key 、value都遍历
	for k, v := range countryMap {
		fmt.Println("国家", k, "首都", v)
	}
	fmt.Println("-----------")

	//(2)、只展示value
	for _, v := range countryMap {
		fmt.Println("国家", "首都", v)
	}
	fmt.Println("-----------")

	//(3)、只展示key
	for k := range countryMap {
		fmt.Println("国家", k, "首都", countryMap[k])
	}
	fmt.Println("-----------")

	//4、查看元素是否在map中存在
	value, ok := countryMap["England"]
	fmt.Printf("%q \n", value)
	fmt.Printf("%T , %v \n", ok, ok)
	if ok {
		fmt.Println("首都：", value)
	} else {
		fmt.Println("首都信息未检索到！")
	}

	//或者
	if value, ok := countryMap["USA"]; ok {
		fmt.Println("首都：", value)
	} else {
		fmt.Println("首都信息未检索到！")
	}

}

func printCountry(x map[string]string) {

	fmt.Printf("address:%p \n", x)
	for k, v := range x {
		fmt.Printf(" %q=%v \n", k, v)
		//fmt.Printf("%s=%s \n", k, v)
	}

}
