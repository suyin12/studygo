package main

import "fmt"

func main() {
	// 十进制
	var a int = 10
	var a2 int = 0b1010
	fmt.Printf("%d \n", a)  // 10
	fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制
	fmt.Printf("%d \n", a2) // 10
	fmt.Printf("%b \n", a2) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77
	fmt.Printf("%d \n", b) // 十进制的63

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF
	fmt.Printf("%d \n", c) // 十进制 255
}
