package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("a value:%v, a address:%p \n", a, &a)
	fmt.Printf("b value:%v, b address:%p b type:%T\n b value2:%v", b, &b, b, *b)
	// fmt.Println(1)

	demo1()
}

func demo1() {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}
