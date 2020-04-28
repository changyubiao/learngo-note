package main

import (
	"fmt"
	"strconv"
)

/*

reference:
	https://docs.studygolang.com/pkg/strconv/


strconv  中Parse类函数的使用

Atoi (string to int)

The most common numeric conversions are Atoi (string to int) and Itoa (int to string).

*/

/*
1、func Atoi(s string) (int, error)
Atoi 返回 ParseInt(s, 10, 0) 转换为 int 类型的结果。【alphabet：字母】

2、func ParseInt(s string, base int, bitSize int) (i int64, err error)
ParseInt 解释给定基础（2到36）中的字符串 s 并返回相应的值 i。如果 base == 0，则基数由字符串的前缀隐含：base 16代表“0x”，base 8代表“0”，否则以10为底数。

3、func ParseUint(s string, base int, bitSize int) (uint64, error)
ParseUint 就像 ParseInt，但是对于无符号数字。

4、func ParseFloat(s string, bitSize int) (float64, error)
ParseFloat 将字符串 s 转换为浮点数，精度由 bitSize：32指定，float32为64; float64为64。
当 bitSize = 32时，结果仍然具有 float64 类型，但可以在不更改其值的情况下将其转换为 float32。

5、func ParseBool(str string) (bool, error)
ParseBool 返回字符串表示的布尔值。
它接受 "1", "t", "T", "true", "TRUE", "True" 的字符串
"0", "f", "F", "false", "FALSE", "False"  的字符串
任何其他值都会返回错误。
*/

func main() {

	fmt.Println("hello strconv pkg")
	//TestItoa()
	//TestAtoi()
	//TestParseInt()

	//TestParseFloat()

	TestParseBool()
}

func TestItoa() {
	/*
		将 int 类型  转成 对应的 字符串
	*/
	var number int = 111
	number2 := strconv.Itoa(number)
	fmt.Printf("%T, r=%v\n", number2, number2)
}

func TestAtoi() {
	/*
		把 数字字符串 转成 对应的 int 类型的数字
	*/
	var number = "1111"
	number2, _ := strconv.Atoi(number)

	fmt.Printf("%T, r=%v\n", number2, number2)

}

func TestParseInt() {
	num, _ := strconv.ParseInt("-4e00", 16, 64)
	fmt.Printf("%T , %v \n", num, num)

	num, _ = strconv.ParseInt("01100001", 2, 64)
	fmt.Printf("%T , %v\n", num, num)

	num, _ = strconv.ParseInt("-01100001", 10, 64)
	fmt.Printf("%T , %v\n", num, num)

	num, _ = strconv.ParseInt("4e00", 10, 64)
	fmt.Printf("%T , %v\n", num, num)
	fmt.Println("---------------")
}

func TestParseFloat() {
	fmt.Println("-------TestParseFloat--------")

	pi := "3.1415926"
	num, _ := strconv.ParseFloat(pi, 64)

	fmt.Printf("%T ,num=%v\n", num, num)

	fmt.Printf("%T ,num*2=%v\n", num, num*2)
	fmt.Println("------TestParseFloat  end ---------")
}

func TestParseBool() {
	fmt.Println("--------TestParseBool-------")

	//1, t, T, TRUE, true, True,
	// 0, f, F, FALSE, false, False.
	flag, err := strconv.ParseBool("steven")
	fmt.Printf("%T , flag=%v , err=%v \n", flag, flag, err)

	flag, err = strconv.ParseBool("1")

	fmt.Println("--------TestParseBool  end -------")
}
