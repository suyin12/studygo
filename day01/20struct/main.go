package main

import (
	"fmt"
	"unsafe"
)

// 类型定义
type NewInt int

// 类型别名
type MyInt = int

type person struct {
	name string
	age  int
	city string
}

func main() {
	// demo1()
	// demo3()
	// demo2()
	// demo4()
	// demo5()
	// demo6()
	demo8()
}

func demo1() {
	var a NewInt
	var b MyInt
	// a = 10
	// b = 10
	fmt.Println(a, b)
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
}

// 实例化结构体
func demo2() {
	var p1 person
	p1.name = "su"
	p1.age = 18
	p1.city = "深圳"
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)
}

// 匿名结构体
func demo3() {
	var user struct {
		name string
		age  int
	}
	user.name = "su"
	user.age = 20
	fmt.Printf("user value is %v\n", user)
	fmt.Printf("user value is %#v\n", user)
}

// 指针类型结构体
func demo4() {
	var p = new(person)
	p.name = "su"
	p.age = 18
	p.city = "北京"
	fmt.Printf("p=%T\n", p)
	fmt.Printf("p=%#v\n", p)
}

// 取结构体的地址实例化
// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次 new 实例化操作
func demo5() {
	p := &person{}
	p.name = "七米" //p3.name = "七米"其实在底层是(*p3).name = "七米"，这是Go语言帮我们实现的语法糖。
	p.age = 18
	p.city = "广州"
	fmt.Printf("p=%T\n", p)
	fmt.Printf("p=%#v\n", p)
}

// 使用键值对初始化
func demo6() {
	p := person{
		name: "su",
		age:  20,
		city: "东莞",
	}
	fmt.Printf("p=%T\n", p)
	fmt.Printf("p=%#v\n", p)

	//结构体指针进行键值对初始化
	p1 := &person{
		name: "建晖",
		age:  22,
		city: "惠州",
	}
	fmt.Printf("p=%T\n", p1)
	fmt.Printf("p=%#v\n", p1)

	//当某些字段没有值可以不填，默认值为原来的
	p2 := &person{
		city: "南京",
	}
	fmt.Printf("p2=%#v\n", p2)

	//使用值的列表初始化
	// 使用这种格式初始化时，需要注意：
	// 必须初始化结构体的所有字段。
	// 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 该方式不能和键值初始化方式混用。
	p3 := &person{
		"晖",
		18,
		"佛山",
	}
	fmt.Printf("p3=%#v\n", p3)

}

// 结构体占用一块连续的内存。
type test struct {
	a int8
	b int8
	c int8
	d int8
}

func demo7() {
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
	fmt.Printf("n.a %v\n", n.a)
	fmt.Printf("n.b %v\n", n.b)
	fmt.Printf("n.c %v\n", n.c)
	fmt.Printf("n.d %v\n", n.d)
	fmt.Println(unsafe.Sizeof(n))

	//空结构体
	// 空结构体是不占用空间的。
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 0

}

type student struct {
	name string
	age  int
}

func demo8() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
