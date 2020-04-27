package main

import (
	"fmt"
	"strings"
)

/*
1、func Compare(a, b string) int
按字典顺序比较a和b字符串大小

2、func EqualFold(s, t string) bool
判断s和t两个utf8字符串是否相等，忽略大小写

3、func Repeat(s string, count int) string
将字符串s重复count次返回

4、func Replace(s, old, new string, n int) string
替换字符串s中old字符为new字符并返回，n<0是替换所有old字符串

n 用来 指定替换字符的次数


4、func ReplaceAll(s, old, new string,) string
替换字符串s中old字符为new字符并返回, 相当于 Replace 中 n < 0 ,即替换所有 字符

5、func Join(a []string, sep string) string
将a中的所有字符串连接成一个字符串，使用字符串sep作为分隔符

*/

func main() {

	// 比较 字符串 忽略大小写 比较 两个字符串是否相等
	fmt.Println(strings.EqualFold("aaa", "AAA"))
	fmt.Println(strings.EqualFold("aba", "ABA"))

	fmt.Println(strings.Repeat("aa", 3))

	fmt.Println("==== Compare  begin=====")
	fmt.Println(strings.Compare("aaa", "abc"))
	fmt.Println(strings.Compare("afa", "abc"))
	fmt.Println(strings.Compare("abc", "abc"))

	fmt.Println("==== Compare  end=====")

	fmt.Println("=== Replace begin === ")

	// n 用来指定 要替换的次数
	fmt.Println(strings.Replace("frank is a good student", "a", "A", 1))
	fmt.Println(strings.Replace("frank is a good student a aa", "a", "A", 2))

	fmt.Println("=== Replace end === ")

	fmt.Println("=== ReplaceAll begin === ")

	fmt.Println(strings.ReplaceAll("frank is a good student a aa", "a", "A"))
	fmt.Println("=== ReplaceAll end  === ")

	TestJoin()

}

func TestJoin() {
	fmt.Println("====  Join begin ==========")

	s := []string{"aa", "bb", "cc", "dd"}
	fmt.Printf("%q \n", strings.Join(s, "-"))

	s2 := []string{"I", "am", "Frank."}
	fmt.Printf("%q \n", strings.Join(s2, " "))

	fmt.Println("====  Join end =============== ")

}
