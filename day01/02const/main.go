package main

import "fmt"

const pi = 3.1415962
const e = 2.7182

func main() {
	//多个常量一起声明
	const (
		con_1 = "con_1"
		con_2 = "con_2"
	)
	fmt.Println(pi, e, con_1, con_2)

	//const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。 例如：
	const (
		n1 = 100
		n2
		n3
	)
	fmt.Println(n1, n2, n3)

}
