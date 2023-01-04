package main

import "fmt"

func main() {
	// demo1()
	// demo2()
	// demo3()
	breakDemo1()
}

func demo1() {
	for i := 0; i < 10; i++ {
		// fmt.Println(i)
	}
}

func demo2() {
	i := 1
	for ; i < 10; i++ {
		// fmt.Println(i)
	}
}

func demo3() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
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
