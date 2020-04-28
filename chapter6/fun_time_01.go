package main

import (
	"fmt"
	"time"
)

/*
reference:
	https://docs.studygolang.com/pkg/time/
*/

/*
go 语言 time.go 时间库 常用的一些方法

//1、Now()返回当前本地时间

//2、Local()将时间转成本地时区，但指向同一时间点的Time。

//3、UTC()将时间转成UTC和零时区，但指向同一时间点的Time。

//4、Date()可以根据指定数值，返回一个本地或国际标准的时间格式。

//5、Parse()能将一个格式化的时间字符串解析成它所代表的时间。就是string转time

//6、Format()根据指定的时间格式，将时间格式化成文本。就是time转string

//7、String()将时间格式化成字符串，格式为："2006-01-02 15:04:05.999999999 -0700 MST"

//8、Unix()将t表示为Unix时间（时间戳，一个int64整数），即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）。

//9、UnixNano()将t表示为Unix时间（时间戳，一个int64整数），即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位纳秒）。

//10、Equal()判断时间是否相等

//11、Before()如果t代表的时间点在u之前，返回真；否则返回假。

//12、After()如果t代表的时间点在u之后，返回真；否则返回假。

//13、Date()返回时间点对应的年、月、日信息

//14、Year()返回时间点对应的年的信息

//15、Month()返回时间点对应的月的信息

//16、Day()返回时间点对应的日的信息

//17、Weekday()返回时间点对应的星期的信息

//18、Clock()返回时间点对应的时、分、秒信息

//19、Hour()返回时间点对应的小时的信息

//20、Minute()返回时间点对应的分的信息

//21、Second()返回时间点对应的秒的信息

//22、Nanosecond()返回时间点对应的纳秒的信息

//23、Sub()返回一个时间段t-u。

//24、Hours()将时间段表示为float64类型的小时数。

//25、Minutes()将时间段表示为float64类型的分钟数。

//26、Seconds()将时间段表示为float64类型的秒数。

//27、Nanoseconds()将时间段表示为int64类型的纳秒数，等价于int64(d)。

//28、String()返回时间段采用"72h3m0.5s"格式的字符串表示。

//29、ParseDuration解析一个时间段字符串。

//30、Add()返回时间点t+d。

//31、AddDate()返回增加了给出的年份、月份和天数的时间点Time。

*/

func main() {

	time1 := time.Now()
	testTime()
	//time.Sleep(5 * time.Second)

	time2 := time.Now()
	//计算函数执行时间
	fmt.Println(time2.Sub(time1).Seconds())

	TestSleep()

}

func expensiveCall() {
	fmt.Println("sleep 2 seconds")
	time.Sleep(time.Second * 2)
}

func TestSleep() {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

func testTime() {
	//1、Now()返回当前本地时间
	t := time.Now()
	fmt.Println("1、", t)

	//2、Local()将时间转成本地时区，但指向同一时间点的Time。
	fmt.Println("2、", t.Local())

	//3、UTC()将时间转成UTC和零时区，但指向同一时间点的Time。
	fmt.Println("3、", t.UTC())

	//4、Date()可以根据指定数值，返回一个本地或国际标准的时间格式。
	t = time.Date(2018, time.January, 1, 1, 1, 1, 0, time.Local)
	fmt.Printf("4、本地时间%s ， 国际统一时间：%s \n", t, t.UTC())

	//5、Parse()能将一个格式化的时间字符串解析成它所代表的时间。就是string转time
	//预定义的ANSIC、UnixDate、RFC3339
	//ANSIC       = "Mon Jan _2 15:04:05 2006"//1 1 2 3 4 5 6
	t, _ = time.Parse("2006-01-02 15:04:05", "2018-07-19 05:47:13")
	fmt.Println("5、", t)

	//6、Format()根据指定的时间格式，将时间格式化成文本。就是time转string
	fmt.Println("6、", time.Now().Format("2006-01-02 15:04:05"))

	//7、String()将时间格式化成字符串，格式为："2006-01-02 15:04:05.999999999 -0700 MST"
	fmt.Println("7、", time.Now().String())

	//8、Unix()将t表示为Unix时间（时间戳，一个int64整数），即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）。
	fmt.Println("8、", time.Now().Unix())

	//9、UnixNano()将t表示为Unix时间（时间戳，一个int64整数），即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位纳秒）。
	fmt.Println("9、", time.Now().UnixNano())

	//10、Equal()判断时间是否相等
	fmt.Println("10、", t.Equal(time.Now()))

	//11、Before()如果t代表的时间点在u之前，返回真；否则返回假。
	fmt.Println("11、", t.Before(time.Now()))

	//12、After()如果t代表的时间点在u之后，返回真；否则返回假。
	fmt.Println("12、", t.After(time.Now()))

	//13、Date()返回时间点对应的年、月、日信息
	year, month, day := time.Now().Date()
	fmt.Println("13、", year, month, day)

	//14、Year()返回时间点对应的年的信息
	fmt.Println("14、", time.Now().Year())

	//15、Month()返回时间点对应的月的信息
	fmt.Println("15、", time.Now().Month())

	//16、Day()返回时间点对应的日的信息
	fmt.Println("16、", time.Now().Day())

	//17、Weekday()返回时间点对应的星期的信息
	fmt.Println("17、", time.Now().Weekday())

	//18、Clock()返回时间点对应的时、分、秒信息
	hour, minute, second := time.Now().Clock()
	fmt.Println("18、", hour, minute, second)

	//19、Hour()返回时间点对应的小时的信息
	fmt.Println("19、", time.Now().Hour())

	//20、Minute()返回时间点对应的分的信息
	fmt.Println("20、", time.Now().Minute())

	//21、Second()返回时间点对应的秒的信息
	fmt.Println("21、", time.Now().Second())

	//22、Nanosecond()返回时间点对应的纳秒的信息
	fmt.Println("22、", time.Now().Nanosecond())

	//23、Sub()返回一个时间段t-u。
	fmt.Println("23、", time.Now().Sub(time.Now()))

	//24、Hours()将时间段表示为float64类型的小时数。
	fmt.Println("24、", time.Now().Sub(time.Now()).Hours())

	//25、Minutes()将时间段表示为float64类型的分钟数。
	fmt.Println("25、", time.Now().Sub(time.Now()).Minutes())

	//26、Seconds()将时间段表示为float64类型的秒数。
	fmt.Println("26、", time.Now().Sub(time.Now()).Seconds())

	//27、Nanoseconds()将时间段表示为int64类型的纳秒数，等价于int64(d)。
	fmt.Println("27、", time.Now().Sub(time.Now()).Nanoseconds())

	//28、String()返回时间段采用"72h3m0.5s"格式的字符串表示。
	fmt.Println("28、", "时间间距：", t.Sub(time.Now()).String())

	//29、ParseDuration解析一个时间段字符串。
	d, _ := time.ParseDuration("1h30m")
	fmt.Println("29、", d)

	//30、Add()返回时间点t+d。
	fmt.Println("30、", "交卷时间：", time.Now().Add(d))

	//31、AddDate()返回增加了给出的年份、月份和天数的时间点Time。
	fmt.Println("31、", "一年一个月零一天之后的日期：", time.Now().AddDate(1, 1, 1))

}
