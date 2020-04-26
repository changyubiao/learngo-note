package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
字符串中 分隔类函数



1、func Fields(s string) []string
将字符串s以空白字符分割，返回一个切片


2、func FieldsFunc(s string, f func(rune) bool) []string
将字符串s以满足f(r)==true的字符分割，返回一个切片

3、func Split(s, sep string) []string
将字符串s以sep作为分隔符进行分割，分割后字符最后去掉sep

4、func SplitAfter(s, sep string) []string
将字符串s以sep作为分隔符进行分割，分割后字符最后附上sep

5、func SplitAfterN(s, sep string, n int) []string
将字符串s以sep作为分隔符进行分割，分割后字符最后附上sep，n决定返回的切片数



6、func SplitN(s, sep string, n int) []string
将字符串s以sep作为分隔符进行分割，分割后字符最后去掉sep，n决定返回的切片数
*/

func main() {

	//var name = "frank is a good student"
	//
	//fmt.Println(name)
	//
	//words := strings.Fields(name)
	//fmt.Printf("%T, %v \n", words, words)

	//TestSplit()

	//TestSplitAfter()

	//TestSplitN()

	//TestSplitAfterN()

	TestFieldsFunc()
}

func TestFields() {
	var name = "frank is a good student"
	fmt.Println(name)
	words := strings.Fields(name)
	fmt.Printf("%T, %v \n", words, words)
}

func TestFieldsFunc() {
	f := func(c rune) bool {
		//return c == '='
		//return c == 'X'
		//return unicode.IsLetter(c)
		// 不是字母并且不是数字
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)

	}
	fmt.Println(strings.FieldsFunc("abc@123*ABC&xyz%XYZ", f)) //[abc 123 ABC xyz XYZ]
}

func TestSplit() {
	/*
		分隔字符串
	*/
	var sentences = "hello world\nhello world2\nhello world3\n"

	sentence := strings.Split(sentences, "\n")

	fmt.Printf("%T , len=%d \n", sentence, len(sentence))
	for _, w := range sentence {

		fmt.Println("w=", w)
	}

}

func TestSplitAfter() {

	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) //["a," "b," "c"]
}

func TestSplitN() {
	//  返回一个 切片,n为切片的个数
	fmt.Printf("%q\n", strings.SplitN("a,b,c,d", ",", 2)) //["a"   "b,c,d"]
	fmt.Printf("%q\n", strings.SplitN("a,b,c,d", ",", 3)) //["a" "b" "c,d"]
}

func TestSplitAfterN() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d", ",", 2)) //["a,"   "b,c,d"]

	// 返回 包含sep 的分隔的一个切片，n 为返回切片的长度
	fmt.Printf("%q\n", strings.SplitAfterN("a#b#c#d#e", "#", 3)) //["a#" "b#" "c#d#e"]

}
