package main

import "fmt"

func main() {
	demo1()
	demo2()
	demo3()
}

func demo1() {
	a := 5
	switch a {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("default")
	}
}

//一个分支可以有多个值，多个case值中间使用英文逗号分隔。
func demo2() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println("未知")
	}
}

//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：
func demo3() {
	age := 20
	switch {
	case age < 20:
		fmt.Println("1")
	case age == 20:
		fmt.Println("2")
	case age > 20:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}
}
