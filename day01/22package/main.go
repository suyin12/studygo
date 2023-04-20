package main

import (
	"fmt"
	"studygo/day01/22package/demo"
	"github.com/q1mi/hello"
)

var x int8 = 10

const pi = 3.14

func init() {
	fmt.Println("x:", x)
	fmt.Println("pi:", pi)
	sayHi()
}

func sayHi() {
	fmt.Println("Hello World!")
}

func main() {
	//首字母小写，对外不可见
	// demo.say()

	//首字母大写，对外可见
	fmt.Println(demo.Model)
	sum := demo.Add(10, 30)
	fmt.Println(sum)
	fmt.Println("你好，世界！")

	hello.SayHi()//调用hello包的SayHi函数
}
