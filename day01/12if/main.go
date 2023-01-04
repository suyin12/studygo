package main

import "fmt"

func main() {
	ifDemo2()
}

func demo1(x int) {
	if i := x; i >= 100 {

	}
}

func ifDemo2() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	fmt.Println(score)
}
