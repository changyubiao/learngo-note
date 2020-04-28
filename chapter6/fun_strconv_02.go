package main

import (
	"fmt"
	"strconv"
)

/*
reference:
	https://docs.studygolang.com/pkg/strconv/

strconv  中  Format 类函数的使用


1、func Itoa(i int) string
Itoa 是 FormatInt(int64(i), 10) 的缩写。

2、func FormatInt(i int64, base int) string
FormatInt 返回给定基数中的i的字符串表示，对于2 <= base <= 36.结果对于数字值> = 10使用小写字母 'a' 到 'z' 。

3、func FormatUint(i uint64, base int) string
FormatUint 返回给定基数中的 i 的字符串表示，对于2 <= base <= 36.结果对于数字值> = 10使用小写字母 'a' 到 'z' 。

	FormatUint  只能转无符号的数字


4、func FormatFloat(f float64, fmt byte, prec, bitSize int) string
FormatFloat 根据格式 fmt 和 precision prec 将浮点数f转换为字符串。它将结果进行四舍五入，假设原始数据是从 bitSize 位的浮点值获得的（float32为32，float64为64）。
格式 fmt 是 'b'，'e'，'E'，'f'，'g'或 'G'。

5、func FormatBool(b bool) string
FormatBool 根据 b 的值返回“true”或“false”




*/

func main() {

	//fmt.Println("hello strconv pkg 2")
	//TestItoa()
	//TestFormatInt()
	//TestFormatUint()

	TestFormatBool()
}

func TestFormatBool() {
	fmt.Println("----------TestFormatBool begin ----------")

	s := strconv.FormatBool(true)
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))

	s = strconv.FormatBool(false)
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))
	fmt.Println("----------TestFormatBool end ----------")
}

func TestFormatUint() {
	//  转 无符号数字
	fmt.Println("--------TestFormatUint begin--------------")

	s := strconv.FormatUint(19968, 16) //4e00
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))

	s = strconv.FormatUint(40869, 16) //9fa5
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))

	s = strconv.FormatUint(1121, 10) //9fa5
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))
	fmt.Println("--------TestFormatUint  end --------------")
}

func TestFormatInt() {
	fmt.Println("-------TestFormatInt begin---------")

	s := strconv.FormatInt(-19968, 16) //4e00
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))

	s = strconv.FormatInt(-40869, 16) //9fa5
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))

	s = strconv.FormatInt(1121, 10) //9fa5
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))

	s = strconv.FormatInt(-1121, 10) //9fa5
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))

	fmt.Println("-------TestFormatInt end---------")
}

func TestItoa() {

	str := strconv.Itoa(123321321321)

	fmt.Printf("%T,  %v\n", str, str)
}
