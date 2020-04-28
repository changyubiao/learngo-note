package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
https://wrfly.kfd.me/posts/go-trimleft-and-trimprefix/



https://docs.studygolang.com/pkg/strings/#ToTitleSpecial

*/

func main_01() {
	//fmt.Println(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!"))
	//fmt.Println(strings.TrimPrefix("Hello, Gophers!!!", "Hello"))

	//cfgs := "mongodb://mongodb://nogiff"
	//cfgs = strings.TrimLeft(cfgs, "mongodb://")
	//
	//fmt.Printf("cfgs:%v \n", cfgs)
	//
	//cfgs2 := strings.TrimPrefix(cfgs, "mongodb://")
	//fmt.Printf("%T, cfgs2:%v \n", cfgs2, cfgs2)

	fmt.Println(strings.TrimLeft("/some/key", "/some"))
	fmt.Println(strings.TrimLeft("/some/sugar", "/some"))

}

func main02() {

	nodeOnlinePrefix := "/nodesli"
	kvs := []string{
		"/nodes/online/1ebbaa95eba5",
		"/nodes/online/d39d0dd1c159",
		"/nodes/online/d66b00076d8c",
	}

	var nodes []string
	for _, v := range kvs {
		str := fmt.Sprintf("%s", v)
		fmt.Printf("online node got key %s\n", str)
		n := strings.TrimLeft(str, nodeOnlinePrefix)
		fmt.Printf("online node got node %s\n", n)
		nodes = append(nodes, strings.TrimSpace(n))
	}
	fmt.Printf("online nodes: %v\n", nodes)

	/*
		这个所谓的cutset就是一个字符集, 而TrimLeft就是从左开始, 如果发现了不在cutset中的字符, 就从这个点返回.

		比如有一个字符串 "1112345111", TrimLeft("1112345111", "1") 之后就成了 2345111,
		对应的, TrimRight("1112345111", "1")就变成了"1112345"; Trim("1112345111", "1")则是从两头都砍, 返回的结果是"2345"

		这是cutset中只有一个字符的情况, 如果集合中有多个字符,那么就从里面开始匹配, 什么时候找不到了, 就返回.
	*/
}

func main() {
	for i := 1; i <= 5; i++ {
		//int转换为字符串：Itoa()
		fmt.Println("I like you very much " + strconv.Itoa(i))

	}

	//number, _ := strconv.Atoi("123")
	//
	//fmt.Printf("%T, number=%d\n", number, number)

	var i = 97
	fmt.Println("I like you very much " + string(i))
	fmt.Println("I like you very much " + strconv.Itoa(i))

	//fmt.Println(strconv.Atoi("123"))
}
