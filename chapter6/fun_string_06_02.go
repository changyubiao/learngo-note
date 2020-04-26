package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
检索字符串

中文文档
https://studygolang.com/pkgdoc


1、func Contains(s, substr string) bool
判断字符串s是否包含substr字符串

2、func ContainsAny(s, chars string) bool
判断字符串s是否包含chars字符串中的任一字符

3、func ContainsRune(s string, r rune) bool
判断字符串s是否包含unicode码值r

4、func Count(s, sep string) int
返回字符串s包含字符串sep的个数

5、func HasPrefix(s, prefix string) bool
判断字符串s是否有前缀字符串prefix

6、func HasSuffix(s, suffix string) bool
判断字符串s是否有后缀字符串suffix

7、func Index(s, sep string) int
返回字符串s中字符串sep首次出现的位置

8、func IndexAny(s, chars string) int
返回字符串chars中的任一unicode码值r在s中首次出现的位置

9、func IndexByte(s string, c byte) int
返回字符串s中字符c首次出现位置

10、func IndexFunc(s string, f func(rune) bool) int
返回字符串s中满足函数f(r)==true字符首次出现的位置

11、func IndexRune(s string, r rune) int
返回unicode码值r在字符串中首次出现的位置

12、func LastIndex(s, sep string) int
返回字符串s中字符串sep最后一次出现的位置

13、func LastIndexAny(s, chars string) int
返回字符串s中任意一个unicode码值r最后一次出现的位置

14、func LastIndexByte(s string, c byte) int
返回字符串s中 字符c 最后一次出现的位置

15、func LastIndexFunc(s string, f func(rune) bool) int
返回字符串s中满足函数f(r)==true字符最后一次出现的位置
*/

func main() {
	//fmt.Println("hello  fun 0602 ")

	//TestContains()
	//TestContainsAny()

	//TestContainsRune()

	//TestCount()
	//	TestHasPrefix()
	//	TestHasSuffix()

	//TestIndex()

	//TestIndexAny()
	//TestIndexRune()

	TestLastIndexAny()

	//TestLastIndexByte()

	//TestIndexFunc()
}

func TestLastIndexByte() {
	fmt.Println(strings.LastIndexByte("abcABCA123", 'A')) //6
}

func TestLastIndexAny() {
	//fmt.Println(strings.LastIndexAny("frank", "arn"))
	//fmt.Println(strings.LastIndexAny("frank", "nar"))
	fmt.Println(strings.LastIndexAny("我是中国人", "中国")) // 9
	fmt.Println(strings.LastIndexAny("我是中国人", "国中")) //9
	//fmt.Println(strings.LastIndexAny("chicken", "aeiouy")) //5
	//fmt.Println(strings.LastIndexAny("crwth", "aeiouy"))   //-1
}

func TestLastIndex() {
	fmt.Println(strings.LastIndex("Steven learn english", "e")) //13
	fmt.Println(strings.Index("go gopher", "go"))               //0
	fmt.Println(strings.LastIndex("go gopher", "go"))           //3
	fmt.Println(strings.LastIndex("go gopher", "rodent"))       //-1
}

func TestIndexRune() {
	fmt.Println(strings.IndexRune("abcABC120", 'C')) //5
	fmt.Println(strings.IndexRune("It培训教育", '教'))    //8
}

func TestIndexByte() {
	fmt.Println(strings.IndexByte("123abc", 'a')) //3
}

func TestIndexAny() {
	//??? 看不懂
	//返回字符串 chars 中的任意字符在字符串 s 中第一次出现的索引位置，若没有出现，返回-1。
	//fmt.Println(strings.IndexAny("abcABC120", "教育基地A")) //3

	fmt.Println(strings.IndexAny("chicken", "eih"))

}

func TestIndex() {
	fmt.Println(strings.Index("chicken", "ken")) //4

	// 没有找到返回 ; -1
	fmt.Println(strings.Index("chicken", "dmr")) //-1
}

func TestHasPrefix() {
	fmt.Println(strings.HasPrefix("1000phone news", "1000"))  //true
	fmt.Println(strings.HasPrefix("1000phone news", "1000a")) //false
}

func TestHasSuffix() {
	fmt.Println(strings.HasSuffix("1000phone news", "news")) //true
	fmt.Println(strings.HasSuffix("1000phone news", "new"))  //false
}

func TestCount() {
	// 统计字符串 s 中非重叠substr的数量。若统计空字符串""，会返回 s 的长度加1。
	fmt.Println(strings.Count("cheese", "e")) //3

	/*  */
	fmt.Println(strings.Count("one", "")) //4
}

func TestContainsRune() {
	fmt.Println(strings.ContainsRune("一丁丂", '丁'))   //true
	fmt.Println(strings.ContainsRune("一丁丂", 19969)) //true

	fmt.Println(strings.ContainsRune("geeksforgeeks", 97))  // true  'a'
	fmt.Println(strings.ContainsRune("geeksforgeeks", 103)) // true  'g'

}

func TestContains() {
	fmt.Println(strings.Contains("seafood", "foo"))   //true
	fmt.Println(strings.Contains("seafood", "bar"))   //false
	fmt.Println(strings.Contains("seafood", ""))      //true
	fmt.Println(strings.Contains("", ""))             //true
	fmt.Println(strings.Contains("steven王2008", "王")) //true
}

func TestContainsAny() {
	fmt.Println(strings.ContainsAny("team", "i"))        //false
	fmt.Println(strings.ContainsAny("failure", "u & i")) //true
	fmt.Println(strings.ContainsAny("foo", ""))          //false
	fmt.Println(strings.ContainsAny("imooc", "o"))       //false
}

func TestLastIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.LastIndexFunc("Hello,世界", f))       //9
	fmt.Println(strings.LastIndexFunc("Hello,world中国人", f)) //17
}

func TestIndexFunc() {
	f := func(c rune) bool {
		return c == 'r'
	}

	fmt.Println(strings.IndexFunc("hello world", f))

}
