package main

import "fmt"

//全局变量
var m = 100

func foo() (int, string) {
	return 10, "jianhui"
}

func main() {
	//标准声明
	var name string = "jianhui"
	var age int = 20
	var isOk bool
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOk)
	//批量声明
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//一次初始化多个变量
	var name1, age1 = "jianhui2", 20
	fmt.Println(name1)
	fmt.Println(age1)

	//类型推导
	var name3 = "jianhui3"
	var age3 = 21
	fmt.Println(name3)
	fmt.Println(age3)

	//短变量声明
	nn := 200
	mm := 300
	fmt.Println(nn, mm)

	//匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
