package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func demo1() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Printf("jpgFunc: %T, jpgFunc: %v\n", jpgFunc, &jpgFunc)
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func lixiang(name string) {
	fmt.Println("name is ", name)
}

func low(f func()) {
	f()
}

// 匿名函数  闭包=函数+引用环境
func test(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

func main() {
	// demo1()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fc := test(lixiang, "sujianhui")
	low(fc)
}
