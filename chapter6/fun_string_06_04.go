package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
字符串中 中大小写 转换
1、func Title(s string) string
将字符串s每个单词首字母大写返回

2、func ToLower(s string) string
将字符串s转换成小写返回

3、func ToLowerSpecial(_case unicode.SpecialCase, s string) string
将字符串s中所有字符按_case指定的映射转换成小写返回

4、func ToTitle(s string) string
将字符串s转换成大写返回

5、func ToTitleSpecial(_case unicode.SpecialCase, s string) string
将字符串s中所有字符按_case指定的映射转换成大写返回

6、func ToUpper(s string) string
将字符串s转换成大写返回

7、func ToUpperSpecial(_case unicode.SpecialCase, s string) string
将字符串s中所有字符按_case指定的映射转换成大写返回



ToTitle   vs.  Title 的区别



参考链接
	https://docs.studygolang.com/pkg/strings/#ToLowerSpecial

*/

func main() {

	//var name = "frank is a GOOD student"

	//fmt.Println(name)

	//fmt.Println(strings.ToLower(name))
	//fmt.Println(strings.ToUpper(name))
	//
	//fmt.Println("name=", name)
	//
	//fmt.Println(strings.ToTitle(name))
	//fmt.Println(strings.Title("her royal highness"))
	//
	//fmt.Println(strings.ToTitle(name) == strings.ToUpper(name))

	//strings.ToLowerSpecial()
	fmt.Println("====ToLowerSpecial  testing=====")
	TestToLowerSpecial()

	TestToTitleSpecial()

}

func TestToLowerSpecial() {

	//fmt.Print("hello  ToLowerSpecial")
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
}

func TestToTitleSpecial() {
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
}
