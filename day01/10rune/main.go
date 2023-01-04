package main

import (
	"fmt"
	"math"
)

func main() {
	s1 := "big"
	bytes1 := []byte(s1)
	bytes1[0] = 'p'
	fmt.Println(string(bytes1))

	s2 := "白萝卜"
	runes2 := []rune(s2)
	runes2[0] = '红'
	fmt.Println(string(runes2))

	sqrtDemo()
}

// 类型转换
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

	fmt.Printf("%T value is %d", c, c)
}
