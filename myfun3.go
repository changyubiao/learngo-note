package main

/*
函数可以作为值传递
*/
import (
	"fmt"
	"strings"
)

func main() {

	result := StringtoLower("aaaaaaaaaaa", processCase)

	fmt.Println(result)

	result2 := StringtoUpper("aaaaaaaaaaa", processCase)

	fmt.Println(result2)

}

func processCase(sequence string) (result string) {

	result = ""

	for i, value := range sequence {

		//fmt.Printf("i =%d,value=%v, type=%T  , value=%s \n", i, value, value, string(value))

		if i%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}

	}

	return result
}

func StringtoLower(str string, f func(string) string) string {

	fmt.Printf("%T  \n", f)

	return f(str)

}

// 声明一个新的类型, 这是一个函数类型
type caseFunc func(string) string

func StringtoUpper(str string, f caseFunc) string {

	fmt.Printf("%T  \n", f)

	return f(str)

}
