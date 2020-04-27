package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*

strings 修剪字符串


1、func Trim(s string, cutset string) string
将字符串s中首尾包含cutset中的任一字符去掉返回

2、func TrimFunc(s string, f func(rune) bool) string
将字符串s首尾满足函数f(r)==true的字符去掉返回

3、func TrimLeft(s string, cutset string) string
将字符串s左边包含cutset中的任一字符去掉返回

4、func TrimLeftFunc(s string, f func(rune) bool) string
将字符串s左边满足函数f(r)==true的字符去掉返回

5、func TrimRight(s string, cutset string) string
将字符串s右边包含cutset中的任一字符去掉返回

6、func TrimRightFunc(s string, f func(rune) bool) string
将字符串s右边满足函数f(r)==true的字符去掉返回
字符串 从右向左看， 满足 f(r) ==true 的字符删除，如果没有满足，停止匹配。


7、func TrimSpace(s string) string
将字符串s首尾空白去掉返回

8、func TrimPrefix(s, prefix string) string
将字符串s中前缀字符串prefix去掉返回

9、func TrimSuffix(s, suffix string) string
将字符串s中后缀字符串prefix去掉返回



注意：

TrimLeft  vs  TrimPrefix    区别

TrimRight  vs
*/

func main() {
	//fmt.Println(strings.Trim("aaabbb ccc   aa", "aaa"))
	//TestTrim()

	//TestTrimFunc()
	//TestTrimLeft()

	//TestTrimRight()

	//TestTrimPrefix()
	//TestTrimSuffix()

	TestTrimSpace()

	TestTrimSuffix()
}

func TestTrimFunc1() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.TrimFunc("！@#￥%steven wang%￥#@", f)) //steven wang
}

func TestTrimSpace() {

	// 去掉两边的空格
	s := "   mysql://  frank@host   "
	fmt.Printf("s=%q \n", strings.TrimSpace(s))
}

func TestTrimPrefix() {

	s := "mysql://frank@host"

	fmt.Printf("s=%q \n", strings.TrimPrefix(s, "mysql://"))

}

func TestTrimSuffix() {
	var s = "AAAWebTool_Prod"
	//s = strings.TrimSuffix(s, "_Prod") //Hello,

	fmt.Printf("s=%q \n", strings.TrimSuffix(s, "_Prod"))

}

func TestTrim() {
	s := strings.Trim("aaa  bbb ccc   aa", "aaa")
	fmt.Printf("s=%q \n", s)

	// 去除空格
	s2 := strings.Trim("   aa bb  cc  dd  ", " ")
	fmt.Printf("s2=%q \n", s2)
}

func TestTrimFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	s1 := strings.TrimFunc("! @#$ frank chang 1 %&*(@#", f)
	fmt.Printf("s1=%q \n", s1) // frank chang 1
}

func TestTrimLeftFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.TrimLeftFunc("！@#￥%steven wang%￥#@", f)) //steven wang%￥#@
}

func TestTrimRightFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.TrimRightFunc("！@#￥%steven wang%￥#@", f)) //！@#￥%steven wang
}

func TestTrimLeft() {
	s1 := "aaa bbb ccc ddd"
	// 从 字符串中从左到右 寻找 有没有在cutset 集合里面，如果有就删掉，如果没有了就停止寻找了
	fmt.Printf("s1=%q\n", strings.TrimLeft(s1, "ab ")) // "ccc ddd"

	fmt.Printf("s1=%q\n", strings.TrimLeft(s1, "ab")) // " bbb ccc ddd"

}

func TestTrimRight() {
	s1 := "aaa bbb cccdddfff"
	// 从 字符串中从左到右 寻找 有没有在cutset 集合里面，如果有就删掉，如果没有了就停止寻找了
	fmt.Printf("s1=%q\n", strings.TrimRight(s1, "cfd")) // "aaa bbb "

	fmt.Printf("s1=%q\n", strings.TrimRight(s1, "f")) // "aaa bbb cccddd"
}
