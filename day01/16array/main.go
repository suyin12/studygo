package main

import "fmt"

//数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。 基本语法：

// 定义一个长度为3元素类型为int的数组a
var a [3]int

func main() {
	// b := [2]string{"1", "2"}
	// c := [3]string{"2", "2", "3"}
	// b = c

	// fmt.Println(a)
	// fmt.Println(b)
	// demo2()
	// demo3()
	// demo4()
	// demo5()
	// demo6()
	// demo7()
	// demo8()
	// demo9()
	demo10()
}

//定义数组
func demo1() {
	var testArray [3]int
	var testnum = [3]int{1, 4, 6}
	var teststring = [3]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)
	fmt.Println(testnum)
	fmt.Println(teststring)
}

//按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度
func demo2() {
	var testArray [3]int
	var testnum = [...]int{1, 4, 6}
	var teststring = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)
	fmt.Println(testnum)
	fmt.Printf("type of testnum:%T\n", testnum) //type of numArray:[2]int
	fmt.Println(teststring)
	fmt.Printf("type of teststring:%T\n", teststring) //type of numArray:[2]int
}

//我们还可以使用指定索引值的方式来初始化数组，例如:
func demo3() {
	var testArray = [...]int{1: 1, 3: 5}
	fmt.Println(testArray)
}

//数组的遍历两种方式
func demo4() {
	var testArray = [...]string{"北京", "上海", "深圳"}
	//方式一
	for i := 0; i < len(testArray); i++ {
		fmt.Println(testArray[i])
	}
	//方式二
	for k, v := range testArray {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}
}

//Go语言是支持多维数组的，我们这里以二维数组为例（数组中又嵌套数组）。
func demo5() {
	var testArray = [2][2]string{
		{"成都", "南宁"},
		{"广州", "东莞"},
	}
	fmt.Println(testArray)       //[[成都 南宁] [广州 东莞]]
	fmt.Println(testArray[1][1]) //支持索引取值:东莞
}

//二维数组的遍历
func demo6() {
	var testArray = [2][2]string{
		{"成都", "南宁"},
		{"广州", "东莞"},
	}
	for _, v1 := range testArray {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}

//注意： 多维数组只有第一层可以使用...来让编译器推导数组长度。例如：
func demo7() {
	var testArray = [...][2]string{
		{"成都", "南宁"},
		{"广州", "东莞"},
		{"惠州", "河源"},
		{"梅州", "江门"},
	}
	//错误的写法
	// var testArray2 = [2][...]string{
	// 	{"成都", "南宁"},
	// 	{"广州", "东莞"},
	// 	{"惠州", "河源"},
	// 	{"梅州", "江门"},
	// }
	for _, v1 := range testArray {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}

//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
func demo8() {
	var testArray = [...]int{100, 200, 300}
	modifyArray1(testArray)
	fmt.Println(testArray)

	var testArray2 = [...][2]int{
		{100, 200},
		{300, 400},
		{500, 600},
	}
	fmt.Println(testArray2)
}

func modifyArray1(x [3]int) {
	x[0] = 1000
}

func modifyArray2(x [3][2]int) {
	x[1][0] = 1000
}

//求和
func demo9() {
	var testArray = [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range testArray {
		sum += v
	}
	fmt.Println(sum)
}

//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
func demo10() {
	var testArray = [...]int{1, 3, 5, 7, 8}
	for k1, v1 := range testArray {
		for k2, v2 := range testArray {
			if (v1 + v2) == 8 {
				fmt.Println(k1, k2)
			}
		}
	}
}
