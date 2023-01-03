package main

import "fmt"

//常量计数器iota
//iota是go语言的常量计数器，只能在常量的表达式中使用。

//iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

func main() {
	const (
		n1 = iota
		n2
		n3
	)
	fmt.Println(n1, n2, n3)

	const (
		m1 = iota //0
		m2        //1
		_         //使用匿名变量跳过某个值
		m4        //3
	)
	fmt.Println(m1, m2, m4)

	//iota声明中间插队
	const (
		x1 = iota //0
		x2 = 100  //100
		x3 = iota //2
		x4        //3
	)
	const x5 = iota //0
	fmt.Println(x1, x2, x3, x4, x5)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println(KB, MB, GB, TB, PB)

	//多个iota定义在一行
	const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
	fmt.Println(a, b, c, d, e, f)
}
