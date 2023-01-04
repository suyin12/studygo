package main

import "fmt"

func main() {
	// demo1()
	// demo2()
	// breakDemo1()
	continueDemo()
}

func demo1() {
	breakFlag := false
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				breakFlag = true
				break
			}
			fmt.Printf("i:%d j:%d\n", i, j)
		}
		if breakFlag {
			break
		}
	}
}

//使用goto语句能简化代码
func demo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakFlag
			}
			fmt.Printf("i:%d j:%d\n", i, j)
		}
	}
	return
breakFlag:
	fmt.Println("结束for循环")
}

func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
