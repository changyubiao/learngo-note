package main // 代码包声明语句

// 代码包导入语句
import (
	"fmt" // 导入代码包fmt。
)

// main函数
func main() { // 代码块由“{”和“}”包裹。

	// 变量声明和赋值语句，由关键字var、变量名num、变量类型uint64、特殊标记=，以及值10组成。
	var num uint64 = 65535

	// 短变量声明语句，由变量名size、特殊标记:=，以及值（需要你来填写）组成。
	size := 8

	// 打印函数调用语句。在这里用于描述一个uint64类型的变量所需占用的比特数。
	// 这里用到了字符串的格式化函数。
	fmt.Printf("类型为 uint64 的整数 %d 需占用的存储空间为 %d 个字节。\n", num, size)
}
