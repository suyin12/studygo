package main

import "fmt"

func main() {
	// demo1()
	// demo2()
	// demo3()
	// demo4()
	// demo5()
	// demo6()
	// demo7()
	demo8()
}

/*声明切片类型的基本语法如下：

var name []T
其中，

name:表示变量名
T:表示切片中的元素类型
*/
func demo1() {
	var a []int
	var b []string
	var c = []bool{true, false}
	var d = []bool{false, true}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //true
	fmt.Println(c == nil) //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}

//基于数组的切片，引用类型， 左包含，右不包含，底层就是一个数组
func demo2() {
	var a = [...]int{1, 2, 3, 4, 5}
	var b = a[1:3]
	// fmt.Println(b)
	a[2] = 100
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", b, len(b), cap(b))
}

func demo3() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
}

/*使用make()函数构造切片
ake([]T, size, cap)
其中：

T:切片的元素类型
size:切片中元素的数量
cap:切片的容量
*/
func demo4() {
	var a = make([]int, 5, 10)

	var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(len(a) == 0)

	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1 == nil)

	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(s2 == nil)

	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
	fmt.Println(s3 == nil)

}

func demo5() {
	var a = make([]int, 3)
	s1 := a
	s1[0] = 100
	fmt.Println(a)
	fmt.Println(cap(a))
	fmt.Println(s1)
}

//切片的遍历
func demo6() {
	a := []int{1, 2, 3}
	//方式1
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
	//方式2
	for i, v := range a {
		fmt.Println(i, v)
	}
}

//append()方法为切片添加元素
func demo7() {
	//通过var声明的零值切片可以在append()函数直接使用，无需初始化。
	// 	没有必要像下面的代码一样初始化一个切片再传入append()函数使用，

	// s := []int{}  // 没有必要初始化
	// s = append(s, 1, 2, 3)

	// var s = make([]int)  // 没有必要初始化
	// s = append(s, 1, 2, 3)
	var s []int
	s = append(s, 1)
	fmt.Println(s)
	s = append(s, 2, 3, 4)
	fmt.Println(s)

	// var s2 = make([]int, 5)
	// s2 = append(s2, 1, 2)
	// fmt.Println(s2)
	var s3 = []int{5, 6, 7}
	s = append(s, s3...)
	fmt.Println(s)
}

func demo8() {
	//append()添加元素和切片扩容
	// 每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，
	// 此时该切片指向的底层数组就会更换。“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	//从上面的结果可以看出：
	// append()函数将元素追加到切片的最后并返回该切片。
	// 切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。
}
