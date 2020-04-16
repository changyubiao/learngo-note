package main

import "fmt"

type user struct {
	name string
	age  byte
}

type manager struct {
	user
	title string
}

func main_01() {
	var m manager
	m.name = "Frank"
	m.age = 18
	m.title = "CTO"
	fmt.Println(m)
}

type X int

func (x *X) inc() {
	*x++
}

func main_02() {
	var x X
	x = 10
	x.inc()

	println("x=", x)
}

func (u user) ToString() string {

	return fmt.Sprintf("%+v", u)
}

func main_03() {

	var m manager
	m.name = "frank"

	m.age = 19
	println(m.ToString())
}

type Printer interface {
	Print()
}

func (u user) Print() {

	fmt.Printf("%+v\n", u)

}

func main() {

	var u user
	u.name = "frank"
	u.age = 22
	var p Printer = u

	p.Print()

}
